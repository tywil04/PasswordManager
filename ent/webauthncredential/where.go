// Code generated by ent, DO NOT EDIT.

package webauthncredential

import (
	"PasswordManager/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCreatedAt, v))
}

// CredentialId applies equality check predicate on the "credentialId" field. It's identical to CredentialIdEQ.
func CredentialId(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCredentialId, v))
}

// PublicKey applies equality check predicate on the "publicKey" field. It's identical to PublicKeyEQ.
func PublicKey(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldPublicKey, v))
}

// AttestationType applies equality check predicate on the "attestationType" field. It's identical to AttestationTypeEQ.
func AttestationType(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldAttestationType, v))
}

// Aaguid applies equality check predicate on the "aaguid" field. It's identical to AaguidEQ.
func Aaguid(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldAaguid, v))
}

// SignCount applies equality check predicate on the "signCount" field. It's identical to SignCountEQ.
func SignCount(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldSignCount, v))
}

// CloneWarning applies equality check predicate on the "cloneWarning" field. It's identical to CloneWarningEQ.
func CloneWarning(v bool) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCloneWarning, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldCreatedAt, v))
}

// CredentialIdEQ applies the EQ predicate on the "credentialId" field.
func CredentialIdEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCredentialId, v))
}

// CredentialIdNEQ applies the NEQ predicate on the "credentialId" field.
func CredentialIdNEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldCredentialId, v))
}

// CredentialIdIn applies the In predicate on the "credentialId" field.
func CredentialIdIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldCredentialId, vs...))
}

// CredentialIdNotIn applies the NotIn predicate on the "credentialId" field.
func CredentialIdNotIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldCredentialId, vs...))
}

// CredentialIdGT applies the GT predicate on the "credentialId" field.
func CredentialIdGT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldCredentialId, v))
}

// CredentialIdGTE applies the GTE predicate on the "credentialId" field.
func CredentialIdGTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldCredentialId, v))
}

// CredentialIdLT applies the LT predicate on the "credentialId" field.
func CredentialIdLT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldCredentialId, v))
}

// CredentialIdLTE applies the LTE predicate on the "credentialId" field.
func CredentialIdLTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldCredentialId, v))
}

// PublicKeyEQ applies the EQ predicate on the "publicKey" field.
func PublicKeyEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldPublicKey, v))
}

// PublicKeyNEQ applies the NEQ predicate on the "publicKey" field.
func PublicKeyNEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldPublicKey, v))
}

// PublicKeyIn applies the In predicate on the "publicKey" field.
func PublicKeyIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldPublicKey, vs...))
}

// PublicKeyNotIn applies the NotIn predicate on the "publicKey" field.
func PublicKeyNotIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldPublicKey, vs...))
}

// PublicKeyGT applies the GT predicate on the "publicKey" field.
func PublicKeyGT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldPublicKey, v))
}

// PublicKeyGTE applies the GTE predicate on the "publicKey" field.
func PublicKeyGTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldPublicKey, v))
}

// PublicKeyLT applies the LT predicate on the "publicKey" field.
func PublicKeyLT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldPublicKey, v))
}

// PublicKeyLTE applies the LTE predicate on the "publicKey" field.
func PublicKeyLTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldPublicKey, v))
}

// AttestationTypeEQ applies the EQ predicate on the "attestationType" field.
func AttestationTypeEQ(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldAttestationType, v))
}

// AttestationTypeNEQ applies the NEQ predicate on the "attestationType" field.
func AttestationTypeNEQ(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldAttestationType, v))
}

// AttestationTypeIn applies the In predicate on the "attestationType" field.
func AttestationTypeIn(vs ...string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldAttestationType, vs...))
}

// AttestationTypeNotIn applies the NotIn predicate on the "attestationType" field.
func AttestationTypeNotIn(vs ...string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldAttestationType, vs...))
}

// AttestationTypeGT applies the GT predicate on the "attestationType" field.
func AttestationTypeGT(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldAttestationType, v))
}

