// Code generated by entc, DO NOT EDIT.

package booking

import (
	"time"
)

const (
	// Label holds the string label denoting the booking type in the database.
	Label = "booking"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldResourceId holds the string denoting the resourceid field in the database.
	FieldResourceId = "resource_id"
	// EdgeMetadata holds the string denoting the metadata edge name in mutations.
	EdgeMetadata = "metadata"
	// EdgeResource holds the string denoting the resource edge name in mutations.
	EdgeResource = "resource"
	// Table holds the table name of the booking in the database.
	Table = "bookings"
	// MetadataTable is the table that holds the metadata relation/edge.
	MetadataTable = "booking_metadata"
	// MetadataInverseTable is the table name for the BookingMetadatum entity.
	// It exists in this package in order to avoid circular dependency with the "bookingmetadatum" package.
	MetadataInverseTable = "booking_metadata"
	// MetadataColumn is the table column denoting the metadata relation/edge.
	MetadataColumn = "booking_id"
	// ResourceTable is the table that holds the resource relation/edge.
	ResourceTable = "bookings"
	// ResourceInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourceInverseTable = "resources"
	// ResourceColumn is the table column denoting the resource relation/edge.
	ResourceColumn = "resource_id"
)

// Columns holds all SQL columns for booking fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldStartTime,
	FieldEndTime,
	FieldResourceId,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updatedAt" field.
	UpdateDefaultUpdatedAt func() time.Time
)
