// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/totpcredential"
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnchallenge"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ChallengeUpdate is the builder for updating Challenge entities.
type ChallengeUpdate struct {
	config
	hooks    []Hook
	mutation *ChallengeMutation
}

// Where appends a list predicates to the ChallengeUpdate builder.
func (cu *ChallengeUpdate) Where(ps ...predicate.Challenge) *ChallengeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetExpiry sets the "expiry" field.
func (cu *ChallengeUpdate) SetExpiry(t time.Time) *ChallengeUpdate {
	cu.mutation.SetExpiry(t)
	return cu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableExpiry(t *time.Time) *ChallengeUpdate {
	if t != nil {
		cu.SetExpiry(*t)
	}
	return cu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cu *ChallengeUpdate) SetUserID(id uuid.UUID) *ChallengeUpdate {
	cu.mutation.SetUserID(id)
	return cu
}

// SetUser sets the "user" edge to the User entity.
func (cu *ChallengeUpdate) SetUser(u *User) *ChallengeUpdate {
	return cu.SetUserID(u.ID)
}

// SetEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID.
func (cu *ChallengeUpdate) SetEmailChallengeID(id uuid.UUID) *ChallengeUpdate {
	cu.mutation.SetEmailChallengeID(id)
	return cu
}

// SetNillableEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableEmailChallengeID(id *uuid.UUID) *ChallengeUpdate {
	if id != nil {
		cu = cu.SetEmailChallengeID(*id)
	}
	return cu
}

// SetEmailChallenge sets the "emailChallenge" edge to the EmailChallenge entity.
func (cu *ChallengeUpdate) SetEmailChallenge(e *EmailChallenge) *ChallengeUpdate {
	return cu.SetEmailChallengeID(e.ID)
}

// SetWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID.
func (cu *ChallengeUpdate) SetWebauthnChallengeID(id uuid.UUID) *ChallengeUpdate {
	cu.mutation.SetWebauthnChallengeID(id)
	return cu
}

// SetNillableWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableWebauthnChallengeID(id *uuid.UUID) *ChallengeUpdate {
	if id != nil {
		cu = cu.SetWebauthnChallengeID(*id)
	}
	return cu
}

// SetWebauthnChallenge sets the "webauthnChallenge" edge to the WebAuthnChallenge entity.
func (cu *ChallengeUpdate) SetWebauthnChallenge(w *WebAuthnChallenge) *ChallengeUpdate {
	return cu.SetWebauthnChallengeID(w.ID)
}

// SetTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID.
func (cu *ChallengeUpdate) SetTotpCredentialID(id uuid.UUID) *ChallengeUpdate {
	cu.mutation.SetTotpCredentialID(id)
	return cu
}

// SetNillableTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableTotpCredentialID(id *uuid.UUID) *ChallengeUpdate {
	if id != nil {
		cu = cu.SetTotpCredentialID(*id)
	}
	return cu
}

// SetTotpCredential sets the "totpCredential" edge to the TotpCredential entity.
func (cu *ChallengeUpdate) SetTotpCredential(t *TotpCredential) *ChallengeUpdate {
	return cu.SetTotpCredentialID(t.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cu *ChallengeUpdate) Mutation() *ChallengeMutation {
	return cu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cu *ChallengeUpdate) ClearUser() *ChallengeUpdate {
	cu.mutation.ClearUser()
	return cu
}

// ClearEmailChallenge clears the "emailChallenge" edge to the EmailChallenge entity.
func (cu *ChallengeUpdate) ClearEmailChallenge() *ChallengeUpdate {
	cu.mutation.ClearEmailChallenge()
	return cu
}

// ClearWebauthnChallenge clears the "webauthnChallenge" edge to the WebAuthnChallenge entity.
func (cu *ChallengeUpdate) ClearWebauthnChallenge() *ChallengeUpdate {
	cu.mutation.ClearWebauthnChallenge()
	return cu
}

// ClearTotpCredential clears the "totpCredential" edge to the TotpCredential entity.
func (cu *ChallengeUpdate) ClearTotpCredential() *ChallengeUpdate {
	cu.mutation.ClearTotpCredential()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChallengeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ChallengeMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChallengeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChallengeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChallengeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChallengeUpdate) check() error {
	if _, ok := cu.mutation.UserID(); cu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Challenge.user"`)
	}
	return nil
}

func (cu *ChallengeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Expiry(); ok {
		_spec.SetField(challenge.FieldExpiry, field.TypeTime, value)
	}
	if cu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.EmailChallengeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EmailChallengeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.WebauthnChallengeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.WebauthnChallengeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.TotpCredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.TotpCredentialIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ChallengeUpdateOne is the builder for updating a single Challenge entity.
type ChallengeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChallengeMutation
}

