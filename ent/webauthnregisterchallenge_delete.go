// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/webauthnregisterchallenge"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebAuthnRegisterChallengeDelete is the builder for deleting a WebAuthnRegisterChallenge entity.
type WebAuthnRegisterChallengeDelete struct {
	config
	hooks    []Hook
	mutation *WebAuthnRegisterChallengeMutation
}

// Where appends a list predicates to the WebAuthnRegisterChallengeDelete builder.
func (warcd *WebAuthnRegisterChallengeDelete) Where(ps ...predicate.WebAuthnRegisterChallenge) *WebAuthnRegisterChallengeDelete {
	warcd.mutation.Where(ps...)
	return warcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (warcd *WebAuthnRegisterChallengeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, WebAuthnRegisterChallengeMutation](ctx, warcd.sqlExec, warcd.mutation, warcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (warcd *WebAuthnRegisterChallengeDelete) ExecX(ctx context.Context) int {
	n, err := warcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (warcd *WebAuthnRegisterChallengeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: webauthnregisterchallenge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: webauthnregisterchallenge.FieldID,
			},
		},
	}
	if ps := warcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, warcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	warcd.mutation.done = true
	return affected, err
}

// WebAuthnRegisterChallengeDeleteOne is the builder for deleting a single WebAuthnRegisterChallenge entity.
type WebAuthnRegisterChallengeDeleteOne struct {
	warcd *WebAuthnRegisterChallengeDelete
}

// Where appends a list predicates to the WebAuthnRegisterChallengeDelete builder.
func (warcdo *WebAuthnRegisterChallengeDeleteOne) Where(ps ...predicate.WebAuthnRegisterChallenge) *WebAuthnRegisterChallengeDeleteOne {
	warcdo.warcd.mutation.Where(ps...)
	return warcdo
}

// Exec executes the deletion query.
func (warcdo *WebAuthnRegisterChallengeDeleteOne) Exec(ctx context.Context) error {
	n, err := warcdo.warcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{webauthnregisterchallenge.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (warcdo *WebAuthnRegisterChallengeDeleteOne) ExecX(ctx context.Context) {
	if err := warcdo.Exec(ctx); err != nil {
		panic(err)
	}
}