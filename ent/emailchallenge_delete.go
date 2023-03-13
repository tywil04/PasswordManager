// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EmailChallengeDelete is the builder for deleting a EmailChallenge entity.
type EmailChallengeDelete struct {
	config
	hooks    []Hook
	mutation *EmailChallengeMutation
}

// Where appends a list predicates to the EmailChallengeDelete builder.
func (ecd *EmailChallengeDelete) Where(ps ...predicate.EmailChallenge) *EmailChallengeDelete {
	ecd.mutation.Where(ps...)
	return ecd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ecd *EmailChallengeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EmailChallengeMutation](ctx, ecd.sqlExec, ecd.mutation, ecd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ecd *EmailChallengeDelete) ExecX(ctx context.Context) int {
	n, err := ecd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ecd *EmailChallengeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(emailchallenge.Table, sqlgraph.NewFieldSpec(emailchallenge.FieldID, field.TypeUUID))
	if ps := ecd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ecd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ecd.mutation.done = true
	return affected, err
}

// EmailChallengeDeleteOne is the builder for deleting a single EmailChallenge entity.
type EmailChallengeDeleteOne struct {
	ecd *EmailChallengeDelete
}

// Where appends a list predicates to the EmailChallengeDelete builder.
func (ecdo *EmailChallengeDeleteOne) Where(ps ...predicate.EmailChallenge) *EmailChallengeDeleteOne {
	ecdo.ecd.mutation.Where(ps...)
	return ecdo
}

// Exec executes the deletion query.
func (ecdo *EmailChallengeDeleteOne) Exec(ctx context.Context) error {
	n, err := ecdo.ecd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{emailchallenge.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ecdo *EmailChallengeDeleteOne) ExecX(ctx context.Context) {
	if err := ecdo.Exec(ctx); err != nil {
		panic(err)
	}
}
