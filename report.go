package booking

import "context"

// Dashboard reports

// - Recent sales
// - Upcoming bookings
// - Today's next bookings
// - Top resources

type ReportService interface {
	GetRecentSalesReport(ctx context.Context, req GetRecentSalesReportRequest) GetRecentSalesReportResponse
	GetUpcomingBookingsReport(ctx context.Context, req GetUpcomingBookingsReportRequest) GetUpcomingBookingsReportResponse
	GetBookingsActivityReport(ctx context.Context, req GetBookingsActivityReportRequest) GetBookingsActivityReportResponse
	GetTodaysBookingsReport(ctx context.Context, req GetTodaysBookingsReportRequest) GetTodaysBookingsReportResponse
	GetTopResourcesReport(ctx context.Context, req GetTopResourcesReportRequest) GetTopResourcesReportResponse
	GetTopEmployeesReport(ctx context.Context, req GetTopEmployeesReportRequest) GetTopEmployeesReportResponse
}

// GetRecentSalesReportRequest represents a payload used by the GetRecentSalesReport method of a ReportService
type GetRecentSalesReportRequest struct{}

// Validate a GetRecentSalesReport. Returns a ValidationError for each requirement that fails.
func (r GetRecentSalesReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetRecentSalesReportResponse represents a response returned by the GetRecentSalesReport method of a ReportService.
type GetRecentSalesReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetRecentSalesReportResponse) Error() error { return r.Err }

// GetUpcomingBookingsReportRequest represents a payload used by the GetUpcomingBookingsReport method of a ReportService
type GetUpcomingBookingsReportRequest struct { // insert request properties here
}

// Validate a GetUpcomingBookingsReport. Returns a ValidationError for each requirement that fails.
func (r GetUpcomingBookingsReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetUpcomingBookingsReportResponse represents a response returned by the GetUpcomingBookingsReport method of a ReportService.
type GetUpcomingBookingsReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetUpcomingBookingsReportResponse) Error() error { return r.Err }

// GetBookingsActivityReportRequest represents a payload used by the GetBookingsActivityReport method of a ReportService
type GetBookingsActivityReportRequest struct{}

// Validate a GetBookingsActivityReport. Returns a ValidationError for each requirement that fails.
func (r GetBookingsActivityReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetBookingsActivityReportResponse represents a response returned by the GetBookingsActivityReport method of a ReportService.
type GetBookingsActivityReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetBookingsActivityReportResponse) Error() error { return r.Err }

// GetTodaysBookingsReportRequest represents a payload used by the GetTodaysBookingsReport method of a ReportService
type GetTodaysBookingsReportRequest struct{}

// Validate a GetTodaysBookingsReport. Returns a ValidationError for each requirement that fails.
func (r GetTodaysBookingsReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetTodaysBookingsReportResponse represents a response returned by the GetTodaysBookingsReport method of a ReportService.
type GetTodaysBookingsReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetTodaysBookingsReportResponse) Error() error { return r.Err }

// GetTopResourcesReportRequest represents a payload used by the GetTopResourcesReport method of a ReportService
type GetTopResourcesReportRequest struct { // insert request properties here
}

// Validate a GetTopResourcesReport. Returns a ValidationError for each requirement that fails.
func (r GetTopResourcesReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetTopResourcesReportResponse represents a response returned by the GetTopResourcesReport method of a ReportService.
type GetTopResourcesReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetTopResourcesReportResponse) Error() error { return r.Err }

// GetTopEmployeesReportRequest represents a payload used by the GetTopEmployeesReport method of a ReportService
type GetTopEmployeesReportRequest struct { // insert request properties here
}

// Validate a GetTopEmployeesReport. Returns a ValidationError for each requirement that fails.
func (r GetTopEmployeesReportRequest) Validate() []ValidationError {
	// insert validation logic here
	return nil
}

// GetTopEmployeesReportResponse represents a response returned by the GetTopEmployeesReport method of a ReportService.
type GetTopEmployeesReportResponse struct {
	Err error `json:"err,omitempty"`
}

// Error implements the errorer interface. Returns property Err from the response.
func (r GetTopEmployeesReportResponse) Error() error { return r.Err }

// ReportServiceMiddleware defines a middleware for ReportService
type ReportServiceMiddleware func(service ReportService) ReportService

type reportValidationMiddleware struct {
	ReportService
}

// ReportValidationMiddleware returns a middleware for validating requests made to a ReportService
func ReportValidationMiddleware() ReportServiceMiddleware {
	return func(next ReportService) ReportService {
		return reportValidationMiddleware{next}
	}
}

// GetRecentSalesReport validates a GetRecentSalesReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetRecentSalesReport(ctx context.Context, req GetRecentSalesReportRequest) GetRecentSalesReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetRecentSalesReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetRecentSalesReport(ctx, req)
}

// GetUpcomingBookingsReport validates a GetUpcomingBookingsReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetUpcomingBookingsReport(ctx context.Context, req GetUpcomingBookingsReportRequest) GetUpcomingBookingsReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetUpcomingBookingsReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetUpcomingBookingsReport(ctx, req)
}

// GetBookingsActivityReport validates a GetBookingsActivityReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetBookingsActivityReport(ctx context.Context, req GetBookingsActivityReportRequest) GetBookingsActivityReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetBookingsActivityReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetBookingsActivityReport(ctx, req)
}

// GetTodaysBookingsReport validates a GetTodaysBookingsReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetTodaysBookingsReport(ctx context.Context, req GetTodaysBookingsReportRequest) GetTodaysBookingsReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetTodaysBookingsReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetTodaysBookingsReport(ctx, req)
}

// GetTopResourcesReport validates a GetTopResourcesReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetTopResourcesReport(ctx context.Context, req GetTopResourcesReportRequest) GetTopResourcesReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetTopResourcesReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetTopResourcesReport(ctx, req)
}

// GetTopEmployeesReport validates a GetTopEmployeesReportRequest. Returns a domain error if any requirements fail and invokes the next middleware otherwise.
func (mw reportValidationMiddleware) GetTopEmployeesReport(ctx context.Context, req GetTopEmployeesReportRequest) GetTopEmployeesReportResponse {
	errs := req.Validate()
	if len(errs) > 0 {
		return GetTopEmployeesReportResponse{Err: WrapValidationErrors(errs)}
	}
	return mw.ReportService.GetTopEmployeesReport(ctx, req)
}
