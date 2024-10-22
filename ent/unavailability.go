// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/unavailability"
)

// Unavailability is the model entity for the Unavailability schema.
type Unavailability struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// StartTime holds the value of the "startTime" field.
	StartTime time.Time `json:"startTime,omitempty"`
	// EndTime holds the value of the "endTime" field.
	EndTime time.Time `json:"endTime,omitempty"`
	// ResourceId holds the value of the "resourceId" field.
	ResourceId int `json:"resourceId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UnavailabilityQuery when eager-loading is set.
	Edges UnavailabilityEdges `json:"edges"`
}

// UnavailabilityEdges holds the relations/edges for other nodes in the graph.
type UnavailabilityEdges struct {
	// Resource holds the value of the resource edge.
	Resource *Resource `json:"resource,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ResourceOrErr returns the Resource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UnavailabilityEdges) ResourceOrErr() (*Resource, error) {
	if e.loadedTypes[0] {
		if e.Resource == nil {
			// The edge resource was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: resource.Label}
		}
		return e.Resource, nil
	}
	return nil, &NotLoadedError{edge: "resource"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Unavailability) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case unavailability.FieldID, unavailability.FieldResourceId:
			values[i] = new(sql.NullInt64)
		case unavailability.FieldCreatedAt, unavailability.FieldUpdatedAt, unavailability.FieldStartTime, unavailability.FieldEndTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Unavailability", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Unavailability fields.
func (u *Unavailability) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case unavailability.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case unavailability.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case unavailability.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case unavailability.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field startTime", values[i])
			} else if value.Valid {
				u.StartTime = value.Time
			}
		case unavailability.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field endTime", values[i])
			} else if value.Valid {
				u.EndTime = value.Time
			}
		case unavailability.FieldResourceId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field resourceId", values[i])
			} else if value.Valid {
				u.ResourceId = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryResource queries the "resource" edge of the Unavailability entity.
func (u *Unavailability) QueryResource() *ResourceQuery {
	return (&UnavailabilityClient{config: u.config}).QueryResource(u)
}

// Update returns a builder for updating this Unavailability.
// Note that you need to call Unavailability.Unwrap() before calling this method if this Unavailability
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Unavailability) Update() *UnavailabilityUpdateOne {
	return (&UnavailabilityClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Unavailability entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Unavailability) Unwrap() *Unavailability {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Unavailability is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Unavailability) String() string {
	var builder strings.Builder
	builder.WriteString("Unavailability(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", createdAt=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", startTime=")
	builder.WriteString(u.StartTime.Format(time.ANSIC))
	builder.WriteString(", endTime=")
	builder.WriteString(u.EndTime.Format(time.ANSIC))
	builder.WriteString(", resourceId=")
	builder.WriteString(fmt.Sprintf("%v", u.ResourceId))
	builder.WriteByte(')')
	return builder.String()
}

// Unavailabilities is a parsable slice of Unavailability.
type Unavailabilities []*Unavailability

func (u Unavailabilities) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
