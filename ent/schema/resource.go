package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Int("quantityAvailable"),
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
