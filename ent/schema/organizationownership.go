package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OrganizationOwnership holds the schema definition for the OrganizationOwnership entity.
type OrganizationOwnership struct {
	ent.Schema
}

// Fields of the OrganizationOwnership.
func (OrganizationOwnership) Fields() []ent.Field {
	return []ent.Field{
		field.Int("userId"),
		field.Int("organizationId"),
	}
}

// Edges of the OrganizationOwnership.
func (OrganizationOwnership) Edges() []ent.Edge {
	return []ent.Edge {
		edge.To("user", User.Type).Field("userId").Unique().Required(),
		edge.To("organization", Organization.Type).Field("organizationId").Unique().Required(),
	}
}

// Indexes of the OrganizationOwnership.
func (OrganizationOwnership) Indexes() []ent.Index {
	return []ent.Index {
		index.Fields("userId", "organizationId").Unique(),
	}
}