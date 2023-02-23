package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Vault struct {
	ent.Schema
}

// Fields of the User.
func (Vault) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("createdAt").Default(time.Now),
		field.String("name").NotEmpty(),
		field.String("colour"),
	}
}

// Edges of the User.
func (Vault) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("passwords", Password.Type),
		edge.From("user", User.Type).
			Ref("vaults").
			Unique(),
	}
}

// Indexes of the User.
func (Vault) Indexes() []ent.Index {
	return nil
}
