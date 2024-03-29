// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/webauthnchallenge"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebAuthnChallengeCreate is the builder for creating a WebAuthnChallenge entity.
type WebAuthnChallengeCreate struct {
	config
	mutation *WebAuthnChallengeMutation
	hooks    []Hook
}

// SetSdChallenge sets the "sdChallenge" field.
func (wacc *WebAuthnChallengeCreate) SetSdChallenge(s string) *WebAuthnChallengeCreate {
	wacc.mutation.SetSdChallenge(s)
	return wacc
}

// SetNillableSdChallenge sets the "sdChallenge" field if the given value is not nil.
func (wacc *WebAuthnChallengeCreate) SetNillableSdChallenge(s *string) *WebAuthnChallengeCreate {
	if s != nil {
		wacc.SetSdChallenge(*s)
	}
	return wacc
}

// SetUserId sets the "userId" field.
func (wacc *WebAuthnChallengeCreate) SetUserId(b []byte) *WebAuthnChallengeCreate {
	wacc.mutation.SetUserId(b)
	return wacc
}

// SetAllowedCredentialIds sets the "allowedCredentialIds" field.
func (wacc *WebAuthnChallengeCreate) SetAllowedCredentialIds(u [][]uint8) *WebAuthnChallengeCreate {
	wacc.mutation.SetAllowedCredentialIds(u)
	return wacc
}

// SetUserVerification sets the "userVerification" field.
func (wacc *WebAuthnChallengeCreate) SetUserVerification(s string) *WebAuthnChallengeCreate {
	wacc.mutation.SetUserVerification(s)
	return wacc
}

// SetNillableUserVerification sets the "userVerification" field if the given value is not nil.
func (wacc *WebAuthnChallengeCreate) SetNillableUserVerification(s *string) *WebAuthnChallengeCreate {
	if s != nil {
		wacc.SetUserVerification(*s)
	}
	return wacc
}

// SetExtensions sets the "extensions" field.
func (wacc *WebAuthnChallengeCreate) SetExtensions(m map[string]interface{}) *WebAuthnChallengeCreate {
	wacc.mutation.SetExtensions(m)
	return wacc
}

// SetID sets the "id" field.
func (wacc *WebAuthnChallengeCreate) SetID(u uuid.UUID) *WebAuthnChallengeCreate {
	wacc.mutation.SetID(u)
	return wacc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wacc *WebAuthnChallengeCreate) SetNillableID(u *uuid.UUID) *WebAuthnChallengeCreate {
	if u != nil {
		wacc.SetID(*u)
	}
	return wacc
}

// SetChallengeID sets the "challenge" edge to the Challenge entity by ID.
func (wacc *WebAuthnChallengeCreate) SetChallengeID(id uuid.UUID) *WebAuthnChallengeCreate {
	wacc.mutation.SetChallengeID(id)
	return wacc
}

// SetNillableChallengeID sets the "challenge" edge to the Challenge entity by ID if the given value is not nil.
func (wacc *WebAuthnChallengeCreate) SetNillableChallengeID(id *uuid.UUID) *WebAuthnChallengeCreate {
	if id != nil {
		wacc = wacc.SetChallengeID(*id)
	}
	return wacc
}

// SetChallenge sets the "challenge" edge to the Challenge entity.
func (wacc *WebAuthnChallengeCreate) SetChallenge(c *Challenge) *WebAuthnChallengeCreate {
	return wacc.SetChallengeID(c.ID)
}

// Mutation returns the WebAuthnChallengeMutation object of the builder.
func (wacc *WebAuthnChallengeCreate) Mutation() *WebAuthnChallengeMutation {
	return wacc.mutation
}

// Save creates the WebAuthnChallenge in the database.
func (wacc *WebAuthnChallengeCreate) Save(ctx context.Context) (*WebAuthnChallenge, error) {
	wacc.defaults()
	return withHooks[*WebAuthnChallenge, WebAuthnChallengeMutation](ctx, wacc.sqlSave, wacc.mutation, wacc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wacc *WebAuthnChallengeCreate) SaveX(ctx context.Context) *WebAuthnChallenge {
	v, err := wacc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wacc *WebAuthnChallengeCreate) Exec(ctx context.Context) error {
	_, err := wacc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wacc *WebAuthnChallengeCreate) ExecX(ctx context.Context) {
	if err := wacc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wacc *WebAuthnChallengeCreate) defaults() {
	if _, ok := wacc.mutation.ID(); !ok {
		v := webauthnchallenge.DefaultID()
		wacc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wacc *WebAuthnChallengeCreate) check() error {
	return nil
}

func (wacc *WebAuthnChallengeCreate) sqlSave(ctx context.Context) (*WebAuthnChallenge, error) {
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

func (wacc *WebAuthnChallengeCreate) createSpec() (*WebAuthnChallenge, *sqlgraph.CreateSpec) {
	var (
		_node = &WebAuthnChallenge{config: wacc.config}
		_spec = sqlgraph.NewCreateSpec(webauthnchallenge.Table, sqlgraph.NewFieldSpec(webauthnchallenge.FieldID, field.TypeUUID))
	)
	if id, ok := wacc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wacc.mutation.SdChallenge(); ok {
		_spec.SetField(webauthnchallenge.FieldSdChallenge, field.TypeString, value)
		_node.SdChallenge = value
	}
	if value, ok := wacc.mutation.UserId(); ok {
		_spec.SetField(webauthnchallenge.FieldUserId, field.TypeBytes, value)
		_node.UserId = value
	}
	if value, ok := wacc.mutation.AllowedCredentialIds(); ok {
		_spec.SetField(webauthnchallenge.FieldAllowedCredentialIds, field.TypeJSON, value)
		_node.AllowedCredentialIds = value
	}
	if value, ok := wacc.mutation.UserVerification(); ok {
		_spec.SetField(webauthnchallenge.FieldUserVerification, field.TypeString, value)
		_node.UserVerification = value
	}
	if value, ok := wacc.mutation.Extensions(); ok {
		_spec.SetField(webauthnchallenge.FieldExtensions, field.TypeJSON, value)
		_node.Extensions = value
	}
	if nodes := wacc.mutation.ChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   webauthnchallenge.ChallengeTable,
			Columns: []string{webauthnchallenge.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(challenge.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.challenge_webauthn_challenge = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebAuthnChallengeCreateBulk is the builder for creating many WebAuthnChallenge entities in bulk.
type WebAuthnChallengeCreateBulk struct {
	config
	builders []*WebAuthnChallengeCreate
}

// Save creates the WebAuthnChallenge entities in the database.
func (waccb *WebAuthnChallengeCreateBulk) Save(ctx context.Context) ([]*WebAuthnChallenge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(waccb.builders))
	nodes := make([]*WebAuthnChallenge, len(waccb.builders))
	mutators := make([]Mutator, len(waccb.builders))
	for i := range waccb.builders {
		func(i int, root context.Context) {
			builder := waccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebAuthnChallengeMutation)
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
func (waccb *WebAuthnChallengeCreateBulk) SaveX(ctx context.Context) []*WebAuthnChallenge {
	v, err := waccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (waccb *WebAuthnChallengeCreateBulk) Exec(ctx context.Context) error {
	_, err := waccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (waccb *WebAuthnChallengeCreateBulk) ExecX(ctx context.Context) {
	if err := waccb.Exec(ctx); err != nil {
		panic(err)
	}
}
