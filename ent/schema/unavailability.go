package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/rule"
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
	}
}

// Mixins of the Unavailability.
func (Unavailability) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}

func (Unavailability) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterUnavailabilityOrganizationQueryRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.FilterResourceOrganizationMutationRule(),
		},
	}
}
