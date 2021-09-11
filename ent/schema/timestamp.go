package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Timestamp struct {
	mixin.Schema
}

// Fields of the Timestamp.
func (Timestamp) Fields() []ent.Field {
	return []ent.Field{
		field.Time("createdAt").Default(time.Now).Immutable(),
		field.Time("updatedAt").Default(time.Now).UpdateDefault(time.Now),
	}
}
