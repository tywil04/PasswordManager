// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthncredential"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebAuthnCredentialCreate is the builder for creating a WebAuthnCredential entity.
type WebAuthnCredentialCreate struct {
	config
	mutation *WebAuthnCredentialMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (wacc *WebAuthnCredentialCreate) SetName(s string) *WebAuthnCredentialCreate {
	wacc.mutation.SetName(s)
	return wacc
}

// SetCreatedAt sets the "createdAt" field.
func (wacc *WebAuthnCredentialCreate) SetCreatedAt(t time.Time) *WebAuthnCredentialCreate {
	wacc.mutation.SetCreatedAt(t)
	return wacc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (wacc *WebAuthnCredentialCreate) SetNillableCreatedAt(t *time.Time) *WebAuthnCredentialCreate {
	if t != nil {
		wacc.SetCreatedAt(*t)
	}
	return wacc
}

// SetCredentialId sets the "credentialId" field.
func (wacc *WebAuthnCredentialCreate) SetCredentialId(b []byte) *WebAuthnCredentialCreate {
	wacc.mutation.SetCredentialId(b)
	return wacc
}

// SetPublicKey sets the "publicKey" field.
func (wacc *WebAuthnCredentialCreate) SetPublicKey(b []byte) *WebAuthnCredentialCreate {
	wacc.mutation.SetPublicKey(b)
	return wacc
}

// SetAttestationType sets the "attestationType" field.
func (wacc *WebAuthnCredentialCreate) SetAttestationType(s string) *WebAuthnCredentialCreate {
	wacc.mutation.SetAttestationType(s)
	return wacc
}

// SetTransport sets the "transport" field.
func (wacc *WebAuthnCredentialCreate) SetTransport(s []string) *WebAuthnCredentialCreate {
	wacc.mutation.SetTransport(s)
	return wacc
}

// SetAaguid sets the "aaguid" field.
func (wacc *WebAuthnCredentialCreate) SetAaguid(b []byte) *WebAuthnCredentialCreate {
	wacc.mutation.SetAaguid(b)
	return wacc
}

// SetSignCount sets the "signCount" field.
func (wacc *WebAuthnCredentialCreate) SetSignCount(u uint32) *WebAuthnCredentialCreate {
	wacc.mutation.SetSignCount(u)
	return wacc
}

// SetCloneWarning sets the "cloneWarning" field.
func (wacc *WebAuthnCredentialCreate) SetCloneWarning(b bool) *WebAuthnCredentialCreate {
	wacc.mutation.SetCloneWarning(b)
	return wacc
}

// SetID sets the "id" field.
func (wacc *WebAuthnCredentialCreate) SetID(u uuid.UUID) *WebAuthnCredentialCreate {
	wacc.mutation.SetID(u)
	return wacc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wacc *WebAuthnCredentialCreate) SetNillableID(u *uuid.UUID) *WebAuthnCredentialCreate {
	if u != nil {
		wacc.SetID(*u)
	}
	return wacc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (wacc *WebAuthnCredentialCreate) SetUserID(id uuid.UUID) *WebAuthnCredentialCreate {
	wacc.mutation.SetUserID(id)
	return wacc
}

// SetUser sets the "user" edge to the User entity.
func (wacc *WebAuthnCredentialCreate) SetUser(u *User) *WebAuthnCredentialCreate {
	return wacc.SetUserID(u.ID)
}

// Mutation returns the WebAuthnCredentialMutation object of the builder.
func (wacc *WebAuthnCredentialCreate) Mutation() *WebAuthnCredentialMutation {
	return wacc.mutation
}

// Save creates the WebAuthnCredential in the database.
func (wacc *WebAuthnCredentialCreate) Save(ctx context.Context) (*WebAuthnCredential, error) {
	wacc.defaults()
	return withHooks[*WebAuthnCredential, WebAuthnCredentialMutation](ctx, wacc.sqlSave, wacc.mutation, wacc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wacc *WebAuthnCredentialCreate) SaveX(ctx context.Context) *WebAuthnCredential {
	v, err := wacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wacc *WebAuthnCredentialCreate) Exec(ctx context.Context) error {
	_, err := wacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacc *WebAuthnCredentialCreate) ExecX(ctx context.Context) {
	if err := wacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wacc *WebAuthnCredentialCreate) defaults() {
	if _, ok := wacc.mutation.CreatedAt(); !ok {
		v := webauthncredential.DefaultCreatedAt()
		wacc.mutation.SetCreatedAt(v)
	}
	if _, ok := wacc.mutation.ID(); !ok {
		v := webauthncredential.DefaultID()
		wacc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wacc *WebAuthnCredentialCreate) check() error {
	if _, ok := wacc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "WebAuthnCredential.name"`)}
	}
	if _, ok := wacc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "WebAuthnCredential.createdAt"`)}
	}
	if _, ok := wacc.mutation.CredentialId(); !ok {
		return &ValidationError{Name: "credentialId", err: errors.New(`ent: missing required field "WebAuthnCredential.credentialId"`)}
	}
	if _, ok := wacc.mutation.PublicKey(); !ok {
		return &ValidationError{Name: "publicKey", err: errors.New(`ent: missing required field "WebAuthnCredential.publicKey"`)}
	}
	if _, ok := wacc.mutation.AttestationType(); !ok {
		return &ValidationError{Name: "attestationType", err: errors.New(`ent: missing required field "WebAuthnCredential.attestationType"`)}
	}
	if _, ok := wacc.mutation.Transport(); !ok {
		return &ValidationError{Name: "transport", err: errors.New(`ent: missing required field "WebAuthnCredential.transport"`)}
	}
	if _, ok := wacc.mutation.Aaguid(); !ok {
		return &ValidationError{Name: "aaguid", err: errors.New(`ent: missing required field "WebAuthnCredential.aaguid"`)}
	}
	if _, ok := wacc.mutation.SignCount(); !ok {
		return &ValidationError{Name: "signCount", err: errors.New(`ent: missing required field "WebAuthnCredential.signCount"`)}
	}
	if _, ok := wacc.mutation.CloneWarning(); !ok {
		return &ValidationError{Name: "cloneWarning", err: errors.New(`ent: missing required field "WebAuthnCredential.cloneWarning"`)}
	}
	if _, ok := wacc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "WebAuthnCredential.user"`)}
	}
	return nil
}

func (wacc *WebAuthnCredentialCreate) sqlSave(ctx context.Context) (*WebAuthnCredential, error) {
	if err := wacc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wacc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wacc.driver, _spec); err != nil {
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
	wacc.mutation.id = &_node.ID
	wacc.mutation.done = true
	return _node, nil
}

func (wacc *WebAuthnCredentialCreate) createSpec() (*WebAuthnCredential, *sqlgraph.CreateSpec) {
	var (
		_node = &WebAuthnCredential{config: wacc.config}
		_spec = sqlgraph.NewCreateSpec(webauthncredential.Table, sqlgraph.NewFieldSpec(webauthncredential.FieldID, field.TypeUUID))
	)
	if id, ok := wacc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wacc.mutation.Name(); ok {
		_spec.SetField(webauthncredential.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := wacc.mutation.CreatedAt(); ok {
		_spec.SetField(webauthncredential.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := wacc.mutation.CredentialId(); ok {
		_spec.SetField(webauthncredential.FieldCredentialId, field.TypeBytes, value)
		_node.CredentialId = value
	}
	if value, ok := wacc.mutation.PublicKey(); ok {
		_spec.SetField(webauthncredential.FieldPublicKey, field.TypeBytes, value)
		_node.PublicKey = value
	}
	if value, ok := wacc.mutation.AttestationType(); ok {
		_spec.SetField(webauthncredential.FieldAttestationType, field.TypeString, value)
		_node.AttestationType = value
	}
	if value, ok := wacc.mutation.Transport(); ok {
		_spec.SetField(webauthncredential.FieldTransport, field.TypeJSON, value)
		_node.Transport = value
	}
	if value, ok := wacc.mutation.Aaguid(); ok {
		_spec.SetField(webauthncredential.FieldAaguid, field.TypeBytes, value)
		_node.Aaguid = value
	}
	if value, ok := wacc.mutation.SignCount(); ok {
		_spec.SetField(webauthncredential.FieldSignCount, field.TypeUint32, value)
		_node.SignCount = value
	}
	if value, ok := wacc.mutation.CloneWarning(); ok {
		_spec.SetField(webauthncredential.FieldCloneWarning, field.TypeBool, value)
		_node.CloneWarning = value
	}
	if nodes := wacc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthncredential.UserTable,
			Columns: []string{webauthncredential.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_webauthn_credentials = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebAuthnCredentialCreateBulk is the builder for creating many WebAuthnCredential entities in bulk.
type WebAuthnCredentialCreateBulk struct {
	config
	builders []*WebAuthnCredentialCreate
}

// Save creates the WebAuthnCredential entities in the database.
func (waccb *WebAuthnCredentialCreateBulk) Save(ctx context.Context) ([]*WebAuthnCredential, error) {
	specs := make([]*sqlgraph.CreateSpec, len(waccb.builders))
	nodes := make([]*WebAuthnCredential, len(waccb.builders))
	mutators := make([]Mutator, len(waccb.builders))
	for i := range waccb.builders {
		func(i int, root context.Context) {
			builder := waccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebAuthnCredentialMutation)
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
					_, err = mutators[i+1].Mutate(root, waccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, waccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, waccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (waccb *WebAuthnCredentialCreateBulk) SaveX(ctx context.Context) []*WebAuthnCredential {
	v, err := waccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (waccb *WebAuthnCredentialCreateBulk) Exec(ctx context.Context) error {
	_, err := waccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (waccb *WebAuthnCredentialCreateBulk) ExecX(ctx context.Context) {
	if err := waccb.Exec(ctx); err != nil {
		panic(err)
	}
}
