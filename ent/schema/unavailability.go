package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Unavailability holds the schema definition for the Unavailability entity.
type Unavailability struct {
	ent.Schema
}

// Fields of the Unavailability.
func (Unavailability) Fields() []ent.Field {
	return []ent.Field{
		field.Time("startTime"),
		field.Time("endTime"),
		field.Int("resourceId"),
		field.Int("organizationId"),
	}
}

// Edges of the Unavailability.
func (Unavailability) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resource", Resource.Type).
			Ref("unavailabilities").
			Field("resourceId").
			Unique().
			Required(),
		edge.From("organization", Organization.Type).
			Ref("unavailabilities").
			Field("organizationId").
			Unique().
			Required(),
	}
}

// Mixins of the Unavailability.
func (Unavailability) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}
