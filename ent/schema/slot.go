package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/rule"
)

// Slot holds the schema definition for the Slot entity.
type Slot struct {
	ent.Schema
}

// Fields of the Slot.
func (Slot) Fields() []ent.Field {
	return []ent.Field{
		field.String("day"),
		field.String("startTime"),
		field.String("endTime"),
		field.Int("quantity").Nillable().Optional(),
		field.Int("resourceId"),
	}
}

// Edges of the Slot.
func (Slot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("resource", Resource.Type).
			Ref("slots").
			Field("resourceId").
			Unique().
			Required(),
	}
}

func (Slot) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterSlotOrganizationQueryRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.FilterSlotOrganizationMutationRule(),
		},
	}
}
