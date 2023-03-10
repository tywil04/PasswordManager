// Code generated by ent, DO NOT EDIT.

package session

import (
	"PasswordManager/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldID, id))
}

// N applies equality check predicate on the "n" field. It's identical to NEQ.
func N(v []byte) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldN, v))
}

// E applies equality check predicate on the "e" field. It's identical to EEQ.
func E(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldE, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiry, v))
}

// NEQ applies the EQ predicate on the "n" field.
func NEQ(v []byte) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldN, v))
}

// NNEQ applies the NEQ predicate on the "n" field.
func NNEQ(v []byte) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldN, v))
}

// NIn applies the In predicate on the "n" field.
func NIn(vs ...[]byte) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldN, vs...))
}

// NNotIn applies the NotIn predicate on the "n" field.
func NNotIn(vs ...[]byte) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldN, vs...))
}

// NGT applies the GT predicate on the "n" field.
func NGT(v []byte) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldN, v))
}

// NGTE applies the GTE predicate on the "n" field.
func NGTE(v []byte) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldN, v))
}

// NLT applies the LT predicate on the "n" field.
func NLT(v []byte) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldN, v))
}

// NLTE applies the LTE predicate on the "n" field.
func NLTE(v []byte) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldN, v))
}

// EEQ applies the EQ predicate on the "e" field.
func EEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldE, v))
}

// ENEQ applies the NEQ predicate on the "e" field.
func ENEQ(v int) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldE, v))
}

// EIn applies the In predicate on the "e" field.
func EIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldE, vs...))
}

// ENotIn applies the NotIn predicate on the "e" field.
func ENotIn(vs ...int) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldE, vs...))
}

// EGT applies the GT predicate on the "e" field.
func EGT(v int) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldE, v))
}

// EGTE applies the GTE predicate on the "e" field.
func EGTE(v int) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldE, v))
}

// ELT applies the LT predicate on the "e" field.
func ELT(v int) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldE, v))
}

// ELTE applies the LTE predicate on the "e" field.
func ELTE(v int) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldE, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldExpiry, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
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
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
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
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(func(s *sql.Selector) {
		p(s.Not())
	})
}
