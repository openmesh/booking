package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/openmesh/booking"
	"github.com/openmesh/booking/ent"
	"github.com/openmesh/booking/ent/migrate"
	"github.com/openmesh/booking/http"
	"github.com/openmesh/booking/metrics"
	"github.com/pelletier/go-toml"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	_ "github.com/openmesh/booking/ent/runtime"
	logging "github.com/openmesh/booking/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

// Build version, injected during build.
var (
	version string
	commit  string
)

// main is the entry point to our application binary. However, it has some poor
// usability, so we mainly use it to delegate out to our Main type.
func main() {
	// Propagate build information to root package to share globally.
	booking.Version = strings.TrimPrefix(version, "")
	booking.Commit = commit

	// Setup signal handlers.
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() { <-c; cancel() }()

	// Instantiate a new type to represent our application.
	// This type lets us shared setup code with our end-to-end tests.
	m := NewMain()

	// Parse command line flags & load configuration.
	if err := m.ParseFlags(ctx, os.Args[1:]); err == flag.ErrHelp {
		os.Exit(1)
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Execute program.
	if err := m.Run(ctx); err != nil {
		m.Close()
		fmt.Fprintln(os.Stderr, err)
		booking.ReportError(ctx, err)
		os.Exit(1)
	}

	// Wait for CTRL-C.
	<-ctx.Done()

	// Clean up program.
	if err := m.Close(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Main represents the program.
type Main struct {
	// Configuration path and parsed config data.
	Config     Config
	ConfigPath string

	// Ent client used by ent service implementations.
	Client *ent.Client

	// HTTP server for handling HTTP communication.
	// Ent services are attached to it before running.
	HTTPServer *http.Server
}

func NewMain() *Main {
	return &Main{
		Config:     DefaultConfig(),
		ConfigPath: DefaultConfigPath,

		Client:     ent.NewClient(),
		HTTPServer: http.NewServer(),
	}
}

// Close gracefully stops the program.
func (m *Main) Close() error {
	if m.HTTPServer != nil {
		if err := m.HTTPServer.Close(); err != nil {
			return err
		}
	}
	if m.Client != nil {
		if err := m.Client.Close(); err != nil {
			return err
		}
	}
	return nil
}

// ParseFlags parses the command line arguments & loads the config.
//
// This exists separately from the Run() function so that we can skip it
// during end-to-end tests. Those tests will configure manually and call Run().
func (m *Main) ParseFlags(ctx context.Context, args []string) error {
	// Our flag set is very simple. It only includes a config path.
	fs := flag.NewFlagSet("booking", flag.ContinueOnError)
	fs.StringVar(&m.ConfigPath, "config", DefaultConfigPath, "config path")
	if err := fs.Parse(args); err != nil {
		return err
	}

	// The expand() function is here to automatically expand "~" to the user's
	// home directory. This is a common task as configuration files are typing
	// under the home directory during local development.
	configPath, err := expand(m.ConfigPath)
	if err != nil {
		return err
	}

	// Read our TOML formatted configuration file.
	config, err := ReadConfigFile(configPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("config file not found: %s", m.ConfigPath)
	} else if err != nil {
		return err
	}
	m.Config = config

	return nil
}

// Run executes the program. The configuration should already be set up before
// calling this function.
func (m *Main) Run(ctx context.Context) (err error) {
	// // Initialize error tracking.
	// if m.Config.Rollbar.Token != "" {
	// 	rollbar.SetToken(m.Config.Rollbar.Token)
	// 	rollbar.SetEnvironment("production")
	// 	rollbar.SetCodeVersion(version)
	// 	rollbar.SetServerRoot("github.com/benbjohnson/wtf")
	// 	wtf.ReportError = rollbarReportError
	// 	wtf.ReportPanic = rollbarReportPanic
	// 	log.Printf("rollbar error tracking enabled")
	// }

	// Initialize event service for real-time events.
	// We are using an in-memory implementation but this could be changed to
	// a more robust service if we expanded out to multiple nodes.

	// eventService := inmem.NewEventService()

	// Attach our event service to the SQLite database so it can publish events.
	// m.DB.EventService = eventService

	// Expand the DSN (in case it is in the user home directory ("~")).
	// Then open the database. This will instantiate the SQLite connection
	// and execute any pending migration files.
	dsn, err := expandDSN(m.Config.DB.DSN)
	if err != nil {
		return fmt.Errorf("cannot expand dsn: %w", err)
	}
	if m.Client, err = ent.Open("postgres", dsn); err != nil {
		return fmt.Errorf("cannot open db: %w", err)
	}

	if err := m.Client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true)); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	// Create dependencies used by service middlewares
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	fieldKeys := []string{"method"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "open_mesh",
		Subsystem: "booking",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	errorCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "open_mesh",
		Subsystem: "booking",
		Name:      "error_count",
		Help:      "Number of errors occurred.",
	}, fieldKeys)
	requestDuration := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "open_mesh",
		Subsystem: "booking",
		Name:      "request_duration",
		Help:      "Total duration of requests in microseconds",
	}, fieldKeys)

	// Instantiate ent-backed services.
	// authService := ent.NewAuthService(m.Client)
	var resourceService booking.ResourceService
	{
		resourceService = ent.NewResourceService(m.Client)
		resourceService = booking.ResourceValidationMiddleware()(resourceService)
		resourceService = logging.ResourceLoggingMiddleware(logger)(resourceService)
		resourceService = metrics.ResourceMetricsMiddleware(requestCount, errorCount, requestDuration)(resourceService)
	}
	var bookingService booking.BookingService
	{
		bookingService = ent.NewBookingService(m.Client)
		bookingService = booking.BookingValidationMiddleware()(bookingService)
		bookingService = logging.BookingLoggingMiddleware(logger)(bookingService)
		bookingService = metrics.BookingMetricsMiddleware(requestCount, errorCount, requestDuration)(bookingService)
	}
	var organizationService booking.OrganizationService
	{
		organizationService = ent.NewOrganizationService(m.Client)
		organizationService = logging.OrganizationLoggingMiddleware(logger)(organizationService)
	}
	var unavailabilityService booking.UnavailabilityService
	{
		unavailabilityService = ent.NewUnavailabilityService(m.Client)
		unavailabilityService = logging.UnavailabilityLoggingMiddleware(logger)(unavailabilityService)
		unavailabilityService = metrics.UnavailabilityMetricsMiddleware(requestCount, errorCount, requestDuration)(unavailabilityService)
	}

	// userService := sqlite.NewUserService(m.DB)

	// Attach user service to Main for testing.
	// m.UserService = userService

	// Set global GA settings.
	// html.MeasurementID = m.Config.GoogleAnalytics.MeasurementID

	// Copy configuration settings to the HTTP server.
	m.HTTPServer.Addr = m.Config.HTTP.Addr
	m.HTTPServer.Domain = m.Config.HTTP.Domain
	m.HTTPServer.HashKey = m.Config.HTTP.HashKey
	m.HTTPServer.BlockKey = m.Config.HTTP.BlockKey
	m.HTTPServer.GitHubClientID = m.Config.GitHub.ClientID
	m.HTTPServer.GitHubClientSecret = m.Config.GitHub.ClientSecret

	// Attach underlying services to the HTTP server.
	// m.HTTPServer.AuthService = authService
	m.HTTPServer.ResourceService = resourceService
	m.HTTPServer.OrganizationService = organizationService
	m.HTTPServer.UnavailabilityService = unavailabilityService
	m.HTTPServer.BookingService = bookingService
	// m.HTTPServer.EventService = eventService
	// m.HTTPServer.UserService = userService

	// Attach logger to server.
	m.HTTPServer.AttachLogger(logger)
	// Register routes that require services to be initialized.
	m.HTTPServer.RegisterRoutes()

	// Start the HTTP server.
	if err := m.HTTPServer.Open(); err != nil {
		return err
	}

	// If TLS enabled, redirect non-TLS connections to TLS.
	if m.HTTPServer.UseTLS() {
		go func() {
			logger.Log(http.ListenAndServeTLSRedirect(m.Config.HTTP.Domain))
		}()
	}

	// Enable internal debug endpoints.
	go func() { http.ListenAndServeDebug() }()

	logger.Log("status", "running", "url", m.HTTPServer.URL(), "debug", "http://localhost:6060", "dsn", m.Config.DB.DSN)

	return nil
}

