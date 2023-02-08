package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for Additional Fields entity.
type Url struct {
	ent.Schema
}

// Fields of Additional Fields.
func (Url) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`form:"id" json:"id" xml:"id"`),
		field.Bytes("url").NotEmpty(),
		field.Bytes("urlIv").NotEmpty(),
	}
}

// Edges of Additional Fields.
func (Url) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("password", Password.Type).
			Ref("urls").
			Unique(),
	}
}

// Indexes of Additional Fields.
func (Url) Indexes() []ent.Index {
	return []ent.Index{}
}
