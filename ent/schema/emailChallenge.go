package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Verification holds the schema definition for the Verification entity.
type EmailChallenge struct {
	ent.Schema
}

// func VerificationExpiryGenerator() time.Time {
// 	current := time.Now()
// 	current.Add(time.Hour * 24)
// 	return current
// }

// Fields of the Verification.
func (EmailChallenge) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Bytes("code").Optional(),
	}
}

// Edges of the Verification.
func (EmailChallenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("challenge", Challenge.Type).Ref("emailChallenge").Unique(),
	}
}

// Indexes of the Verification.
func (EmailChallenge) Indexes() []ent.Index {
	return nil
}
