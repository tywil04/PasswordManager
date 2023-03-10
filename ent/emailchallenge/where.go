// Code generated by ent, DO NOT EDIT.

package emailchallenge

import (
	"PasswordManager/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldEQ(FieldCode, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...[]byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...[]byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v []byte) predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldLTE(FieldCode, v))
}

// CodeIsNil applies the IsNil predicate on the "code" field.
func CodeIsNil() predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldIsNull(FieldCode))
}

// CodeNotNil applies the NotNil predicate on the "code" field.
func CodeNotNil() predicate.EmailChallenge {
	return predicate.EmailChallenge(sql.FieldNotNull(FieldCode))
}

// HasChallenge applies the HasEdge predicate on the "challenge" edge.
func HasChallenge() predicate.EmailChallenge {
	return predicate.EmailChallenge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChallengeTable, ChallengeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChallengeWith applies the HasEdge predicate on the "challenge" edge with a given conditions (other predicates).
func HasChallengeWith(preds ...predicate.Challenge) predicate.EmailChallenge {
	return predicate.EmailChallenge(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChallengeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChallengeTable, ChallengeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EmailChallenge) predicate.EmailChallenge {
	return predicate.EmailChallenge(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EmailChallenge) predicate.EmailChallenge {
	return predicate.EmailChallenge(func(s *sql.Selector) {
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
func Not(p predicate.EmailChallenge) predicate.EmailChallenge {
	return predicate.EmailChallenge(func(s *sql.Selector) {
		p(s.Not())
	})
}
