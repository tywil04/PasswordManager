// Code generated by ent, DO NOT EDIT.

package user

import (
	"PasswordManager/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// StrengthenedMasterHash applies equality check predicate on the "strengthenedMasterHash" field. It's identical to StrengthenedMasterHashEQ.
func StrengthenedMasterHash(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashSalt applies equality check predicate on the "strengthenedMasterHashSalt" field. It's identical to StrengthenedMasterHashSaltEQ.
func StrengthenedMasterHashSalt(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStrengthenedMasterHashSalt, v))
}

// ProtectedDatabaseKey applies equality check predicate on the "protectedDatabaseKey" field. It's identical to ProtectedDatabaseKeyEQ.
func ProtectedDatabaseKey(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyIv applies equality check predicate on the "protectedDatabaseKeyIv" field. It's identical to ProtectedDatabaseKeyIvEQ.
func ProtectedDatabaseKeyIv(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProtectedDatabaseKeyIv, v))
}

// Verified applies equality check predicate on the "verified" field. It's identical to VerifiedEQ.
func Verified(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVerified, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// StrengthenedMasterHashEQ applies the EQ predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashNEQ applies the NEQ predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashIn applies the In predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldStrengthenedMasterHash, vs...))
}

// StrengthenedMasterHashNotIn applies the NotIn predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStrengthenedMasterHash, vs...))
}

// StrengthenedMasterHashGT applies the GT predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashGTE applies the GTE predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashLT applies the LT predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashLTE applies the LTE predicate on the "strengthenedMasterHash" field.
func StrengthenedMasterHashLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStrengthenedMasterHash, v))
}

// StrengthenedMasterHashSaltEQ applies the EQ predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldStrengthenedMasterHashSalt, v))
}

// StrengthenedMasterHashSaltNEQ applies the NEQ predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldStrengthenedMasterHashSalt, v))
}

// StrengthenedMasterHashSaltIn applies the In predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldStrengthenedMasterHashSalt, vs...))
}

// StrengthenedMasterHashSaltNotIn applies the NotIn predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldStrengthenedMasterHashSalt, vs...))
}

// StrengthenedMasterHashSaltGT applies the GT predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldStrengthenedMasterHashSalt, v))
}

// StrengthenedMasterHashSaltGTE applies the GTE predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldStrengthenedMasterHashSalt, v))
}

// StrengthenedMasterHashSaltLT applies the LT predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldStrengthenedMasterHashSalt, v))
}

// StrengthenedMasterHashSaltLTE applies the LTE predicate on the "strengthenedMasterHashSalt" field.
func StrengthenedMasterHashSaltLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldStrengthenedMasterHashSalt, v))
}

// ProtectedDatabaseKeyEQ applies the EQ predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyNEQ applies the NEQ predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyIn applies the In predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldProtectedDatabaseKey, vs...))
}

// ProtectedDatabaseKeyNotIn applies the NotIn predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProtectedDatabaseKey, vs...))
}

// ProtectedDatabaseKeyGT applies the GT predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyGTE applies the GTE predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyLT applies the LT predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyLTE applies the LTE predicate on the "protectedDatabaseKey" field.
func ProtectedDatabaseKeyLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProtectedDatabaseKey, v))
}

// ProtectedDatabaseKeyIvEQ applies the EQ predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProtectedDatabaseKeyIv, v))
}

// ProtectedDatabaseKeyIvNEQ applies the NEQ predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvNEQ(v []byte) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProtectedDatabaseKeyIv, v))
}

// ProtectedDatabaseKeyIvIn applies the In predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldIn(FieldProtectedDatabaseKeyIv, vs...))
}

// ProtectedDatabaseKeyIvNotIn applies the NotIn predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvNotIn(vs ...[]byte) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProtectedDatabaseKeyIv, vs...))
}

// ProtectedDatabaseKeyIvGT applies the GT predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvGT(v []byte) predicate.User {
	return predicate.User(sql.FieldGT(FieldProtectedDatabaseKeyIv, v))
}

// ProtectedDatabaseKeyIvGTE applies the GTE predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvGTE(v []byte) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProtectedDatabaseKeyIv, v))
}

// ProtectedDatabaseKeyIvLT applies the LT predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvLT(v []byte) predicate.User {
	return predicate.User(sql.FieldLT(FieldProtectedDatabaseKeyIv, v))
}

// ProtectedDatabaseKeyIvLTE applies the LTE predicate on the "protectedDatabaseKeyIv" field.
func ProtectedDatabaseKeyIvLTE(v []byte) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProtectedDatabaseKeyIv, v))
}

// Default2FAEQ applies the EQ predicate on the "default2FA" field.
func Default2FAEQ(v Default2FA) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDefault2FA, v))
}

// Default2FANEQ applies the NEQ predicate on the "default2FA" field.
func Default2FANEQ(v Default2FA) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDefault2FA, v))
}

// Default2FAIn applies the In predicate on the "default2FA" field.
func Default2FAIn(vs ...Default2FA) predicate.User {
	return predicate.User(sql.FieldIn(FieldDefault2FA, vs...))
}

// Default2FANotIn applies the NotIn predicate on the "default2FA" field.
func Default2FANotIn(vs ...Default2FA) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDefault2FA, vs...))
}

// VerifiedEQ applies the EQ predicate on the "verified" field.
func VerifiedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldVerified, v))
}

// VerifiedNEQ applies the NEQ predicate on the "verified" field.
func VerifiedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldVerified, v))
}

// HasEmailChallenges applies the HasEdge predicate on the "emailChallenges" edge.
func HasEmailChallenges() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmailChallengesTable, EmailChallengesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEmailChallengesWith applies the HasEdge predicate on the "emailChallenges" edge with a given conditions (other predicates).
func HasEmailChallengesWith(preds ...predicate.EmailChallenge) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EmailChallengesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EmailChallengesTable, EmailChallengesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWebauthnCredentials applies the HasEdge predicate on the "webauthnCredentials" edge.
func HasWebauthnCredentials() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebauthnCredentialsTable, WebauthnCredentialsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebauthnCredentialsWith applies the HasEdge predicate on the "webauthnCredentials" edge with a given conditions (other predicates).
func HasWebauthnCredentialsWith(preds ...predicate.WebAuthnCredential) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WebauthnCredentialsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebauthnCredentialsTable, WebauthnCredentialsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasWebauthnChallenges applies the HasEdge predicate on the "webauthnChallenges" edge.
func HasWebauthnChallenges() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebauthnChallengesTable, WebauthnChallengesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWebauthnChallengesWith applies the HasEdge predicate on the "webauthnChallenges" edge with a given conditions (other predicates).
func HasWebauthnChallengesWith(preds ...predicate.WebAuthnChallenge) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WebauthnChallengesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WebauthnChallengesTable, WebauthnChallengesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPasswords applies the HasEdge predicate on the "passwords" edge.
func HasPasswords() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PasswordsTable, PasswordsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPasswordsWith applies the HasEdge predicate on the "passwords" edge with a given conditions (other predicates).
func HasPasswordsWith(preds ...predicate.Password) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PasswordsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PasswordsTable, PasswordsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSessions applies the HasEdge predicate on the "sessions" edge.
func HasSessions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionsWith applies the HasEdge predicate on the "sessions" edge with a given conditions (other predicates).
func HasSessionsWith(preds ...predicate.Session) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SessionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SessionsTable, SessionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
