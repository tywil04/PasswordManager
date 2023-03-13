// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnregisterchallenge"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebAuthnRegisterChallengeCreate is the builder for creating a WebAuthnRegisterChallenge entity.
type WebAuthnRegisterChallengeCreate struct {
	config
	mutation *WebAuthnRegisterChallengeMutation
	hooks    []Hook
}

// SetSdChallenge sets the "sdChallenge" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetSdChallenge(s string) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetSdChallenge(s)
	return warcc
}

// SetNillableSdChallenge sets the "sdChallenge" field if the given value is not nil.
func (warcc *WebAuthnRegisterChallengeCreate) SetNillableSdChallenge(s *string) *WebAuthnRegisterChallengeCreate {
	if s != nil {
		warcc.SetSdChallenge(*s)
	}
	return warcc
}

// SetUserId sets the "userId" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetUserId(b []byte) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetUserId(b)
	return warcc
}

// SetAllowedCredentialIds sets the "allowedCredentialIds" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetAllowedCredentialIds(u [][]uint8) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetAllowedCredentialIds(u)
	return warcc
}

// SetUserVerification sets the "userVerification" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetUserVerification(s string) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetUserVerification(s)
	return warcc
}

// SetNillableUserVerification sets the "userVerification" field if the given value is not nil.
func (warcc *WebAuthnRegisterChallengeCreate) SetNillableUserVerification(s *string) *WebAuthnRegisterChallengeCreate {
	if s != nil {
		warcc.SetUserVerification(*s)
	}
	return warcc
}

// SetExtensions sets the "extensions" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetExtensions(m map[string]interface{}) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetExtensions(m)
	return warcc
}

// SetID sets the "id" field.
func (warcc *WebAuthnRegisterChallengeCreate) SetID(u uuid.UUID) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetID(u)
	return warcc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (warcc *WebAuthnRegisterChallengeCreate) SetNillableID(u *uuid.UUID) *WebAuthnRegisterChallengeCreate {
	if u != nil {
		warcc.SetID(*u)
	}
	return warcc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (warcc *WebAuthnRegisterChallengeCreate) SetUserID(id uuid.UUID) *WebAuthnRegisterChallengeCreate {
	warcc.mutation.SetUserID(id)
	return warcc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (warcc *WebAuthnRegisterChallengeCreate) SetNillableUserID(id *uuid.UUID) *WebAuthnRegisterChallengeCreate {
	if id != nil {
		warcc = warcc.SetUserID(*id)
	}
	return warcc
}

// SetUser sets the "user" edge to the User entity.
func (warcc *WebAuthnRegisterChallengeCreate) SetUser(u *User) *WebAuthnRegisterChallengeCreate {
	return warcc.SetUserID(u.ID)
}

// Mutation returns the WebAuthnRegisterChallengeMutation object of the builder.
func (warcc *WebAuthnRegisterChallengeCreate) Mutation() *WebAuthnRegisterChallengeMutation {
	return warcc.mutation
}

// Save creates the WebAuthnRegisterChallenge in the database.
func (warcc *WebAuthnRegisterChallengeCreate) Save(ctx context.Context) (*WebAuthnRegisterChallenge, error) {
	warcc.defaults()
	return withHooks[*WebAuthnRegisterChallenge, WebAuthnRegisterChallengeMutation](ctx, warcc.sqlSave, warcc.mutation, warcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (warcc *WebAuthnRegisterChallengeCreate) SaveX(ctx context.Context) *WebAuthnRegisterChallenge {
	v, err := warcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (warcc *WebAuthnRegisterChallengeCreate) Exec(ctx context.Context) error {
	_, err := warcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (warcc *WebAuthnRegisterChallengeCreate) ExecX(ctx context.Context) {
	if err := warcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (warcc *WebAuthnRegisterChallengeCreate) defaults() {
	if _, ok := warcc.mutation.ID(); !ok {
		v := webauthnregisterchallenge.DefaultID()
		warcc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (warcc *WebAuthnRegisterChallengeCreate) check() error {
	return nil
}

func (warcc *WebAuthnRegisterChallengeCreate) sqlSave(ctx context.Context) (*WebAuthnRegisterChallenge, error) {
	if err := warcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := warcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, warcc.driver, _spec); err != nil {
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
	warcc.mutation.id = &_node.ID
	warcc.mutation.done = true
	return _node, nil
}

func (warcc *WebAuthnRegisterChallengeCreate) createSpec() (*WebAuthnRegisterChallenge, *sqlgraph.CreateSpec) {
	var (
		_node = &WebAuthnRegisterChallenge{config: warcc.config}
		_spec = sqlgraph.NewCreateSpec(webauthnregisterchallenge.Table, sqlgraph.NewFieldSpec(webauthnregisterchallenge.FieldID, field.TypeUUID))
	)
	if id, ok := warcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := warcc.mutation.SdChallenge(); ok {
		_spec.SetField(webauthnregisterchallenge.FieldSdChallenge, field.TypeString, value)
		_node.SdChallenge = value
	}
	if value, ok := warcc.mutation.UserId(); ok {
		_spec.SetField(webauthnregisterchallenge.FieldUserId, field.TypeBytes, value)
		_node.UserId = value
	}
	if value, ok := warcc.mutation.AllowedCredentialIds(); ok {
		_spec.SetField(webauthnregisterchallenge.FieldAllowedCredentialIds, field.TypeJSON, value)
		_node.AllowedCredentialIds = value
	}
	if value, ok := warcc.mutation.UserVerification(); ok {
		_spec.SetField(webauthnregisterchallenge.FieldUserVerification, field.TypeString, value)
		_node.UserVerification = value
	}
	if value, ok := warcc.mutation.Extensions(); ok {
		_spec.SetField(webauthnregisterchallenge.FieldExtensions, field.TypeJSON, value)
		_node.Extensions = value
	}
	if nodes := warcc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   webauthnregisterchallenge.UserTable,
			Columns: []string{webauthnregisterchallenge.UserColumn},
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
		_node.user_webauthn_register_challenges = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// WebAuthnRegisterChallengeCreateBulk is the builder for creating many WebAuthnRegisterChallenge entities in bulk.
type WebAuthnRegisterChallengeCreateBulk struct {
	config
	builders []*WebAuthnRegisterChallengeCreate
}

// Save creates the WebAuthnRegisterChallenge entities in the database.
func (warccb *WebAuthnRegisterChallengeCreateBulk) Save(ctx context.Context) ([]*WebAuthnRegisterChallenge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(warccb.builders))
	nodes := make([]*WebAuthnRegisterChallenge, len(warccb.builders))
	mutators := make([]Mutator, len(warccb.builders))
	for i := range warccb.builders {
		func(i int, root context.Context) {
			builder := warccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WebAuthnRegisterChallengeMutation)
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
					_, err = mutators[i+1].Mutate(root, warccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, warccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, warccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (warccb *WebAuthnRegisterChallengeCreateBulk) SaveX(ctx context.Context) []*WebAuthnRegisterChallenge {
	v, err := warccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (warccb *WebAuthnRegisterChallengeCreateBulk) Exec(ctx context.Context) error {
	_, err := warccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (warccb *WebAuthnRegisterChallengeCreateBulk) ExecX(ctx context.Context) {
	if err := warccb.Exec(ctx); err != nil {
		panic(err)
	}
}