// AttestationTypeGTE applies the GTE predicate on the "attestationType" field.
func AttestationTypeGTE(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldAttestationType, v))
}

// AttestationTypeLT applies the LT predicate on the "attestationType" field.
func AttestationTypeLT(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldAttestationType, v))
}

// AttestationTypeLTE applies the LTE predicate on the "attestationType" field.
func AttestationTypeLTE(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldAttestationType, v))
}

// AttestationTypeContains applies the Contains predicate on the "attestationType" field.
func AttestationTypeContains(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldContains(FieldAttestationType, v))
}

// AttestationTypeHasPrefix applies the HasPrefix predicate on the "attestationType" field.
func AttestationTypeHasPrefix(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldHasPrefix(FieldAttestationType, v))
}

// AttestationTypeHasSuffix applies the HasSuffix predicate on the "attestationType" field.
func AttestationTypeHasSuffix(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldHasSuffix(FieldAttestationType, v))
}

// AttestationTypeEqualFold applies the EqualFold predicate on the "attestationType" field.
func AttestationTypeEqualFold(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEqualFold(FieldAttestationType, v))
}

// AttestationTypeContainsFold applies the ContainsFold predicate on the "attestationType" field.
func AttestationTypeContainsFold(v string) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldContainsFold(FieldAttestationType, v))
}

// AaguidEQ applies the EQ predicate on the "aaguid" field.
func AaguidEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldAaguid, v))
}

// AaguidNEQ applies the NEQ predicate on the "aaguid" field.
func AaguidNEQ(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldAaguid, v))
}

// AaguidIn applies the In predicate on the "aaguid" field.
func AaguidIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldAaguid, vs...))
}

// AaguidNotIn applies the NotIn predicate on the "aaguid" field.
func AaguidNotIn(vs ...[]byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldAaguid, vs...))
}

// AaguidGT applies the GT predicate on the "aaguid" field.
func AaguidGT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldAaguid, v))
}

// AaguidGTE applies the GTE predicate on the "aaguid" field.
func AaguidGTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldAaguid, v))
}

// AaguidLT applies the LT predicate on the "aaguid" field.
func AaguidLT(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldAaguid, v))
}

// AaguidLTE applies the LTE predicate on the "aaguid" field.
func AaguidLTE(v []byte) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldAaguid, v))
}

// SignCountEQ applies the EQ predicate on the "signCount" field.
func SignCountEQ(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldSignCount, v))
}

// SignCountNEQ applies the NEQ predicate on the "signCount" field.
func SignCountNEQ(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldSignCount, v))
}

// SignCountIn applies the In predicate on the "signCount" field.
func SignCountIn(vs ...uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldIn(FieldSignCount, vs...))
}

// SignCountNotIn applies the NotIn predicate on the "signCount" field.
func SignCountNotIn(vs ...uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNotIn(FieldSignCount, vs...))
}

// SignCountGT applies the GT predicate on the "signCount" field.
func SignCountGT(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGT(FieldSignCount, v))
}

// SignCountGTE applies the GTE predicate on the "signCount" field.
func SignCountGTE(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldGTE(FieldSignCount, v))
}

// SignCountLT applies the LT predicate on the "signCount" field.
func SignCountLT(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLT(FieldSignCount, v))
}

// SignCountLTE applies the LTE predicate on the "signCount" field.
func SignCountLTE(v uint32) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldLTE(FieldSignCount, v))
}

// CloneWarningEQ applies the EQ predicate on the "cloneWarning" field.
func CloneWarningEQ(v bool) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldEQ(FieldCloneWarning, v))
}

// CloneWarningNEQ applies the NEQ predicate on the "cloneWarning" field.
func CloneWarningNEQ(v bool) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(sql.FieldNEQ(FieldCloneWarning, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.WebAuthnCredential) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.WebAuthnCredential) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(func(s *sql.Selector) {
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
func Not(p predicate.WebAuthnCredential) predicate.WebAuthnCredential {
	return predicate.WebAuthnCredential(func(s *sql.Selector) {
		p(s.Not())
	})
}