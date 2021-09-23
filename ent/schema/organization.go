package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("publicKey"),
		field.String("privateKey"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("resources", Resource.Type),
	}
}

// Mixin of the Organization.
func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("privateKey").Unique(),
		index.Fields("publicKey").Unique(),
	}
}
