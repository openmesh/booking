package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BookingMetadatum holds the schema definition for the BookingMetadatum entity.
type BookingMetadatum struct {
	ent.Schema
}

// Fields of the BookingMetadatum.
func (BookingMetadatum) Fields() []ent.Field {
	return []ent.Field{
		field.String("key"),
		field.String("value"),
		field.Int("bookingId"),
	}
}

// Edges of the BookingMetadatum.
func (BookingMetadatum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("booking", Booking.Type).
			Ref("metadata").
			Field("bookingId").
			Unique().
			Required(),
	}
}
