// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/predicate"
	"context"
	"errors"
	"fmt"

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

// SetNillableCode sets the "code" field if the given value is not nil.
func (ecu *EmailChallengeUpdate) SetNillableCode(s *string) *EmailChallengeUpdate {
	if s != nil {
		ecu.SetCode(*s)
	}
	return ecu
}

// ClearCode clears the value of the "code" field.
func (ecu *EmailChallengeUpdate) ClearCode() *EmailChallengeUpdate {
	ecu.mutation.ClearCode()
	return ecu
}

// SetChallengeID sets the "challenge" edge to the Challenge entity by ID.
func (ecu *EmailChallengeUpdate) SetChallengeID(id uuid.UUID) *EmailChallengeUpdate {
	ecu.mutation.SetChallengeID(id)
	return ecu
}

// SetNillableChallengeID sets the "challenge" edge to the Challenge entity by ID if the given value is not nil.
func (ecu *EmailChallengeUpdate) SetNillableChallengeID(id *uuid.UUID) *EmailChallengeUpdate {
	if id != nil {
		ecu = ecu.SetChallengeID(*id)
	}
	return ecu
}

// SetChallenge sets the "challenge" edge to the Challenge entity.
func (ecu *EmailChallengeUpdate) SetChallenge(c *Challenge) *EmailChallengeUpdate {
	return ecu.SetChallengeID(c.ID)
}

// Mutation returns the EmailChallengeMutation object of the builder.
func (ecu *EmailChallengeUpdate) Mutation() *EmailChallengeMutation {
	return ecu.mutation
}

// ClearChallenge clears the "challenge" edge to the Challenge entity.
func (ecu *EmailChallengeUpdate) ClearChallenge() *EmailChallengeUpdate {
	ecu.mutation.ClearChallenge()
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

func (ecu *EmailChallengeUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
	if ecu.mutation.CodeCleared() {
		_spec.ClearField(emailchallenge.FieldCode, field.TypeString)
	}
	if ecu.mutation.ChallengeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailchallenge.ChallengeTable,
			Columns: []string{emailchallenge.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecu.mutation.ChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailchallenge.ChallengeTable,
			Columns: []string{emailchallenge.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
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

// SetNillableCode sets the "code" field if the given value is not nil.
func (ecuo *EmailChallengeUpdateOne) SetNillableCode(s *string) *EmailChallengeUpdateOne {
	if s != nil {
		ecuo.SetCode(*s)
	}
	return ecuo
}

// ClearCode clears the value of the "code" field.
func (ecuo *EmailChallengeUpdateOne) ClearCode() *EmailChallengeUpdateOne {
	ecuo.mutation.ClearCode()
	return ecuo
}

// SetChallengeID sets the "challenge" edge to the Challenge entity by ID.
func (ecuo *EmailChallengeUpdateOne) SetChallengeID(id uuid.UUID) *EmailChallengeUpdateOne {
	ecuo.mutation.SetChallengeID(id)
	return ecuo
}

// SetNillableChallengeID sets the "challenge" edge to the Challenge entity by ID if the given value is not nil.
func (ecuo *EmailChallengeUpdateOne) SetNillableChallengeID(id *uuid.UUID) *EmailChallengeUpdateOne {
	if id != nil {
		ecuo = ecuo.SetChallengeID(*id)
	}
	return ecuo
}

// SetChallenge sets the "challenge" edge to the Challenge entity.
func (ecuo *EmailChallengeUpdateOne) SetChallenge(c *Challenge) *EmailChallengeUpdateOne {
	return ecuo.SetChallengeID(c.ID)
}

// Mutation returns the EmailChallengeMutation object of the builder.
func (ecuo *EmailChallengeUpdateOne) Mutation() *EmailChallengeMutation {
	return ecuo.mutation
}

// ClearChallenge clears the "challenge" edge to the Challenge entity.
func (ecuo *EmailChallengeUpdateOne) ClearChallenge() *EmailChallengeUpdateOne {
	ecuo.mutation.ClearChallenge()
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

func (ecuo *EmailChallengeUpdateOne) sqlSave(ctx context.Context) (_node *EmailChallenge, err error) {
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
	if ecuo.mutation.CodeCleared() {
		_spec.ClearField(emailchallenge.FieldCode, field.TypeString)
	}
	if ecuo.mutation.ChallengeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailchallenge.ChallengeTable,
			Columns: []string{emailchallenge.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ecuo.mutation.ChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   emailchallenge.ChallengeTable,
			Columns: []string{emailchallenge.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
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
