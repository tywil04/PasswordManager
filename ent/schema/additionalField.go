package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for Additional Fields entity.
type AdditionalField struct {
	ent.Schema
}

// Fields of Additional Fields.
func (AdditionalField) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`form:"id" json:"id" xml:"id"`),
		field.Bytes("key").NotEmpty().StructTag(`form:"key" json:"key" xml:"key"`),
		field.Bytes("keyIv").NotEmpty().StructTag(`form:"keyIv" json:"keyIv" xml:"keyIv"`),
		field.Bytes("value").NotEmpty().StructTag(`form:"value" json:"value" xml:"value"`),
		field.Bytes("valueIv").NotEmpty().StructTag(`form:"valueIv" json:"valueIv" xml:"valueIv"`),
	}
}

// Edges of Additional Fields.
func (AdditionalField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("password", Password.Type).
			Ref("additionalFields").
			Unique(),
	}
}

// Indexes of Additional Fields.
func (AdditionalField) Indexes() []ent.Index {
	return []ent.Index{}
}
