package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type WebAuthnChallenge struct {
	ent.Schema
}

func (WebAuthnChallenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("sdChallenge").Optional(),
		field.Bytes("userId").Optional(),
		field.JSON("allowedCredentialIds", [][]byte{}).Optional(),
		field.String("userVerification").Optional(),
		field.JSON("extensions", map[string]any{}).Optional(),
	}
}

func (WebAuthnChallenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("challenge", Challenge.Type).Ref("webauthnChallenge").Unique(),
	}
}

func (WebAuthnChallenge) Indexes() []ent.Index {
	return nil
}
