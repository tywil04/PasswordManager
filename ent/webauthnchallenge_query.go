// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/predicate"
	"PasswordManager/ent/user"
	"PasswordManager/ent/webauthnchallenge"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// WebAuthnChallengeQuery is the builder for querying WebAuthnChallenge entities.
type WebAuthnChallengeQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.WebAuthnChallenge
	withUser   *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebAuthnChallengeQuery builder.
func (wacq *WebAuthnChallengeQuery) Where(ps ...predicate.WebAuthnChallenge) *WebAuthnChallengeQuery {
	wacq.predicates = append(wacq.predicates, ps...)
	return wacq
}

// Limit the number of records to be returned by this query.
func (wacq *WebAuthnChallengeQuery) Limit(limit int) *WebAuthnChallengeQuery {
	wacq.ctx.Limit = &limit
	return wacq
}

// Offset to start from.
func (wacq *WebAuthnChallengeQuery) Offset(offset int) *WebAuthnChallengeQuery {
	wacq.ctx.Offset = &offset
	return wacq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (wacq *WebAuthnChallengeQuery) Unique(unique bool) *WebAuthnChallengeQuery {
	wacq.ctx.Unique = &unique
	return wacq
}

// Order specifies how the records should be ordered.
func (wacq *WebAuthnChallengeQuery) Order(o ...OrderFunc) *WebAuthnChallengeQuery {
	wacq.order = append(wacq.order, o...)
	return wacq
}

// QueryUser chains the current query on the "user" edge.
func (wacq *WebAuthnChallengeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: wacq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := wacq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := wacq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(webauthnchallenge.Table, webauthnchallenge.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, webauthnchallenge.UserTable, webauthnchallenge.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(wacq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first WebAuthnChallenge entity from the query.
// Returns a *NotFoundError when no WebAuthnChallenge was found.
func (wacq *WebAuthnChallengeQuery) First(ctx context.Context) (*WebAuthnChallenge, error) {
	nodes, err := wacq.Limit(1).All(setContextOp(ctx, wacq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{webauthnchallenge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) FirstX(ctx context.Context) *WebAuthnChallenge {
	node, err := wacq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebAuthnChallenge ID from the query.
// Returns a *NotFoundError when no WebAuthnChallenge ID was found.
func (wacq *WebAuthnChallengeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wacq.Limit(1).IDs(setContextOp(ctx, wacq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webauthnchallenge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := wacq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebAuthnChallenge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebAuthnChallenge entity is found.
// Returns a *NotFoundError when no WebAuthnChallenge entities are found.
func (wacq *WebAuthnChallengeQuery) Only(ctx context.Context) (*WebAuthnChallenge, error) {
	nodes, err := wacq.Limit(2).All(setContextOp(ctx, wacq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{webauthnchallenge.Label}
	default:
		return nil, &NotSingularError{webauthnchallenge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) OnlyX(ctx context.Context) *WebAuthnChallenge {
	node, err := wacq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebAuthnChallenge ID in the query.
// Returns a *NotSingularError when more than one WebAuthnChallenge ID is found.
// Returns a *NotFoundError when no entities are found.
func (wacq *WebAuthnChallengeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = wacq.Limit(2).IDs(setContextOp(ctx, wacq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webauthnchallenge.Label}
	default:
		err = &NotSingularError{webauthnchallenge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := wacq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebAuthnChallenges.
func (wacq *WebAuthnChallengeQuery) All(ctx context.Context) ([]*WebAuthnChallenge, error) {
	ctx = setContextOp(ctx, wacq.ctx, "All")
	if err := wacq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebAuthnChallenge, *WebAuthnChallengeQuery]()
	return withInterceptors[[]*WebAuthnChallenge](ctx, wacq, qr, wacq.inters)
}

// AllX is like All, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) AllX(ctx context.Context) []*WebAuthnChallenge {
	nodes, err := wacq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebAuthnChallenge IDs.
func (wacq *WebAuthnChallengeQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	ctx = setContextOp(ctx, wacq.ctx, "IDs")
	if err := wacq.Select(webauthnchallenge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := wacq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (wacq *WebAuthnChallengeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, wacq.ctx, "Count")
	if err := wacq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, wacq, querierCount[*WebAuthnChallengeQuery](), wacq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) CountX(ctx context.Context) int {
	count, err := wacq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (wacq *WebAuthnChallengeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, wacq.ctx, "Exist")
	switch _, err := wacq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (wacq *WebAuthnChallengeQuery) ExistX(ctx context.Context) bool {
	exist, err := wacq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebAuthnChallengeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (wacq *WebAuthnChallengeQuery) Clone() *WebAuthnChallengeQuery {
	if wacq == nil {
		return nil
	}
	return &WebAuthnChallengeQuery{
		config:     wacq.config,
		ctx:        wacq.ctx.Clone(),
		order:      append([]OrderFunc{}, wacq.order...),
		inters:     append([]Interceptor{}, wacq.inters...),
		predicates: append([]predicate.WebAuthnChallenge{}, wacq.predicates...),
		withUser:   wacq.withUser.Clone(),
		// clone intermediate query.
		sql:  wacq.sql.Clone(),
		path: wacq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (wacq *WebAuthnChallengeQuery) WithUser(opts ...func(*UserQuery)) *WebAuthnChallengeQuery {
	query := (&UserClient{config: wacq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	wacq.withUser = query
	return wacq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Challenge string `json:"challenge,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebAuthnChallenge.Query().
//		GroupBy(webauthnchallenge.FieldChallenge).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (wacq *WebAuthnChallengeQuery) GroupBy(field string, fields ...string) *WebAuthnChallengeGroupBy {
	wacq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebAuthnChallengeGroupBy{build: wacq}
	grbuild.flds = &wacq.ctx.Fields
	grbuild.label = webauthnchallenge.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Challenge string `json:"challenge,omitempty"`
//	}
//
//	client.WebAuthnChallenge.Query().
//		Select(webauthnchallenge.FieldChallenge).
//		Scan(ctx, &v)
//
func (wacq *WebAuthnChallengeQuery) Select(fields ...string) *WebAuthnChallengeSelect {
	wacq.ctx.Fields = append(wacq.ctx.Fields, fields...)
	sbuild := &WebAuthnChallengeSelect{WebAuthnChallengeQuery: wacq}
	sbuild.label = webauthnchallenge.Label
	sbuild.flds, sbuild.scan = &wacq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebAuthnChallengeSelect configured with the given aggregations.
func (wacq *WebAuthnChallengeQuery) Aggregate(fns ...AggregateFunc) *WebAuthnChallengeSelect {
	return wacq.Select().Aggregate(fns...)
}

func (wacq *WebAuthnChallengeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range wacq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, wacq); err != nil {
				return err
			}
		}
	}
	for _, f := range wacq.ctx.Fields {
		if !webauthnchallenge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if wacq.path != nil {
		prev, err := wacq.path(ctx)
		if err != nil {
			return err
		}
		wacq.sql = prev
	}
	return nil
}

func (wacq *WebAuthnChallengeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebAuthnChallenge, error) {
	var (
		nodes       = []*WebAuthnChallenge{}
		withFKs     = wacq.withFKs
		_spec       = wacq.querySpec()
		loadedTypes = [1]bool{
			wacq.withUser != nil,
		}
	)
	if wacq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, webauthnchallenge.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebAuthnChallenge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebAuthnChallenge{config: wacq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, wacq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := wacq.withUser; query != nil {
		if err := wacq.loadUser(ctx, query, nodes, nil,
			func(n *WebAuthnChallenge, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (wacq *WebAuthnChallengeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*WebAuthnChallenge, init func(*WebAuthnChallenge), assign func(*WebAuthnChallenge, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*WebAuthnChallenge)
	for i := range nodes {
		if nodes[i].user_webauthn_challenges == nil {
			continue
		}
		fk := *nodes[i].user_webauthn_challenges
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
			return fmt.Errorf(`unexpected foreign-key "user_webauthn_challenges" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (wacq *WebAuthnChallengeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := wacq.querySpec()
	_spec.Node.Columns = wacq.ctx.Fields
	if len(wacq.ctx.Fields) > 0 {
		_spec.Unique = wacq.ctx.Unique != nil && *wacq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, wacq.driver, _spec)
}

func (wacq *WebAuthnChallengeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   webauthnchallenge.Table,
			Columns: webauthnchallenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: webauthnchallenge.FieldID,
			},
		},
		From:   wacq.sql,
		Unique: true,
	}
	if unique := wacq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := wacq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webauthnchallenge.FieldID)
		for i := range fields {
			if fields[i] != webauthnchallenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := wacq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := wacq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := wacq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := wacq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (wacq *WebAuthnChallengeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(wacq.driver.Dialect())
	t1 := builder.Table(webauthnchallenge.Table)
	columns := wacq.ctx.Fields
	if len(columns) == 0 {
		columns = webauthnchallenge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if wacq.sql != nil {
		selector = wacq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if wacq.ctx.Unique != nil && *wacq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range wacq.predicates {
		p(selector)
	}
	for _, p := range wacq.order {
		p(selector)
	}
	if offset := wacq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := wacq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebAuthnChallengeGroupBy is the group-by builder for WebAuthnChallenge entities.
type WebAuthnChallengeGroupBy struct {
	selector
	build *WebAuthnChallengeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (wacgb *WebAuthnChallengeGroupBy) Aggregate(fns ...AggregateFunc) *WebAuthnChallengeGroupBy {
	wacgb.fns = append(wacgb.fns, fns...)
	return wacgb
}

// Scan applies the selector query and scans the result into the given value.
func (wacgb *WebAuthnChallengeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wacgb.build.ctx, "GroupBy")
	if err := wacgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebAuthnChallengeQuery, *WebAuthnChallengeGroupBy](ctx, wacgb.build, wacgb, wacgb.build.inters, v)
}

func (wacgb *WebAuthnChallengeGroupBy) sqlScan(ctx context.Context, root *WebAuthnChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(wacgb.fns))
	for _, fn := range wacgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*wacgb.flds)+len(wacgb.fns))
		for _, f := range *wacgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*wacgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wacgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebAuthnChallengeSelect is the builder for selecting fields of WebAuthnChallenge entities.
type WebAuthnChallengeSelect struct {
	*WebAuthnChallengeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (wacs *WebAuthnChallengeSelect) Aggregate(fns ...AggregateFunc) *WebAuthnChallengeSelect {
	wacs.fns = append(wacs.fns, fns...)
	return wacs
}

// Scan applies the selector query and scans the result into the given value.
func (wacs *WebAuthnChallengeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, wacs.ctx, "Select")
	if err := wacs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebAuthnChallengeQuery, *WebAuthnChallengeSelect](ctx, wacs.WebAuthnChallengeQuery, wacs, wacs.inters, v)
}

func (wacs *WebAuthnChallengeSelect) sqlScan(ctx context.Context, root *WebAuthnChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(wacs.fns))
	for _, fn := range wacs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*wacs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := wacs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}