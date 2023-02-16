package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Challenge struct {
	ent.Schema
}

func (Challenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("expiry").Default(func() time.Time {
			return time.Now().Add(time.Hour)
		}),
	}
}

func (Challenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("challenges").Unique().Required(),
		edge.To("emailChallenge", EmailChallenge.Type).Unique(),
		edge.To("webauthnChallenge", WebAuthnChallenge.Type).Unique(),
		edge.To("totpCredential", TotpCredential.Type).Unique(),
	}
}

func (Challenge) Indexes() []ent.Index {
	return nil
}
