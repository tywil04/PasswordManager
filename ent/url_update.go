// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/password"
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/url"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// URLUpdate is the builder for updating Url entities.
type URLUpdate struct {
	config
	hooks    []Hook
	mutation *URLMutation
}

// Where appends a list predicates to the URLUpdate builder.
func (uu *URLUpdate) Where(ps ...predicate.Url) *URLUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetURL sets the "url" field.
func (uu *URLUpdate) SetURL(b []byte) *URLUpdate {
	uu.mutation.SetURL(b)
	return uu
}

// SetUrlIv sets the "urlIv" field.
func (uu *URLUpdate) SetUrlIv(b []byte) *URLUpdate {
	uu.mutation.SetUrlIv(b)
	return uu
}

// SetPasswordID sets the "password" edge to the Password entity by ID.
func (uu *URLUpdate) SetPasswordID(id uuid.UUID) *URLUpdate {
	uu.mutation.SetPasswordID(id)
	return uu
}

// SetNillablePasswordID sets the "password" edge to the Password entity by ID if the given value is not nil.
func (uu *URLUpdate) SetNillablePasswordID(id *uuid.UUID) *URLUpdate {
	if id != nil {
		uu = uu.SetPasswordID(*id)
	}
	return uu
}

// SetPassword sets the "password" edge to the Password entity.
func (uu *URLUpdate) SetPassword(p *Password) *URLUpdate {
	return uu.SetPasswordID(p.ID)
}

// Mutation returns the URLMutation object of the builder.
func (uu *URLUpdate) Mutation() *URLMutation {
	return uu.mutation
}

// ClearPassword clears the "password" edge to the Password entity.
func (uu *URLUpdate) ClearPassword() *URLUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *URLUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, URLMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *URLUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *URLUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *URLUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *URLUpdate) check() error {
	if v, ok := uu.mutation.URL(); ok {
		if err := url.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Url.url": %w`, err)}
		}
	}
	if v, ok := uu.mutation.UrlIv(); ok {
		if err := url.UrlIvValidator(v); err != nil {
			return &ValidationError{Name: "urlIv", err: fmt.Errorf(`ent: validator failed for field "Url.urlIv": %w`, err)}
		}
	}
	return nil
}

func (uu *URLUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   url.Table,
			Columns: url.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: url.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeBytes, value)
	}
	if value, ok := uu.mutation.UrlIv(); ok {
		_spec.SetField(url.FieldUrlIv, field.TypeBytes, value)
	}
	if uu.mutation.PasswordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.PasswordTable,
			Columns: []string{url.PasswordColumn},
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
	if nodes := uu.mutation.PasswordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.PasswordTable,
			Columns: []string{url.PasswordColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// URLUpdateOne is the builder for updating a single Url entity.
type URLUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *URLMutation
}

// SetURL sets the "url" field.
func (uuo *URLUpdateOne) SetURL(b []byte) *URLUpdateOne {
	uuo.mutation.SetURL(b)
	return uuo
}

// SetUrlIv sets the "urlIv" field.
func (uuo *URLUpdateOne) SetUrlIv(b []byte) *URLUpdateOne {
	uuo.mutation.SetUrlIv(b)
	return uuo
}

// SetPasswordID sets the "password" edge to the Password entity by ID.
func (uuo *URLUpdateOne) SetPasswordID(id uuid.UUID) *URLUpdateOne {
	uuo.mutation.SetPasswordID(id)
	return uuo
}

// SetNillablePasswordID sets the "password" edge to the Password entity by ID if the given value is not nil.
func (uuo *URLUpdateOne) SetNillablePasswordID(id *uuid.UUID) *URLUpdateOne {
	if id != nil {
		uuo = uuo.SetPasswordID(*id)
	}
	return uuo
}

// SetPassword sets the "password" edge to the Password entity.
func (uuo *URLUpdateOne) SetPassword(p *Password) *URLUpdateOne {
	return uuo.SetPasswordID(p.ID)
}

// Mutation returns the URLMutation object of the builder.
func (uuo *URLUpdateOne) Mutation() *URLMutation {
	return uuo.mutation
}

// ClearPassword clears the "password" edge to the Password entity.
func (uuo *URLUpdateOne) ClearPassword() *URLUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *URLUpdateOne) Select(field string, fields ...string) *URLUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Url entity.
func (uuo *URLUpdateOne) Save(ctx context.Context) (*Url, error) {
	return withHooks[*Url, URLMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *URLUpdateOne) SaveX(ctx context.Context) *Url {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *URLUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *URLUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *URLUpdateOne) check() error {
	if v, ok := uuo.mutation.URL(); ok {
		if err := url.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Url.url": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.UrlIv(); ok {
		if err := url.UrlIvValidator(v); err != nil {
			return &ValidationError{Name: "urlIv", err: fmt.Errorf(`ent: validator failed for field "Url.urlIv": %w`, err)}
		}
	}
	return nil
}

func (uuo *URLUpdateOne) sqlSave(ctx context.Context) (_node *Url, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   url.Table,
			Columns: url.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: url.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Url.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, url.FieldID)
		for _, f := range fields {
			if !url.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != url.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(url.FieldURL, field.TypeBytes, value)
	}
	if value, ok := uuo.mutation.UrlIv(); ok {
		_spec.SetField(url.FieldUrlIv, field.TypeBytes, value)
	}
	if uuo.mutation.PasswordCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.PasswordTable,
			Columns: []string{url.PasswordColumn},
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
	if nodes := uuo.mutation.PasswordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   url.PasswordTable,
			Columns: []string{url.PasswordColumn},
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
	_node = &Url{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{url.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}