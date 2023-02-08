// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/additionalfield"
	"PasswordManager/ent/password"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdditionalFieldCreate is the builder for creating a AdditionalField entity.
type AdditionalFieldCreate struct {
	config
	mutation *AdditionalFieldMutation
	hooks    []Hook
}

// SetKey sets the "key" field.
func (afc *AdditionalFieldCreate) SetKey(b []byte) *AdditionalFieldCreate {
	afc.mutation.SetKey(b)
	return afc
}

// SetKeyIv sets the "keyIv" field.
func (afc *AdditionalFieldCreate) SetKeyIv(b []byte) *AdditionalFieldCreate {
	afc.mutation.SetKeyIv(b)
	return afc
}

// SetValue sets the "value" field.
func (afc *AdditionalFieldCreate) SetValue(b []byte) *AdditionalFieldCreate {
	afc.mutation.SetValue(b)
	return afc
}

// SetValueIv sets the "valueIv" field.
func (afc *AdditionalFieldCreate) SetValueIv(b []byte) *AdditionalFieldCreate {
	afc.mutation.SetValueIv(b)
	return afc
}

// SetID sets the "id" field.
func (afc *AdditionalFieldCreate) SetID(u uuid.UUID) *AdditionalFieldCreate {
	afc.mutation.SetID(u)
	return afc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (afc *AdditionalFieldCreate) SetNillableID(u *uuid.UUID) *AdditionalFieldCreate {
	if u != nil {
		afc.SetID(*u)
	}
	return afc
}

// SetPasswordID sets the "password" edge to the Password entity by ID.
func (afc *AdditionalFieldCreate) SetPasswordID(id uuid.UUID) *AdditionalFieldCreate {
	afc.mutation.SetPasswordID(id)
	return afc
}

// SetNillablePasswordID sets the "password" edge to the Password entity by ID if the given value is not nil.
func (afc *AdditionalFieldCreate) SetNillablePasswordID(id *uuid.UUID) *AdditionalFieldCreate {
	if id != nil {
		afc = afc.SetPasswordID(*id)
	}
	return afc
}

// SetPassword sets the "password" edge to the Password entity.
func (afc *AdditionalFieldCreate) SetPassword(p *Password) *AdditionalFieldCreate {
	return afc.SetPasswordID(p.ID)
}

// Mutation returns the AdditionalFieldMutation object of the builder.
func (afc *AdditionalFieldCreate) Mutation() *AdditionalFieldMutation {
	return afc.mutation
}

// Save creates the AdditionalField in the database.
func (afc *AdditionalFieldCreate) Save(ctx context.Context) (*AdditionalField, error) {
	afc.defaults()
	return withHooks[*AdditionalField, AdditionalFieldMutation](ctx, afc.sqlSave, afc.mutation, afc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (afc *AdditionalFieldCreate) SaveX(ctx context.Context) *AdditionalField {
	v, err := afc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afc *AdditionalFieldCreate) Exec(ctx context.Context) error {
	_, err := afc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afc *AdditionalFieldCreate) ExecX(ctx context.Context) {
	if err := afc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afc *AdditionalFieldCreate) defaults() {
	if _, ok := afc.mutation.ID(); !ok {
		v := additionalfield.DefaultID()
		afc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afc *AdditionalFieldCreate) check() error {
	if _, ok := afc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "AdditionalField.key"`)}
	}
	if v, ok := afc.mutation.Key(); ok {
		if err := additionalfield.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.key": %w`, err)}
		}
	}
	if _, ok := afc.mutation.KeyIv(); !ok {
		return &ValidationError{Name: "keyIv", err: errors.New(`ent: missing required field "AdditionalField.keyIv"`)}
	}
	if v, ok := afc.mutation.KeyIv(); ok {
		if err := additionalfield.KeyIvValidator(v); err != nil {
			return &ValidationError{Name: "keyIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.keyIv": %w`, err)}
		}
	}
	if _, ok := afc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "AdditionalField.value"`)}
	}
	if v, ok := afc.mutation.Value(); ok {
		if err := additionalfield.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.value": %w`, err)}
		}
	}
	if _, ok := afc.mutation.ValueIv(); !ok {
		return &ValidationError{Name: "valueIv", err: errors.New(`ent: missing required field "AdditionalField.valueIv"`)}
	}
	if v, ok := afc.mutation.ValueIv(); ok {
		if err := additionalfield.ValueIvValidator(v); err != nil {
			return &ValidationError{Name: "valueIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.valueIv": %w`, err)}
		}
	}
	return nil
}

func (afc *AdditionalFieldCreate) sqlSave(ctx context.Context) (*AdditionalField, error) {
	if err := afc.check(); err != nil {
		return nil, err
	}
	_node, _spec := afc.createSpec()
	if err := sqlgraph.CreateNode(ctx, afc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	afc.mutation.id = &_node.ID
	afc.mutation.done = true
	return _node, nil
}

func (afc *AdditionalFieldCreate) createSpec() (*AdditionalField, *sqlgraph.CreateSpec) {
	var (
		_node = &AdditionalField{config: afc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: additionalfield.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: additionalfield.FieldID,
			},
		}
	)
	if id, ok := afc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := afc.mutation.Key(); ok {
		_spec.SetField(additionalfield.FieldKey, field.TypeBytes, value)
		_node.Key = value
	}
	if value, ok := afc.mutation.KeyIv(); ok {
		_spec.SetField(additionalfield.FieldKeyIv, field.TypeBytes, value)
		_node.KeyIv = value
	}
	if value, ok := afc.mutation.Value(); ok {
		_spec.SetField(additionalfield.FieldValue, field.TypeBytes, value)
		_node.Value = value
	}
	if value, ok := afc.mutation.ValueIv(); ok {
		_spec.SetField(additionalfield.FieldValueIv, field.TypeBytes, value)
		_node.ValueIv = value
	}
	if nodes := afc.mutation.PasswordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   additionalfield.PasswordTable,
			Columns: []string{additionalfield.PasswordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: password.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.password_additional_fields = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AdditionalFieldCreateBulk is the builder for creating many AdditionalField entities in bulk.
type AdditionalFieldCreateBulk struct {
	config
	builders []*AdditionalFieldCreate
}

// Save creates the AdditionalField entities in the database.
func (afcb *AdditionalFieldCreateBulk) Save(ctx context.Context) ([]*AdditionalField, error) {
	specs := make([]*sqlgraph.CreateSpec, len(afcb.builders))
	nodes := make([]*AdditionalField, len(afcb.builders))
	mutators := make([]Mutator, len(afcb.builders))
	for i := range afcb.builders {
		func(i int, root context.Context) {
			builder := afcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdditionalFieldMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, afcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, afcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, afcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (afcb *AdditionalFieldCreateBulk) SaveX(ctx context.Context) []*AdditionalField {
	v, err := afcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afcb *AdditionalFieldCreateBulk) Exec(ctx context.Context) error {
	_, err := afcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afcb *AdditionalFieldCreateBulk) ExecX(ctx context.Context) {
	if err := afcb.Exec(ctx); err != nil {
		panic(err)
	}
}