// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/additionalfield"
	"PasswordManager/ent/password"
	"PasswordManager/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdditionalFieldUpdate is the builder for updating AdditionalField entities.
type AdditionalFieldUpdate struct {
	config
	hooks    []Hook
	mutation *AdditionalFieldMutation
}

// Where appends a list predicates to the AdditionalFieldUpdate builder.
func (afu *AdditionalFieldUpdate) Where(ps ...predicate.AdditionalField) *AdditionalFieldUpdate {
	afu.mutation.Where(ps...)
	return afu
}

// SetKey sets the "key" field.
func (afu *AdditionalFieldUpdate) SetKey(b []byte) *AdditionalFieldUpdate {
	afu.mutation.SetKey(b)
	return afu
}

// SetKeyIv sets the "keyIv" field.
func (afu *AdditionalFieldUpdate) SetKeyIv(b []byte) *AdditionalFieldUpdate {
	afu.mutation.SetKeyIv(b)
	return afu
}

// SetValue sets the "value" field.
func (afu *AdditionalFieldUpdate) SetValue(b []byte) *AdditionalFieldUpdate {
	afu.mutation.SetValue(b)
	return afu
}

// SetValueIv sets the "valueIv" field.
func (afu *AdditionalFieldUpdate) SetValueIv(b []byte) *AdditionalFieldUpdate {
	afu.mutation.SetValueIv(b)
	return afu
}

// SetPasswordID sets the "password" edge to the Password entity by ID.
func (afu *AdditionalFieldUpdate) SetPasswordID(id uuid.UUID) *AdditionalFieldUpdate {
	afu.mutation.SetPasswordID(id)
	return afu
}

// SetNillablePasswordID sets the "password" edge to the Password entity by ID if the given value is not nil.
func (afu *AdditionalFieldUpdate) SetNillablePasswordID(id *uuid.UUID) *AdditionalFieldUpdate {
	if id != nil {
		afu = afu.SetPasswordID(*id)
	}
	return afu
}

// SetPassword sets the "password" edge to the Password entity.
func (afu *AdditionalFieldUpdate) SetPassword(p *Password) *AdditionalFieldUpdate {
	return afu.SetPasswordID(p.ID)
}

// Mutation returns the AdditionalFieldMutation object of the builder.
func (afu *AdditionalFieldUpdate) Mutation() *AdditionalFieldMutation {
	return afu.mutation
}

