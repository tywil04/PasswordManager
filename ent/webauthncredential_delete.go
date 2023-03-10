// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/webauthncredential"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WebAuthnCredentialDelete is the builder for deleting a WebAuthnCredential entity.
type WebAuthnCredentialDelete struct {
	config
	hooks    []Hook
	mutation *WebAuthnCredentialMutation
}

// Where appends a list predicates to the WebAuthnCredentialDelete builder.
func (wacd *WebAuthnCredentialDelete) Where(ps ...predicate.WebAuthnCredential) *WebAuthnCredentialDelete {
	wacd.mutation.Where(ps...)
	return wacd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wacd *WebAuthnCredentialDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, WebAuthnCredentialMutation](ctx, wacd.sqlExec, wacd.mutation, wacd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wacd *WebAuthnCredentialDelete) ExecX(ctx context.Context) int {
	n, err := wacd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wacd *WebAuthnCredentialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(webauthncredential.Table, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeUUID))
	if ps := wacd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wacd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wacd.mutation.done = true
	return affected, err
}

// WebAuthnCredentialDeleteOne is the builder for deleting a single WebAuthnCredential entity.
type WebAuthnCredentialDeleteOne struct {
	wacd *WebAuthnCredentialDelete
}

// Where appends a list predicates to the WebAuthnCredentialDelete builder.
func (wacdo *WebAuthnCredentialDeleteOne) Where(ps ...predicate.WebAuthnCredential) *WebAuthnCredentialDeleteOne {
	wacdo.wacd.mutation.Where(ps...)
	return wacdo
}

// Exec executes the deletion query.
func (wacdo *WebAuthnCredentialDeleteOne) Exec(ctx context.Context) error {
	n, err := wacdo.wacd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{webauthncredential.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wacdo *WebAuthnCredentialDeleteOne) ExecX(ctx context.Context) {
	if err := wacdo.Exec(ctx); err != nil {
		panic(err)
	}
}
