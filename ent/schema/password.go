package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for Passwords entity.
type Password struct {
	ent.Schema
}

// Fields of Passwords.
func (Password) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StructTag(`form:"id" json:"id" xml:"id"`),
		field.Bytes("name").NotEmpty(),
		field.Bytes("nameIv").NotEmpty(),
		field.Bytes("username").NotEmpty(),
		field.Bytes("usernameIv").NotEmpty(),
		field.Bytes("password").NotEmpty(),
		field.Bytes("passwordIv").NotEmpty(),
		field.Bytes("colour").NotEmpty(),
		field.Bytes("colourIv").NotEmpty(),
	}
}

// Edges of Passwords.
func (Password) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("additionalFields", AdditionalField.Type),
		edge.To("urls", Url.Type),
		edge.From("vault", Vault.Type).
			Ref("passwords").
			Unique(),
	}
}

// Indexes of Passwords.
func (Password) Indexes() []ent.Index {
	return []ent.Index{}
}
