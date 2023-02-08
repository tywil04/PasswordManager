package schema

import (
	"time"

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
		field.String("code").NotEmpty(),
		field.Time("expiry").Default(time.Now),
		field.Enum("for").Values("signup", "signin"),
	}
}

// Edges of the Verification.
func (EmailChallenge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("emailChallenges").Unique().Required(),
	}
}

// Indexes of the Verification.
func (EmailChallenge) Indexes() []ent.Index {
	return nil
}
