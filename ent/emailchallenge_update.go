// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailChallengeUpdate is the builder for updating EmailChallenge entities.
type EmailChallengeUpdate struct {
	config
	hooks    []Hook
	mutation *EmailChallengeMutation
}

// Where appends a list predicates to the EmailChallengeUpdate builder.
func (ecu *EmailChallengeUpdate) Where(ps ...predicate.EmailChallenge) *EmailChallengeUpdate {
	ecu.mutation.Where(ps...)
	return ecu
}

// SetCode sets the "code" field.
func (ecu *EmailChallengeUpdate) SetCode(s string) *EmailChallengeUpdate {
	ecu.mutation.SetCode(s)
	return ecu
}

// SetExpiry sets the "expiry" field.
func (ecu *EmailChallengeUpdate) SetExpiry(t time.Time) *EmailChallengeUpdate {
	ecu.mutation.SetExpiry(t)
	return ecu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (ecu *EmailChallengeUpdate) SetNillableExpiry(t *time.Time) *EmailChallengeUpdate {
	if t != nil {
		ecu.SetExpiry(*t)
	}
	return ecu
}

// SetFor sets the "for" field.
func (ecu *EmailChallengeUpdate) SetFor(e emailchallenge.For) *EmailChallengeUpdate {
	ecu.mutation.SetFor(e)
	return ecu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ecu *EmailChallengeUpdate) SetUserID(id uuid.UUID) *EmailChallengeUpdate {
	ecu.mutation.SetUserID(id)
	return ecu
}

// SetUser sets the "user" edge to the User entity.
func (ecu *EmailChallengeUpdate) SetUser(u *User) *EmailChallengeUpdate {
	return ecu.SetUserID(u.ID)
}

// Mutation returns the EmailChallengeMutation object of the builder.
func (ecu *EmailChallengeUpdate) Mutation() *EmailChallengeMutation {
	return ecu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ecu *EmailChallengeUpdate) ClearUser() *EmailChallengeUpdate {
	ecu.mutation.ClearUser()
	return ecu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ecu *EmailChallengeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, EmailChallengeMutation](ctx, ecu.sqlSave, ecu.mutation, ecu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecu *EmailChallengeUpdate) SaveX(ctx context.Context) int {
	affected, err := ecu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ecu *EmailChallengeUpdate) Exec(ctx context.Context) error {
	_, err := ecu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecu *EmailChallengeUpdate) ExecX(ctx context.Context) {
	if err := ecu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecu *EmailChallengeUpdate) check() error {
	if v, ok := ecu.mutation.Code(); ok {
		if err := emailchallenge.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EmailChallenge.code": %w`, err)}
		}
	}
	if v, ok := ecu.mutation.For(); ok {
		if err := emailchallenge.ForValidator(v); err != nil {
			return &ValidationError{Name: "for", err: fmt.Errorf(`ent: validator failed for field "EmailChallenge.for": %w`, err)}
		}
	}
	if _, ok := ecu.mutation.UserID(); ecu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EmailChallenge.user"`)
	}
	return nil
}