const (
	// DefaultConfigPath is the default path to the application configuration.
	DefaultConfigPath = "./booking.conf"

	// DefaultDSN is the default datasource name.
	DefaultDSN = "~/.booking/db"
)

// Config represents the CLI configuration file.
type Config struct {
	DB struct {
		DSN string `toml:"dsn"`
	} `toml:"db"`

	HTTP struct {
		Addr     string `toml:"addr"`
		Domain   string `toml:"domain"`
		HashKey  string `toml:"hash-key"`
		BlockKey string `toml:"block-key"`
	} `toml:"http"`

	GoogleAnalytics struct {
		MeasurementID string `toml:"measurement-id"`
	} `toml:"google-analytics"`

	GitHub struct {
		ClientID     string `toml:"client-id"`
		ClientSecret string `toml:"client-secret"`
	} `toml:"github"`

	Rollbar struct {
		Token string `toml:"token"`
	} `toml:"rollbar"`
}

// DefaultConfig returns a new instance of Config with defaults set.
func DefaultConfig() Config {
	var config Config
	config.DB.DSN = DefaultDSN
	return config
}

// ReadConfigFile unmarshals config from
func ReadConfigFile(filename string) (Config, error) {
	config := DefaultConfig()
	if buf, err := ioutil.ReadFile(filename); err != nil {
		return config, err
	} else if err := toml.Unmarshal(buf, &config); err != nil {
		return config, err
	}
	return config, nil
}

