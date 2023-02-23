// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/password"
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/user"
	"PasswordManager/ent/vault"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VaultQuery is the builder for querying Vault entities.
type VaultQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.Vault
	withPasswords *PasswordQuery
	withUser      *UserQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VaultQuery builder.
func (vq *VaultQuery) Where(ps ...predicate.Vault) *VaultQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VaultQuery) Limit(limit int) *VaultQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VaultQuery) Offset(offset int) *VaultQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VaultQuery) Unique(unique bool) *VaultQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VaultQuery) Order(o ...OrderFunc) *VaultQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryPasswords chains the current query on the "passwords" edge.
func (vq *VaultQuery) QueryPasswords() *PasswordQuery {
	query := (&PasswordClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vault.Table, vault.FieldID, selector),
			sqlgraph.To(password.Table, password.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, vault.PasswordsTable, vault.PasswordsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (vq *VaultQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(vault.Table, vault.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, vault.UserTable, vault.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Vault entity from the query.
// Returns a *NotFoundError when no Vault was found.
func (vq *VaultQuery) First(ctx context.Context) (*Vault, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vault.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VaultQuery) FirstX(ctx context.Context) *Vault {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Vault ID from the query.
// Returns a *NotFoundError when no Vault ID was found.
func (vq *VaultQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vault.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VaultQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Vault entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Vault entity is found.
// Returns a *NotFoundError when no Vault entities are found.
func (vq *VaultQuery) Only(ctx context.Context) (*Vault, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vault.Label}
	default:
		return nil, &NotSingularError{vault.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VaultQuery) OnlyX(ctx context.Context) *Vault {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Vault ID in the query.
// Returns a *NotSingularError when more than one Vault ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VaultQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vault.Label}
	default:
		err = &NotSingularError{vault.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VaultQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vaults.
func (vq *VaultQuery) All(ctx context.Context) ([]*Vault, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Vault, *VaultQuery]()
	return withInterceptors[[]*Vault](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VaultQuery) AllX(ctx context.Context) []*Vault {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Vault IDs.
func (vq *VaultQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err := vq.Select(vault.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VaultQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VaultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VaultQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VaultQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VaultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VaultQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VaultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VaultQuery) Clone() *VaultQuery {
	if vq == nil {
		return nil
	}
	return &VaultQuery{
		config:        vq.config,
		ctx:           vq.ctx.Clone(),
		order:         append([]OrderFunc{}, vq.order...),
		inters:        append([]Interceptor{}, vq.inters...),
		predicates:    append([]predicate.Vault{}, vq.predicates...),
		withPasswords: vq.withPasswords.Clone(),
		withUser:      vq.withUser.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithPasswords tells the query-builder to eager-load the nodes that are connected to
// the "passwords" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VaultQuery) WithPasswords(opts ...func(*PasswordQuery)) *VaultQuery {
	query := (&PasswordClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withPasswords = query
	return vq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VaultQuery) WithUser(opts ...func(*UserQuery)) *VaultQuery {
	query := (&UserClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withUser = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Vault.Query().
//		GroupBy(vault.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vq *VaultQuery) GroupBy(field string, fields ...string) *VaultGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VaultGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = vault.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.Vault.Query().
//		Select(vault.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (vq *VaultQuery) Select(fields ...string) *VaultSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VaultSelect{VaultQuery: vq}
	sbuild.label = vault.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VaultSelect configured with the given aggregations.
func (vq *VaultQuery) Aggregate(fns ...AggregateFunc) *VaultSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VaultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !vault.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VaultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Vault, error) {
	var (
		nodes       = []*Vault{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withPasswords != nil,
			vq.withUser != nil,
		}
	)
	if vq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, vault.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Vault).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Vault{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withPasswords; query != nil {
		if err := vq.loadPasswords(ctx, query, nodes,
			func(n *Vault) { n.Edges.Passwords = []*Password{} },
			func(n *Vault, e *Password) { n.Edges.Passwords = append(n.Edges.Passwords, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withUser; query != nil {
		if err := vq.loadUser(ctx, query, nodes, nil,
			func(n *Vault, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VaultQuery) loadPasswords(ctx context.Context, query *PasswordQuery, nodes []*Vault, init func(*Vault), assign func(*Vault, *Password)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Vault)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Password(func(s *sql.Selector) {
		s.Where(sql.InValues(vault.PasswordsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.vault_passwords
		if fk == nil {
			return fmt.Errorf(`foreign-key "vault_passwords" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "vault_passwords" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VaultQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Vault, init func(*Vault), assign func(*Vault, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Vault)
	for i := range nodes {
		if nodes[i].user_vaults == nil {
			continue
		}
		fk := *nodes[i].user_vaults
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_vaults" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vq *VaultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VaultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vault.Table,
			Columns: vault.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vault.FieldID,
			},
		},
		From:   vq.sql,
		Unique: true,
	}
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vault.FieldID)
		for i := range fields {
			if fields[i] != vault.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VaultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(vault.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = vault.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VaultGroupBy is the group-by builder for Vault entities.
type VaultGroupBy struct {
	selector
	build *VaultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VaultGroupBy) Aggregate(fns ...AggregateFunc) *VaultGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VaultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VaultQuery, *VaultGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VaultGroupBy) sqlScan(ctx context.Context, root *VaultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VaultSelect is the builder for selecting fields of Vault entities.
type VaultSelect struct {
	*VaultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VaultSelect) Aggregate(fns ...AggregateFunc) *VaultSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VaultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VaultQuery, *VaultSelect](ctx, vs.VaultQuery, vs, vs.inters, v)
}

func (vs *VaultSelect) sqlScan(ctx context.Context, root *VaultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
