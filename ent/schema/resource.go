package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/rule"
	// "github.com/openmesh/booking/ent/privacy"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.String("timezone"),
		field.String("password"),
		field.Int("price"),
		field.Int("bookingPrice"),
		field.Int("organizationId"),
		field.Int("quantityAvailable").
			Optional().
			Nillable(),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("slots", Slot.Type),
		edge.To("bookings", Booking.Type),
		edge.To("unavailabilities", Unavailability.Type),
		edge.From("organization", Organization.Type).
			Ref("resources").
			Field("organizationId").
			Unique().
			Required(),
	}
}

// Mixins of the Resource.
func (Resource) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}

func (Resource) Policy() ent.Policy {
	return privacy.Policy{
		// Mutation: privacy.MutationPolicy{
		// 	rule.
		// },
		Query: privacy.QueryPolicy{
			rule.FilterOrganizationRuleResource(),
		},
	}
}