func (ecu *EmailChallengeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ecu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailchallenge.Table,
			Columns: emailchallenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailchallenge.FieldID,
			},
		},
	}
	if ps := ecu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecu.mutation.Code(); ok {
		_spec.SetField(emailchallenge.FieldCode, field.TypeString, value)
	}
	if value, ok := ecu.mutation.Expiry(); ok {
		_spec.SetField(emailchallenge.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := ecu.mutation.For(); ok {
		_spec.SetField(emailchallenge.FieldFor, field.TypeEnum, value)
	}
	if ecu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailchallenge.UserTable,
			Columns: []string{emailchallenge.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailchallenge.UserTable,
			Columns: []string{emailchallenge.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ecu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailchallenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ecu.mutation.done = true
	return n, nil
}

// EmailChallengeUpdateOne is the builder for updating a single EmailChallenge entity.
type EmailChallengeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EmailChallengeMutation
}

// SetCode sets the "code" field.
func (ecuo *EmailChallengeUpdateOne) SetCode(s string) *EmailChallengeUpdateOne {
	ecuo.mutation.SetCode(s)
	return ecuo
}

// SetExpiry sets the "expiry" field.
func (ecuo *EmailChallengeUpdateOne) SetExpiry(t time.Time) *EmailChallengeUpdateOne {
	ecuo.mutation.SetExpiry(t)
	return ecuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (ecuo *EmailChallengeUpdateOne) SetNillableExpiry(t *time.Time) *EmailChallengeUpdateOne {
	if t != nil {
		ecuo.SetExpiry(*t)
	}
	return ecuo
}

// SetFor sets the "for" field.
func (ecuo *EmailChallengeUpdateOne) SetFor(e emailchallenge.For) *EmailChallengeUpdateOne {
	ecuo.mutation.SetFor(e)
	return ecuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ecuo *EmailChallengeUpdateOne) SetUserID(id uuid.UUID) *EmailChallengeUpdateOne {
	ecuo.mutation.SetUserID(id)
	return ecuo
}

// SetUser sets the "user" edge to the User entity.
func (ecuo *EmailChallengeUpdateOne) SetUser(u *User) *EmailChallengeUpdateOne {
	return ecuo.SetUserID(u.ID)
}

// Mutation returns the EmailChallengeMutation object of the builder.
func (ecuo *EmailChallengeUpdateOne) Mutation() *EmailChallengeMutation {
	return ecuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ecuo *EmailChallengeUpdateOne) ClearUser() *EmailChallengeUpdateOne {
	ecuo.mutation.ClearUser()
	return ecuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ecuo *EmailChallengeUpdateOne) Select(field string, fields ...string) *EmailChallengeUpdateOne {
	ecuo.fields = append([]string{field}, fields...)
	return ecuo
}

// Save executes the query and returns the updated EmailChallenge entity.
func (ecuo *EmailChallengeUpdateOne) Save(ctx context.Context) (*EmailChallenge, error) {
	return withHooks[*EmailChallenge, EmailChallengeMutation](ctx, ecuo.sqlSave, ecuo.mutation, ecuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ecuo *EmailChallengeUpdateOne) SaveX(ctx context.Context) *EmailChallenge {
	node, err := ecuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ecuo *EmailChallengeUpdateOne) Exec(ctx context.Context) error {
	_, err := ecuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecuo *EmailChallengeUpdateOne) ExecX(ctx context.Context) {
	if err := ecuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ecuo *EmailChallengeUpdateOne) check() error {
	if v, ok := ecuo.mutation.Code(); ok {
		if err := emailchallenge.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "EmailChallenge.code": %w`, err)}
		}
	}
	if v, ok := ecuo.mutation.For(); ok {
		if err := emailchallenge.ForValidator(v); err != nil {
			return &ValidationError{Name: "for", err: fmt.Errorf(`ent: validator failed for field "EmailChallenge.for": %w`, err)}
		}
	}
	if _, ok := ecuo.mutation.UserID(); ecuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "EmailChallenge.user"`)
	}
	return nil
}

func (ecuo *EmailChallengeUpdateOne) sqlSave(ctx context.Context) (_node *EmailChallenge, err error) {
	if err := ecuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   emailchallenge.Table,
			Columns: emailchallenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: emailchallenge.FieldID,
			},
		},
	}
	id, ok := ecuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EmailChallenge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ecuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailchallenge.FieldID)
		for _, f := range fields {
			if !emailchallenge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != emailchallenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ecuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ecuo.mutation.Code(); ok {
		_spec.SetField(emailchallenge.FieldCode, field.TypeString, value)
	}
	if value, ok := ecuo.mutation.Expiry(); ok {
		_spec.SetField(emailchallenge.FieldExpiry, field.TypeTime, value)
	}
	if value, ok := ecuo.mutation.For(); ok {
		_spec.SetField(emailchallenge.FieldFor, field.TypeEnum, value)
	}
	if ecuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailchallenge.UserTable,
			Columns: []string{emailchallenge.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   emailchallenge.UserTable,
			Columns: []string{emailchallenge.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EmailChallenge{config: ecuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ecuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{emailchallenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ecuo.mutation.done = true
	return _node, nil
}