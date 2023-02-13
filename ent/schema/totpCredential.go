package schema

import (
	"time"

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
		field.Time("createdAt").Default(time.Now),
		field.String("secret").NotEmpty(),
		field.Bool("validated").Default(false),
	}
}

func (TotpCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("totpCredential").Unique().Required(),
	}
}

func (TotpCredential) Indexes() []ent.Index {
	return nil
}
