package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for Additional Fields entity.
type Session struct {
	ent.Schema
}

// Fields of Additional Fields.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bytes("n").NotEmpty().Sensitive(),
		field.Int("e"),
		field.Time("expiry"),
	}
}

// Edges of Additional Fields.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("sessions").
			Unique(),
	}
}

// Indexes of Additional Fields.
func (Session) Indexes() []ent.Index {
	return []ent.Index{}
}