// ClearPassword clears the "password" edge to the Password entity.
func (afu *AdditionalFieldUpdate) ClearPassword() *AdditionalFieldUpdate {
	afu.mutation.ClearPassword()
	return afu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (afu *AdditionalFieldUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, AdditionalFieldMutation](ctx, afu.sqlSave, afu.mutation, afu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (afu *AdditionalFieldUpdate) SaveX(ctx context.Context) int {
	affected, err := afu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (afu *AdditionalFieldUpdate) Exec(ctx context.Context) error {
	_, err := afu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afu *AdditionalFieldUpdate) ExecX(ctx context.Context) {
	if err := afu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afu *AdditionalFieldUpdate) check() error {
	if v, ok := afu.mutation.Key(); ok {
		if err := additionalfield.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.key": %w`, err)}
		}
	}
	if v, ok := afu.mutation.KeyIv(); ok {
		if err := additionalfield.KeyIvValidator(v); err != nil {
			return &ValidationError{Name: "keyIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.keyIv": %w`, err)}
		}
	}
	if v, ok := afu.mutation.Value(); ok {
		if err := additionalfield.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.value": %w`, err)}
		}
	}
	if v, ok := afu.mutation.ValueIv(); ok {
		if err := additionalfield.ValueIvValidator(v); err != nil {
			return &ValidationError{Name: "valueIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.valueIv": %w`, err)}
		}
	}
	return nil
}

func (afu *AdditionalFieldUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := afu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   additionalfield.Table,
			Columns: additionalfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: additionalfield.FieldID,
			},
		},
	}
	if ps := afu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afu.mutation.Key(); ok {
		_spec.SetField(additionalfield.FieldKey, field.TypeBytes, value)
	}
	if value, ok := afu.mutation.KeyIv(); ok {
		_spec.SetField(additionalfield.FieldKeyIv, field.TypeBytes, value)
	}
	if value, ok := afu.mutation.Value(); ok {
		_spec.SetField(additionalfield.FieldValue, field.TypeBytes, value)
	}
	if value, ok := afu.mutation.ValueIv(); ok {
		_spec.SetField(additionalfield.FieldValueIv, field.TypeBytes, value)
	}
	if afu.mutation.PasswordCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := afu.mutation.PasswordIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, afu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{additionalfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	afu.mutation.done = true
	return n, nil
}

// AdditionalFieldUpdateOne is the builder for updating a single AdditionalField entity.
type AdditionalFieldUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdditionalFieldMutation
}

// SetKey sets the "key" field.
func (afuo *AdditionalFieldUpdateOne) SetKey(b []byte) *AdditionalFieldUpdateOne {
	afuo.mutation.SetKey(b)
	return afuo
}

// SetKeyIv sets the "keyIv" field.
func (afuo *AdditionalFieldUpdateOne) SetKeyIv(b []byte) *AdditionalFieldUpdateOne {
	afuo.mutation.SetKeyIv(b)
	return afuo
}

// SetValue sets the "value" field.
func (afuo *AdditionalFieldUpdateOne) SetValue(b []byte) *AdditionalFieldUpdateOne {
	afuo.mutation.SetValue(b)
	return afuo
}

// SetValueIv sets the "valueIv" field.
func (afuo *AdditionalFieldUpdateOne) SetValueIv(b []byte) *AdditionalFieldUpdateOne {
	afuo.mutation.SetValueIv(b)
	return afuo
}

// SetPasswordID sets the "password" edge to the Password entity by ID.
func (afuo *AdditionalFieldUpdateOne) SetPasswordID(id uuid.UUID) *AdditionalFieldUpdateOne {
	afuo.mutation.SetPasswordID(id)
	return afuo
}

// SetNillablePasswordID sets the "password" edge to the Password entity by ID if the given value is not nil.
func (afuo *AdditionalFieldUpdateOne) SetNillablePasswordID(id *uuid.UUID) *AdditionalFieldUpdateOne {
	if id != nil {
		afuo = afuo.SetPasswordID(*id)
	}
	return afuo
}

// SetPassword sets the "password" edge to the Password entity.
func (afuo *AdditionalFieldUpdateOne) SetPassword(p *Password) *AdditionalFieldUpdateOne {
	return afuo.SetPasswordID(p.ID)
}

// Mutation returns the AdditionalFieldMutation object of the builder.
func (afuo *AdditionalFieldUpdateOne) Mutation() *AdditionalFieldMutation {
	return afuo.mutation
}

// ClearPassword clears the "password" edge to the Password entity.
func (afuo *AdditionalFieldUpdateOne) ClearPassword() *AdditionalFieldUpdateOne {
	afuo.mutation.ClearPassword()
	return afuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (afuo *AdditionalFieldUpdateOne) Select(field string, fields ...string) *AdditionalFieldUpdateOne {
	afuo.fields = append([]string{field}, fields...)
	return afuo
}

// Save executes the query and returns the updated AdditionalField entity.
func (afuo *AdditionalFieldUpdateOne) Save(ctx context.Context) (*AdditionalField, error) {
	return withHooks[*AdditionalField, AdditionalFieldMutation](ctx, afuo.sqlSave, afuo.mutation, afuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (afuo *AdditionalFieldUpdateOne) SaveX(ctx context.Context) *AdditionalField {
	node, err := afuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (afuo *AdditionalFieldUpdateOne) Exec(ctx context.Context) error {
	_, err := afuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afuo *AdditionalFieldUpdateOne) ExecX(ctx context.Context) {
	if err := afuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (afuo *AdditionalFieldUpdateOne) check() error {
	if v, ok := afuo.mutation.Key(); ok {
		if err := additionalfield.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.key": %w`, err)}
		}
	}
	if v, ok := afuo.mutation.KeyIv(); ok {
		if err := additionalfield.KeyIvValidator(v); err != nil {
			return &ValidationError{Name: "keyIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.keyIv": %w`, err)}
		}
	}
	if v, ok := afuo.mutation.Value(); ok {
		if err := additionalfield.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.value": %w`, err)}
		}
	}
	if v, ok := afuo.mutation.ValueIv(); ok {
		if err := additionalfield.ValueIvValidator(v); err != nil {
			return &ValidationError{Name: "valueIv", err: fmt.Errorf(`ent: validator failed for field "AdditionalField.valueIv": %w`, err)}
		}
	}
	return nil
}

func (afuo *AdditionalFieldUpdateOne) sqlSave(ctx context.Context) (_node *AdditionalField, err error) {
	if err := afuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   additionalfield.Table,
			Columns: additionalfield.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: additionalfield.FieldID,
			},
		},
	}
	id, ok := afuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdditionalField.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := afuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, additionalfield.FieldID)
		for _, f := range fields {
			if !additionalfield.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != additionalfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := afuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afuo.mutation.Key(); ok {
		_spec.SetField(additionalfield.FieldKey, field.TypeBytes, value)
	}
	if value, ok := afuo.mutation.KeyIv(); ok {
		_spec.SetField(additionalfield.FieldKeyIv, field.TypeBytes, value)
	}
	if value, ok := afuo.mutation.Value(); ok {
		_spec.SetField(additionalfield.FieldValue, field.TypeBytes, value)
	}
	if value, ok := afuo.mutation.ValueIv(); ok {
		_spec.SetField(additionalfield.FieldValueIv, field.TypeBytes, value)
	}
	if afuo.mutation.PasswordCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := afuo.mutation.PasswordIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdditionalField{config: afuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, afuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{additionalfield.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	afuo.mutation.done = true
	return _node, nil
}