// SetExpiry sets the "expiry" field.
func (cuo *ChallengeUpdateOne) SetExpiry(t time.Time) *ChallengeUpdateOne {
	cuo.mutation.SetExpiry(t)
	return cuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableExpiry(t *time.Time) *ChallengeUpdateOne {
	if t != nil {
		cuo.SetExpiry(*t)
	}
	return cuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (cuo *ChallengeUpdateOne) SetUserID(id uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.SetUserID(id)
	return cuo
}

// SetUser sets the "user" edge to the User entity.
func (cuo *ChallengeUpdateOne) SetUser(u *User) *ChallengeUpdateOne {
	return cuo.SetUserID(u.ID)
}

// SetEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID.
func (cuo *ChallengeUpdateOne) SetEmailChallengeID(id uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.SetEmailChallengeID(id)
	return cuo
}

// SetNillableEmailChallengeID sets the "emailChallenge" edge to the EmailChallenge entity by ID if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableEmailChallengeID(id *uuid.UUID) *ChallengeUpdateOne {
	if id != nil {
		cuo = cuo.SetEmailChallengeID(*id)
	}
	return cuo
}

// SetEmailChallenge sets the "emailChallenge" edge to the EmailChallenge entity.
func (cuo *ChallengeUpdateOne) SetEmailChallenge(e *EmailChallenge) *ChallengeUpdateOne {
	return cuo.SetEmailChallengeID(e.ID)
}

// SetWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID.
func (cuo *ChallengeUpdateOne) SetWebauthnChallengeID(id uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.SetWebauthnChallengeID(id)
	return cuo
}

// SetNillableWebauthnChallengeID sets the "webauthnChallenge" edge to the WebAuthnChallenge entity by ID if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableWebauthnChallengeID(id *uuid.UUID) *ChallengeUpdateOne {
	if id != nil {
		cuo = cuo.SetWebauthnChallengeID(*id)
	}
	return cuo
}

// SetWebauthnChallenge sets the "webauthnChallenge" edge to the WebAuthnChallenge entity.
func (cuo *ChallengeUpdateOne) SetWebauthnChallenge(w *WebAuthnChallenge) *ChallengeUpdateOne {
	return cuo.SetWebauthnChallengeID(w.ID)
}

// SetTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID.
func (cuo *ChallengeUpdateOne) SetTotpCredentialID(id uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.SetTotpCredentialID(id)
	return cuo
}

// SetNillableTotpCredentialID sets the "totpCredential" edge to the TotpCredential entity by ID if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableTotpCredentialID(id *uuid.UUID) *ChallengeUpdateOne {
	if id != nil {
		cuo = cuo.SetTotpCredentialID(*id)
	}
	return cuo
}

// SetTotpCredential sets the "totpCredential" edge to the TotpCredential entity.
func (cuo *ChallengeUpdateOne) SetTotpCredential(t *TotpCredential) *ChallengeUpdateOne {
	return cuo.SetTotpCredentialID(t.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cuo *ChallengeUpdateOne) Mutation() *ChallengeMutation {
	return cuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cuo *ChallengeUpdateOne) ClearUser() *ChallengeUpdateOne {
	cuo.mutation.ClearUser()
	return cuo
}

// ClearEmailChallenge clears the "emailChallenge" edge to the EmailChallenge entity.
func (cuo *ChallengeUpdateOne) ClearEmailChallenge() *ChallengeUpdateOne {
	cuo.mutation.ClearEmailChallenge()
	return cuo
}

// ClearWebauthnChallenge clears the "webauthnChallenge" edge to the WebAuthnChallenge entity.
func (cuo *ChallengeUpdateOne) ClearWebauthnChallenge() *ChallengeUpdateOne {
	cuo.mutation.ClearWebauthnChallenge()
	return cuo
}

// ClearTotpCredential clears the "totpCredential" edge to the TotpCredential entity.
func (cuo *ChallengeUpdateOne) ClearTotpCredential() *ChallengeUpdateOne {
	cuo.mutation.ClearTotpCredential()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChallengeUpdateOne) Select(field string, fields ...string) *ChallengeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Challenge entity.
func (cuo *ChallengeUpdateOne) Save(ctx context.Context) (*Challenge, error) {
	return withHooks[*Challenge, ChallengeMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) SaveX(ctx context.Context) *Challenge {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChallengeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChallengeUpdateOne) check() error {
	if _, ok := cuo.mutation.UserID(); cuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Challenge.user"`)
	}
	return nil
}

func (cuo *ChallengeUpdateOne) sqlSave(ctx context.Context) (_node *Challenge, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Challenge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, challenge.FieldID)
		for _, f := range fields {
			if !challenge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != challenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Expiry(); ok {
		_spec.SetField(challenge.FieldExpiry, field.TypeTime, value)
	}
	if cuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.EmailChallengeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EmailChallengeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.WebauthnChallengeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.WebauthnChallengeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.TotpCredentialCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.TotpCredentialIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Challenge{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}