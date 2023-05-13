// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/challenge"
	"PasswordManager/ent/emailchallenge"
	"PasswordManager/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EmailChallengeQuery is the builder for querying EmailChallenge entities.
type EmailChallengeQuery struct {
	config
	ctx           *QueryContext
	order         []OrderFunc
	inters        []Interceptor
	predicates    []predicate.EmailChallenge
	withChallenge *ChallengeQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailChallengeQuery builder.
func (ecq *EmailChallengeQuery) Where(ps ...predicate.EmailChallenge) *EmailChallengeQuery {
	ecq.predicates = append(ecq.predicates, ps...)
	return ecq
}

// Limit the number of records to be returned by this query.
func (ecq *EmailChallengeQuery) Limit(limit int) *EmailChallengeQuery {
	ecq.ctx.Limit = &limit
	return ecq
}

// Offset to start from.
func (ecq *EmailChallengeQuery) Offset(offset int) *EmailChallengeQuery {
	ecq.ctx.Offset = &offset
	return ecq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ecq *EmailChallengeQuery) Unique(unique bool) *EmailChallengeQuery {
	ecq.ctx.Unique = &unique
	return ecq
}

// Order specifies how the records should be ordered.
func (ecq *EmailChallengeQuery) Order(o ...OrderFunc) *EmailChallengeQuery {
	ecq.order = append(ecq.order, o...)
	return ecq
}

// QueryChallenge chains the current query on the "challenge" edge.
func (ecq *EmailChallengeQuery) QueryChallenge() *ChallengeQuery {
	query := (&ChallengeClient{config: ecq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ecq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ecq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(emailchallenge.Table, emailchallenge.FieldID, selector),
			sqlgraph.To(challenge.Table, challenge.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, emailchallenge.ChallengeTable, emailchallenge.ChallengeColumn),
		)
		fromU = sqlgraph.SetNeighbors(ecq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EmailChallenge entity from the query.
// Returns a *NotFoundError when no EmailChallenge was found.
func (ecq *EmailChallengeQuery) First(ctx context.Context) (*EmailChallenge, error) {
	nodes, err := ecq.Limit(1).All(setContextOp(ctx, ecq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailchallenge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ecq *EmailChallengeQuery) FirstX(ctx context.Context) *EmailChallenge {
	node, err := ecq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailChallenge ID from the query.
// Returns a *NotFoundError when no EmailChallenge ID was found.
func (ecq *EmailChallengeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ecq.Limit(1).IDs(setContextOp(ctx, ecq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailchallenge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ecq *EmailChallengeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := ecq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailChallenge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailChallenge entity is found.
// Returns a *NotFoundError when no EmailChallenge entities are found.
func (ecq *EmailChallengeQuery) Only(ctx context.Context) (*EmailChallenge, error) {
	nodes, err := ecq.Limit(2).All(setContextOp(ctx, ecq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailchallenge.Label}
	default:
		return nil, &NotSingularError{emailchallenge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ecq *EmailChallengeQuery) OnlyX(ctx context.Context) *EmailChallenge {
	node, err := ecq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailChallenge ID in the query.
// Returns a *NotSingularError when more than one EmailChallenge ID is found.
// Returns a *NotFoundError when no entities are found.
func (ecq *EmailChallengeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = ecq.Limit(2).IDs(setContextOp(ctx, ecq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailchallenge.Label}
	default:
		err = &NotSingularError{emailchallenge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ecq *EmailChallengeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := ecq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailChallenges.
func (ecq *EmailChallengeQuery) All(ctx context.Context) ([]*EmailChallenge, error) {
	ctx = setContextOp(ctx, ecq.ctx, "All")
	if err := ecq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmailChallenge, *EmailChallengeQuery]()
	return withInterceptors[[]*EmailChallenge](ctx, ecq, qr, ecq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ecq *EmailChallengeQuery) AllX(ctx context.Context) []*EmailChallenge {
	nodes, err := ecq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailChallenge IDs.
func (ecq *EmailChallengeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if ecq.ctx.Unique == nil && ecq.path != nil {
		ecq.Unique(true)
	}
	ctx = setContextOp(ctx, ecq.ctx, "IDs")
	if err = ecq.Select(emailchallenge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ecq *EmailChallengeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := ecq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ecq *EmailChallengeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Count")
	if err := ecq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ecq, querierCount[*EmailChallengeQuery](), ecq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ecq *EmailChallengeQuery) CountX(ctx context.Context) int {
	count, err := ecq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ecq *EmailChallengeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ecq.ctx, "Exist")
	switch _, err := ecq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ecq *EmailChallengeQuery) ExistX(ctx context.Context) bool {
	exist, err := ecq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailChallengeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ecq *EmailChallengeQuery) Clone() *EmailChallengeQuery {
	if ecq == nil {
		return nil
	}
	return &EmailChallengeQuery{
		config:        ecq.config,
		ctx:           ecq.ctx.Clone(),
		order:         append([]OrderFunc{}, ecq.order...),
		inters:        append([]Interceptor{}, ecq.inters...),
		predicates:    append([]predicate.EmailChallenge{}, ecq.predicates...),
		withChallenge: ecq.withChallenge.Clone(),
		// clone intermediate query.
		sql:  ecq.sql.Clone(),
		path: ecq.path,
	}
}

// WithChallenge tells the query-builder to eager-load the nodes that are connected to
// the "challenge" edge. The optional arguments are used to configure the query builder of the edge.
func (ecq *EmailChallengeQuery) WithChallenge(opts ...func(*ChallengeQuery)) *EmailChallengeQuery {
	query := (&ChallengeClient{config: ecq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ecq.withChallenge = query
	return ecq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Code []byte `json:"code,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailChallenge.Query().
//		GroupBy(emailchallenge.FieldCode).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ecq *EmailChallengeQuery) GroupBy(field string, fields ...string) *EmailChallengeGroupBy {
	ecq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmailChallengeGroupBy{build: ecq}
	grbuild.flds = &ecq.ctx.Fields
	grbuild.label = emailchallenge.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Code []byte `json:"code,omitempty"`
//	}
//
//	client.EmailChallenge.Query().
//		Select(emailchallenge.FieldCode).
//		Scan(ctx, &v)
func (ecq *EmailChallengeQuery) Select(fields ...string) *EmailChallengeSelect {
	ecq.ctx.Fields = append(ecq.ctx.Fields, fields...)
	sbuild := &EmailChallengeSelect{EmailChallengeQuery: ecq}
	sbuild.label = emailchallenge.Label
	sbuild.flds, sbuild.scan = &ecq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmailChallengeSelect configured with the given aggregations.
func (ecq *EmailChallengeQuery) Aggregate(fns ...AggregateFunc) *EmailChallengeSelect {
	return ecq.Select().Aggregate(fns...)
}

func (ecq *EmailChallengeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ecq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ecq); err != nil {
				return err
			}
		}
	}
	for _, f := range ecq.ctx.Fields {
		if !emailchallenge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ecq.path != nil {
		prev, err := ecq.path(ctx)
		if err != nil {
			return err
		}
		ecq.sql = prev
	}
	return nil
}

func (ecq *EmailChallengeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailChallenge, error) {
	var (
		nodes       = []*EmailChallenge{}
		withFKs     = ecq.withFKs
		_spec       = ecq.querySpec()
		loadedTypes = [1]bool{
			ecq.withChallenge != nil,
		}
	)
	if ecq.withChallenge != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, emailchallenge.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailChallenge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailChallenge{config: ecq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ecq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ecq.withChallenge; query != nil {
		if err := ecq.loadChallenge(ctx, query, nodes, nil,
			func(n *EmailChallenge, e *Challenge) { n.Edges.Challenge = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ecq *EmailChallengeQuery) loadChallenge(ctx context.Context, query *ChallengeQuery, nodes []*EmailChallenge, init func(*EmailChallenge), assign func(*EmailChallenge, *Challenge)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*EmailChallenge)
	for i := range nodes {
		if nodes[i].challenge_email_challenge == nil {
			continue
		}
		fk := *nodes[i].challenge_email_challenge
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(challenge.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "challenge_email_challenge" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ecq *EmailChallengeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ecq.querySpec()
	_spec.Node.Columns = ecq.ctx.Fields
	if len(ecq.ctx.Fields) > 0 {
		_spec.Unique = ecq.ctx.Unique != nil && *ecq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ecq.driver, _spec)
}

func (ecq *EmailChallengeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(emailchallenge.Table, emailchallenge.Columns, sqlgraph.NewFieldSpec(emailchallenge.FieldID, field.TypeUUID))
	_spec.From = ecq.sql
	if unique := ecq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ecq.path != nil {
		_spec.Unique = true
	}
	if fields := ecq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailchallenge.FieldID)
		for i := range fields {
			if fields[i] != emailchallenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ecq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ecq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ecq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ecq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ecq *EmailChallengeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ecq.driver.Dialect())
	t1 := builder.Table(emailchallenge.Table)
	columns := ecq.ctx.Fields
	if len(columns) == 0 {
		columns = emailchallenge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ecq.sql != nil {
		selector = ecq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ecq.ctx.Unique != nil && *ecq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ecq.predicates {
		p(selector)
	}
	for _, p := range ecq.order {
		p(selector)
	}
	if offset := ecq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ecq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmailChallengeGroupBy is the group-by builder for EmailChallenge entities.
type EmailChallengeGroupBy struct {
	selector
	build *EmailChallengeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ecgb *EmailChallengeGroupBy) Aggregate(fns ...AggregateFunc) *EmailChallengeGroupBy {
	ecgb.fns = append(ecgb.fns, fns...)
	return ecgb
}

// Scan applies the selector query and scans the result into the given value.
func (ecgb *EmailChallengeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecgb.build.ctx, "GroupBy")
	if err := ecgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailChallengeQuery, *EmailChallengeGroupBy](ctx, ecgb.build, ecgb, ecgb.build.inters, v)
}

func (ecgb *EmailChallengeGroupBy) sqlScan(ctx context.Context, root *EmailChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ecgb.fns))
	for _, fn := range ecgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ecgb.flds)+len(ecgb.fns))
		for _, f := range *ecgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ecgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmailChallengeSelect is the builder for selecting fields of EmailChallenge entities.
type EmailChallengeSelect struct {
	*EmailChallengeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ecs *EmailChallengeSelect) Aggregate(fns ...AggregateFunc) *EmailChallengeSelect {
	ecs.fns = append(ecs.fns, fns...)
	return ecs
}

// Scan applies the selector query and scans the result into the given value.
func (ecs *EmailChallengeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ecs.ctx, "Select")
	if err := ecs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailChallengeQuery, *EmailChallengeSelect](ctx, ecs.EmailChallengeQuery, ecs, ecs.inters, v)
}

func (ecs *EmailChallengeSelect) sqlScan(ctx context.Context, root *EmailChallengeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ecs.fns))
	for _, fn := range ecs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ecs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ecs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