// expand returns path using tilde expansion. This means that a file path that
// begins with the "~" will be expanded to prefix the user's home directory.
func expand(path string) (string, error) {
	// Ignore if path has no leading tilde.
	if path != "~" && !strings.HasPrefix(path, "~"+string(os.PathSeparator)) {
		return path, nil
	}

	// Fetch the current user to determine the home path.
	u, err := user.Current()
	if err != nil {
		return path, err
	} else if u.HomeDir == "" {
		return path, fmt.Errorf("home directory unset")
	}

	if path == "~" {
		return u.HomeDir, nil
	}
	return filepath.Join(u.HomeDir, strings.TrimPrefix(path, "~"+string(os.PathSeparator))), nil
}

// expandDSN expands a datasource name. Ignores in-memory databases.
func expandDSN(dsn string) (string, error) {
	if dsn == ":memory:" {
		return dsn, nil
	}
	return expand(dsn)
}

func seedData(ctx context.Context, s booking.OrganizationService) {
	organization := booking.Organization{
		Name: "Default",
		Owner: &booking.User{
			Name:  "Admin",
			Email: "admin@local",
		},
	}
	err := s.CreateOrganization(ctx, &organization)
	if err != nil {
		os.Exit(1)
	}
}

// // rollbarReportError reports internal errors to rollbar.
// func rollbarReportError(ctx context.Context, err error, args ...interface{}) {
// 	if booking.ErrorCode(err) != booking.EINTERNAL {
// 		return
// 	}

// 	// Set user information for error, if available.
// 	if u := booking.UserFromContext(ctx); u != nil {
// 		rollbar.SetPerson(fmt.Sprint(u.ID), u.Name, u.Email)
// 	} else {
// 		rollbar.ClearPerson()
// 	}

// 	log.Printf("error: %v", err)
// 	rollbar.Error(append([]interface{}{err}, args)...)
// }

// // rollbarReportPanic reports panics to rollbar.
// func rollbarReportPanic(err interface{}) {
// 	log.Printf("panic: %v", err)
// 	rollbar.LogPanic(err, true)
// }
