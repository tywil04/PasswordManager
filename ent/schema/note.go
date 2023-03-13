package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for Passwords entity.
type Note struct {
	ent.Schema
}

// Fields of Passwords.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bytes("name").NotEmpty(),
		field.Bytes("nameIv").NotEmpty(),
		field.Bytes("title").NotEmpty(),
		field.Bytes("titleIv").NotEmpty(),
		field.Bytes("content").NotEmpty(),
		field.Bytes("contentIv").NotEmpty(),
		field.Bytes("colour").NotEmpty(),
		field.Bytes("colourIv").NotEmpty(),
	}
}

// Edges of Passwords.
func (Note) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("vault", Vault.Type).
			Ref("notes").
			Unique(),
	}
}

// Indexes of Passwords.
func (Note) Indexes() []ent.Index {
	return []ent.Index{}
}
