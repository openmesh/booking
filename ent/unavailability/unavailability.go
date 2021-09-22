// Code generated by entc, DO NOT EDIT.

package unavailability

import (
	"time"
)

const (
	// Label holds the string label denoting the unavailability type in the database.
	Label = "unavailability"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStartTime holds the string denoting the starttime field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the endtime field in the database.
	FieldEndTime = "end_time"
	// FieldResourceId holds the string denoting the resourceid field in the database.
	FieldResourceId = "resource_id"
	// FieldOrganizationId holds the string denoting the organizationid field in the database.
	FieldOrganizationId = "organization_id"
	// EdgeResource holds the string denoting the resource edge name in mutations.
	EdgeResource = "resource"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the unavailability in the database.
	Table = "unavailabilities"
	// ResourceTable is the table that holds the resource relation/edge.
	ResourceTable = "unavailabilities"
	// ResourceInverseTable is the table name for the Resource entity.
	// It exists in this package in order to avoid circular dependency with the "resource" package.
	ResourceInverseTable = "resources"
	// ResourceColumn is the table column denoting the resource relation/edge.
	ResourceColumn = "resource_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "unavailabilities"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
)

// Columns holds all SQL columns for unavailability fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStartTime,
	FieldEndTime,
	FieldResourceId,
	FieldOrganizationId,
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
