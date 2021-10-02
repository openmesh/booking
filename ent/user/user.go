// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldOrganizationId holds the string denoting the organizationid field in the database.
	FieldOrganizationId = "organization_id"
	// EdgeAuths holds the string denoting the auths edge name in mutations.
	EdgeAuths = "auths"
	// EdgeTokens holds the string denoting the tokens edge name in mutations.
	EdgeTokens = "tokens"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AuthsTable is the table that holds the auths relation/edge.
	AuthsTable = "auths"
	// AuthsInverseTable is the table name for the Auth entity.
	// It exists in this package in order to avoid circular dependency with the "auth" package.
	AuthsInverseTable = "auths"
	// AuthsColumn is the table column denoting the auths relation/edge.
	AuthsColumn = "user_id"
	// TokensTable is the table that holds the tokens relation/edge.
	TokensTable = "tokens"
	// TokensInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokensInverseTable = "tokens"
	// TokensColumn is the table column denoting the tokens relation/edge.
	TokensColumn = "user_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "users"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldEmail,
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
