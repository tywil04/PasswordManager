// Code generated by ent, DO NOT EDIT.

package url

import (
	"PasswordManager/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldID, id))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v []byte) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// UrlIv applies equality check predicate on the "urlIv" field. It's identical to UrlIvEQ.
func UrlIv(v []byte) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldUrlIv, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v []byte) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v []byte) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...[]byte) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...[]byte) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v []byte) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v []byte) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v []byte) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v []byte) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldURL, v))
}

// UrlIvEQ applies the EQ predicate on the "urlIv" field.
func UrlIvEQ(v []byte) predicate.Url {
	return predicate.Url(sql.FieldEQ(FieldUrlIv, v))
}

// UrlIvNEQ applies the NEQ predicate on the "urlIv" field.
func UrlIvNEQ(v []byte) predicate.Url {
	return predicate.Url(sql.FieldNEQ(FieldUrlIv, v))
}

// UrlIvIn applies the In predicate on the "urlIv" field.
func UrlIvIn(vs ...[]byte) predicate.Url {
	return predicate.Url(sql.FieldIn(FieldUrlIv, vs...))
}

// UrlIvNotIn applies the NotIn predicate on the "urlIv" field.
func UrlIvNotIn(vs ...[]byte) predicate.Url {
	return predicate.Url(sql.FieldNotIn(FieldUrlIv, vs...))
}

// UrlIvGT applies the GT predicate on the "urlIv" field.
func UrlIvGT(v []byte) predicate.Url {
	return predicate.Url(sql.FieldGT(FieldUrlIv, v))
}

// UrlIvGTE applies the GTE predicate on the "urlIv" field.
func UrlIvGTE(v []byte) predicate.Url {
	return predicate.Url(sql.FieldGTE(FieldUrlIv, v))
}

// UrlIvLT applies the LT predicate on the "urlIv" field.
func UrlIvLT(v []byte) predicate.Url {
	return predicate.Url(sql.FieldLT(FieldUrlIv, v))
}

// UrlIvLTE applies the LTE predicate on the "urlIv" field.
func UrlIvLTE(v []byte) predicate.Url {
	return predicate.Url(sql.FieldLTE(FieldUrlIv, v))
}

// HasPassword applies the HasEdge predicate on the "password" edge.
func HasPassword() predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PasswordTable, PasswordColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPasswordWith applies the HasEdge predicate on the "password" edge with a given conditions (other predicates).
func HasPasswordWith(preds ...predicate.Password) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PasswordInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PasswordTable, PasswordColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
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
func Not(p predicate.Url) predicate.Url {
	return predicate.Url(func(s *sql.Selector) {
		p(s.Not())
	})
}
