package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type TotpCredential struct {
	ent.Schema
}

func (TotpCredential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("secret").NotEmpty(),
		field.Bool("validated").Default(false),
	}
}

func (TotpCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("totpCredential").Unique(),
		edge.From("challenge", Challenge.Type).Ref("totpCredential").Unique(),
	}
}

func (TotpCredential) Indexes() []ent.Index {
	return nil
}
