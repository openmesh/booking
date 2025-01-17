package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/privacy"
	"github.com/openmesh/booking/ent/rule"
)

// Booking holds the schema definition for the Booking entity.
type Booking struct {
	ent.Schema
}

// Fields of the Booking.
func (Booking) Fields() []ent.Field {
	return []ent.Field{
		field.String("status"),
		field.Time("startTime"),
		field.Time("endTime"),
		field.Int("resourceId"),
	}
}

// Edges of the Booking.
func (Booking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", BookingMetadatum.Type),
		edge.From("resource", Resource.Type).
			Ref("bookings").
			Field("resourceId").
			Unique().
			Required(),
	}
}

// Mixins of the Booking.
func (Booking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Timestamp{},
	}
}

func (Booking) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterBookingOrganizationQueryRule(),
		},
		Mutation: privacy.MutationPolicy{
			rule.FilterBookingOrganizationMutationRule(),
		},
	}
}
