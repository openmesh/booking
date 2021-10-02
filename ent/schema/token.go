package schema

import (
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/rule"
	"github.com/openmesh/booking/rand"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Default(fmt.Sprintf("omb_%s", rand.Key())).
			Immutable(),
		field.String("name").
			Immutable(),
		field.Time("expiry").
			Immutable().
			Nillable().
			Optional(),
		field.Int("userId"),
		field.Int("organizationId"),
	}
}

func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tokens").
			Field("userId").
			Unique().
			Required(),
		edge.From("organization", Organization.Type).
			Ref("tokens").
			Field("organizationId").
			Unique().
			Required(),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}

func (Token) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.FilterTokenUserMutationRule(),
			rule.FilterTokenOrganizationMutationRule(),
		},
	}
}
