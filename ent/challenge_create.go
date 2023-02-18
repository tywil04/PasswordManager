// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/totpcredential"
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnchallenge"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ChallengeCreate is the builder for creating a Challenge entity.
type ChallengeCreate struct {
	config
	mutation *ChallengeMutation
	hooks    []Hook
}

// SetExpiry sets the "expiry" field.
func (cc *ChallengeCreate) SetExpiry(t time.Time) *ChallengeCreate {
	cc.mutation.SetExpiry(t)
	return cc
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableExpiry(t *time.Time) *ChallengeCreate {
	if t != nil {
		cc.SetExpiry(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChallengeCreate) SetID(u uuid.UUID) *ChallengeCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChallengeCreate) SetNillableID(u *uuid.UUID) *ChallengeCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cc *ChallengeCreate) SetUserID(id uuid.UUID) *ChallengeCreate {
	cc.mutation.SetUserID(id)
	return cc
}

// SetUser sets the "user" edge to the User entity.
func (cc *ChallengeCreate) SetUser(u *User) *ChallengeCreate {
	return cc.SetUserID(u.ID)
}

// SetEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID.
func (cc *ChallengeCreate) SetEmailChallengeID(id uuid.UUID) *ChallengeCreate {
	cc.mutation.SetEmailChallengeID(id)
	return cc
}

// SetNillableEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID if the given value is not nil.
func (cc *ChallengeCreate) SetNillableEmailChallengeID(id *uuid.UUID) *ChallengeCreate {
	if id != nil {
		cc = cc.SetEmailChallengeID(*id)
	}
	return cc
}

// SetEmailChallenge sets the "emailChallenge" edge to the EmailChallenge entity.
func (cc *ChallengeCreate) SetEmailChallenge(e *EmailChallenge) *ChallengeCreate {
	return cc.SetEmailChallengeID(e.ID)
}

// SetWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID.
func (cc *ChallengeCreate) SetWebauthnChallengeID(id uuid.UUID) *ChallengeCreate {
	cc.mutation.SetWebauthnChallengeID(id)
	return cc
}

// SetNillableWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID if the given value is not nil.
func (cc *ChallengeCreate) SetNillableWebauthnChallengeID(id *uuid.UUID) *ChallengeCreate {
	if id != nil {
		cc = cc.SetWebauthnChallengeID(*id)
	}
	return cc
}

// SetWebauthnChallenge sets the "webauthnChallenge" edge to the WebAuthnChallenge entity.
func (cc *ChallengeCreate) SetWebauthnChallenge(w *WebAuthnChallenge) *ChallengeCreate {
	return cc.SetWebauthnChallengeID(w.ID)
}

// SetTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID.
func (cc *ChallengeCreate) SetTotpCredentialID(id uuid.UUID) *ChallengeCreate {
	cc.mutation.SetTotpCredentialID(id)
	return cc
}

// SetNillableTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID if the given value is not nil.
func (cc *ChallengeCreate) SetNillableTotpCredentialID(id *uuid.UUID) *ChallengeCreate {
	if id != nil {
		cc = cc.SetTotpCredentialID(*id)
	}
	return cc
}

// SetTotpCredential sets the "totpCredential" edge to the TotpCredential entity.
func (cc *ChallengeCreate) SetTotpCredential(t *TotpCredential) *ChallengeCreate {
	return cc.SetTotpCredentialID(t.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cc *ChallengeCreate) Mutation() *ChallengeMutation {
	return cc.mutation
}

// Save creates the Challenge in the database.
func (cc *ChallengeCreate) Save(ctx context.Context) (*Challenge, error) {
	cc.defaults()
	return withHooks[*Challenge, ChallengeMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChallengeCreate) SaveX(ctx context.Context) *Challenge {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChallengeCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChallengeCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChallengeCreate) defaults() {
	if _, ok := cc.mutation.Expiry(); !ok {
		v := challenge.DefaultExpiry()
		cc.mutation.SetExpiry(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := challenge.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChallengeCreate) check() error {
	if _, ok := cc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "Challenge.expiry"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Challenge.user"`)}
	}
	return nil
}

func (cc *ChallengeCreate) sqlSave(ctx context.Context) (*Challenge, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChallengeCreate) createSpec() (*Challenge, *sqlgraph.CreateSpec) {
	var (
		_node = &Challenge{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: challenge.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.Expiry(); ok {
		_spec.SetField(challenge.FieldExpiry, field.TypeTime, value)
		_node.Expiry = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.UserTable,
			Columns: []string{challenge.UserColumn},
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
		_node.user_challenges = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.EmailChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   challenge.EmailChallengeTable,
			Columns: []string{challenge.EmailChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: emailchallenge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.WebauthnChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   challenge.WebauthnChallengeTable,
			Columns: []string{challenge.WebauthnChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: webauthnchallenge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TotpCredentialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   challenge.TotpCredentialTable,
			Columns: []string{challenge.TotpCredentialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: totpcredential.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChallengeCreateBulk is the builder for creating many Challenge entities in bulk.
type ChallengeCreateBulk struct {
	config
	builders []*ChallengeCreate
}

// Save creates the Challenge entities in the database.
func (ccb *ChallengeCreateBulk) Save(ctx context.Context) ([]*Challenge, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Challenge, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChallengeMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) SaveX(ctx context.Context) []*Challenge {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChallengeCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChallengeCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}