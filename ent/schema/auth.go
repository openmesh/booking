package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		field.String("source"),
		field.String("sourceId"),
		field.String("accessToken").Nillable().Optional(),
		field.String("refreshToken").Nillable().Optional(),
		field.Time("expiry").Nillable().Optional(),
		field.Int("userId"),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auths").
			Field("userId").
			Unique().
			Required(),
	}
}

// Mixins of the Auth.
func (Auth) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}
