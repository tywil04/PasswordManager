package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type WebAuthnCredential struct {
	ent.Schema
}

func (WebAuthnCredential) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name"),
		field.Time("createdAt").Default(time.Now),
		field.Bytes("credentialId"),
		field.Bytes("publicKey"),
		field.String("attestationType"),
		field.Strings("transport"),
		field.Bytes("aaguid"),
		field.Uint32("signCount"),
		field.Bool("cloneWarning"),
	}
}

func (WebAuthnCredential) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("webauthnCredentials").Unique().Required(),
	}
}

func (WebAuthnCredential) Indexes() []ent.Index {
	return nil
}
