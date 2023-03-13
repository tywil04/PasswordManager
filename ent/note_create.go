// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/note"
	"PasswordManager/ent/vault"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// NoteCreate is the builder for creating a Note entity.
type NoteCreate struct {
	config
	mutation *NoteMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (nc *NoteCreate) SetName(b []byte) *NoteCreate {
	nc.mutation.SetName(b)
	return nc
}

// SetNameIv sets the "nameIv" field.
func (nc *NoteCreate) SetNameIv(b []byte) *NoteCreate {
	nc.mutation.SetNameIv(b)
	return nc
}

// SetTitle sets the "title" field.
func (nc *NoteCreate) SetTitle(b []byte) *NoteCreate {
	nc.mutation.SetTitle(b)
	return nc
}

// SetTitleIv sets the "titleIv" field.
func (nc *NoteCreate) SetTitleIv(b []byte) *NoteCreate {
	nc.mutation.SetTitleIv(b)
	return nc
}

// SetContent sets the "content" field.
func (nc *NoteCreate) SetContent(b []byte) *NoteCreate {
	nc.mutation.SetContent(b)
	return nc
}

// SetContentIv sets the "contentIv" field.
func (nc *NoteCreate) SetContentIv(b []byte) *NoteCreate {
	nc.mutation.SetContentIv(b)
	return nc
}

// SetColour sets the "colour" field.
func (nc *NoteCreate) SetColour(b []byte) *NoteCreate {
	nc.mutation.SetColour(b)
	return nc
}

// SetColourIv sets the "colourIv" field.
func (nc *NoteCreate) SetColourIv(b []byte) *NoteCreate {
	nc.mutation.SetColourIv(b)
	return nc
}

// SetID sets the "id" field.
func (nc *NoteCreate) SetID(u uuid.UUID) *NoteCreate {
	nc.mutation.SetID(u)
	return nc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (nc *NoteCreate) SetNillableID(u *uuid.UUID) *NoteCreate {
	if u != nil {
		nc.SetID(*u)
	}
	return nc
}

// SetVaultID sets the "vault" edge to the Vault entity by ID.
func (nc *NoteCreate) SetVaultID(id uuid.UUID) *NoteCreate {
	nc.mutation.SetVaultID(id)
	return nc
}

// SetNillableVaultID sets the "vault" edge to the Vault entity by ID if the given value is not nil.
func (nc *NoteCreate) SetNillableVaultID(id *uuid.UUID) *NoteCreate {
	if id != nil {
		nc = nc.SetVaultID(*id)
	}
	return nc
}

// SetVault sets the "vault" edge to the Vault entity.
func (nc *NoteCreate) SetVault(v *Vault) *NoteCreate {
	return nc.SetVaultID(v.ID)
}

// Mutation returns the NoteMutation object of the builder.
func (nc *NoteCreate) Mutation() *NoteMutation {
	return nc.mutation
}

// Save creates the Note in the database.
func (nc *NoteCreate) Save(ctx context.Context) (*Note, error) {
	nc.defaults()
	return withHooks[*Note, NoteMutation](ctx, nc.sqlSave, nc.mutation, nc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NoteCreate) SaveX(ctx context.Context) *Note {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nc *NoteCreate) Exec(ctx context.Context) error {
	_, err := nc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nc *NoteCreate) ExecX(ctx context.Context) {
	if err := nc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nc *NoteCreate) defaults() {
	if _, ok := nc.mutation.ID(); !ok {
		v := note.DefaultID()
		nc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nc *NoteCreate) check() error {
	if _, ok := nc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Note.name"`)}
	}
	if v, ok := nc.mutation.Name(); ok {
		if err := note.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Note.name": %w`, err)}
		}
	}
	if _, ok := nc.mutation.NameIv(); !ok {
		return &ValidationError{Name: "nameIv", err: errors.New(`ent: missing required field "Note.nameIv"`)}
	}
	if v, ok := nc.mutation.NameIv(); ok {
		if err := note.NameIvValidator(v); err != nil {
			return &ValidationError{Name: "nameIv", err: fmt.Errorf(`ent: validator failed for field "Note.nameIv": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Note.title"`)}
	}
	if v, ok := nc.mutation.Title(); ok {
		if err := note.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Note.title": %w`, err)}
		}
	}
	if _, ok := nc.mutation.TitleIv(); !ok {
		return &ValidationError{Name: "titleIv", err: errors.New(`ent: missing required field "Note.titleIv"`)}
	}
	if v, ok := nc.mutation.TitleIv(); ok {
		if err := note.TitleIvValidator(v); err != nil {
			return &ValidationError{Name: "titleIv", err: fmt.Errorf(`ent: validator failed for field "Note.titleIv": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Note.content"`)}
	}
	if v, ok := nc.mutation.Content(); ok {
		if err := note.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Note.content": %w`, err)}
		}
	}
	if _, ok := nc.mutation.ContentIv(); !ok {
		return &ValidationError{Name: "contentIv", err: errors.New(`ent: missing required field "Note.contentIv"`)}
	}
	if v, ok := nc.mutation.ContentIv(); ok {
		if err := note.ContentIvValidator(v); err != nil {
			return &ValidationError{Name: "contentIv", err: fmt.Errorf(`ent: validator failed for field "Note.contentIv": %w`, err)}
		}
	}
	if _, ok := nc.mutation.Colour(); !ok {
		return &ValidationError{Name: "colour", err: errors.New(`ent: missing required field "Note.colour"`)}
	}
	if v, ok := nc.mutation.Colour(); ok {
		if err := note.ColourValidator(v); err != nil {
			return &ValidationError{Name: "colour", err: fmt.Errorf(`ent: validator failed for field "Note.colour": %w`, err)}
		}
	}
	if _, ok := nc.mutation.ColourIv(); !ok {
		return &ValidationError{Name: "colourIv", err: errors.New(`ent: missing required field "Note.colourIv"`)}
	}
	if v, ok := nc.mutation.ColourIv(); ok {
		if err := note.ColourIvValidator(v); err != nil {
			return &ValidationError{Name: "colourIv", err: fmt.Errorf(`ent: validator failed for field "Note.colourIv": %w`, err)}
		}
	}
	return nil
}

func (nc *NoteCreate) sqlSave(ctx context.Context) (*Note, error) {
	if err := nc.check(); err != nil {
		return nil, err
	}
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
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
	nc.mutation.id = &_node.ID
	nc.mutation.done = true
	return _node, nil
}

func (nc *NoteCreate) createSpec() (*Note, *sqlgraph.CreateSpec) {
	var (
		_node = &Note{config: nc.config}
		_spec = sqlgraph.NewCreateSpec(note.Table, sqlgraph.NewFieldSpec(note.FieldID, field.TypeUUID))
	)
	if id, ok := nc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := nc.mutation.Name(); ok {
		_spec.SetField(note.FieldName, field.TypeBytes, value)
		_node.Name = value
	}
	if value, ok := nc.mutation.NameIv(); ok {
		_spec.SetField(note.FieldNameIv, field.TypeBytes, value)
		_node.NameIv = value
	}
	if value, ok := nc.mutation.Title(); ok {
		_spec.SetField(note.FieldTitle, field.TypeBytes, value)
		_node.Title = value
	}
	if value, ok := nc.mutation.TitleIv(); ok {
		_spec.SetField(note.FieldTitleIv, field.TypeBytes, value)
		_node.TitleIv = value
	}
	if value, ok := nc.mutation.Content(); ok {
		_spec.SetField(note.FieldContent, field.TypeBytes, value)
		_node.Content = value
	}
	if value, ok := nc.mutation.ContentIv(); ok {
		_spec.SetField(note.FieldContentIv, field.TypeBytes, value)
		_node.ContentIv = value
	}
	if value, ok := nc.mutation.Colour(); ok {
		_spec.SetField(note.FieldColour, field.TypeBytes, value)
		_node.Colour = value
	}
	if value, ok := nc.mutation.ColourIv(); ok {
		_spec.SetField(note.FieldColourIv, field.TypeBytes, value)
		_node.ColourIv = value
	}
	if nodes := nc.mutation.VaultIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   note.VaultTable,
			Columns: []string{note.VaultColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: vault.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.vault_notes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NoteCreateBulk is the builder for creating many Note entities in bulk.
type NoteCreateBulk struct {
	config
	builders []*NoteCreate
}

// Save creates the Note entities in the database.
func (ncb *NoteCreateBulk) Save(ctx context.Context) ([]*Note, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Note, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NoteMutation)
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
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NoteCreateBulk) SaveX(ctx context.Context) []*Note {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ncb *NoteCreateBulk) Exec(ctx context.Context) error {
	_, err := ncb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ncb *NoteCreateBulk) ExecX(ctx context.Context) {
	if err := ncb.Exec(ctx); err != nil {
		panic(err)
	}
}