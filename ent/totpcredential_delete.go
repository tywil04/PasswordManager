// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/totpcredential"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TotpCredentialDelete is the builder for deleting a TotpCredential entity.
type TotpCredentialDelete struct {
	config
	hooks    []Hook
	mutation *TotpCredentialMutation
}

// Where appends a list predicates to the TotpCredentialDelete builder.
func (tcd *TotpCredentialDelete) Where(ps ...predicate.TotpCredential) *TotpCredentialDelete {
	tcd.mutation.Where(ps...)
	return tcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tcd *TotpCredentialDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, TotpCredentialMutation](ctx, tcd.sqlExec, tcd.mutation, tcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tcd *TotpCredentialDelete) ExecX(ctx context.Context) int {
	n, err := tcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tcd *TotpCredentialDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(totpcredential.Table, sqlgraph.NewFieldSpec(totpcredential.FieldID, field.TypeUUID))
	if ps := tcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tcd.mutation.done = true
	return affected, err
}

// TotpCredentialDeleteOne is the builder for deleting a single TotpCredential entity.
type TotpCredentialDeleteOne struct {
	tcd *TotpCredentialDelete
}

// Where appends a list predicates to the TotpCredentialDelete builder.
func (tcdo *TotpCredentialDeleteOne) Where(ps ...predicate.TotpCredential) *TotpCredentialDeleteOne {
	tcdo.tcd.mutation.Where(ps...)
	return tcdo
}

// Exec executes the deletion query.
func (tcdo *TotpCredentialDeleteOne) Exec(ctx context.Context) error {
	n, err := tcdo.tcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{totpcredential.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tcdo *TotpCredentialDeleteOne) ExecX(ctx context.Context) {
	if err := tcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
