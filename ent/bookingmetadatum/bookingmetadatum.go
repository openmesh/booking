// Code generated by entc, DO NOT EDIT.

package bookingmetadatum

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the bookingmetadatum type in the database.
	Label = "booking_metadatum"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldBookingId holds the string denoting the bookingid field in the database.
	FieldBookingId = "booking_id"
	// EdgeBooking holds the string denoting the booking edge name in mutations.
	EdgeBooking = "booking"
	// Table holds the table name of the bookingmetadatum in the database.
	Table = "booking_metadata"
	// BookingTable is the table that holds the booking relation/edge.
	BookingTable = "booking_metadata"
	// BookingInverseTable is the table name for the Booking entity.
	// It exists in this package in order to avoid circular dependency with the "booking" package.
	BookingInverseTable = "bookings"
	// BookingColumn is the table column denoting the booking relation/edge.
	BookingColumn = "booking_id"
)

// Columns holds all SQL columns for bookingmetadatum fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldValue,
	FieldBookingId,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/openmesh/booking/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
)
