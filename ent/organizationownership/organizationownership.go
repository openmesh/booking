// Code generated by entc, DO NOT EDIT.

package organizationownership

const (
	// Label holds the string label denoting the organizationownership type in the database.
	Label = "organization_ownership"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserId holds the string denoting the userid field in the database.
	FieldUserId = "user_id"
	// FieldOrganizationId holds the string denoting the organizationid field in the database.
	FieldOrganizationId = "organization_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// Table holds the table name of the organizationownership in the database.
	Table = "organization_ownerships"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "organization_ownerships"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "organization_ownerships"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
)

// Columns holds all SQL columns for organizationownership fields.
var Columns = []string{
	FieldID,
	FieldUserId,
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
