package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("email").Unique(),
		field.Bytes("strengthenedMasterHash").NotEmpty().Sensitive(),
		field.Bytes("strengthenedMasterHashSalt").NotEmpty().Sensitive(),
		field.Bytes("protectedDatabaseKey").NotEmpty().Sensitive(),
		field.Bytes("protectedDatabaseKeyIv").NotEmpty().Sensitive(),
		field.Bool("webauthnEnabled").Default(false),
		field.Bool("totpEnabled").Default(false),
		field.Bool("verified").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("totpCredential", TotpCredential.Type).Unique(),
		edge.To("webauthnCredentials", WebAuthnCredential.Type),
		edge.To("webauthnRegisterChallenges", WebAuthnRegisterChallenge.Type),
		edge.To("vaults", Vault.Type),
		edge.To("sessions", Session.Type),
		edge.To("challenges", Challenge.Type),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Unique(),
	}
}
