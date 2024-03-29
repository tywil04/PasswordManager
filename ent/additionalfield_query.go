// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PasswordManager/ent/additionalfield"
	"PasswordManager/ent/password"
	"PasswordManager/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AdditionalFieldQuery is the builder for querying AdditionalField entities.
type AdditionalFieldQuery struct {
	config
	ctx          *QueryContext
	order        []OrderFunc
	inters       []Interceptor
	predicates   []predicate.AdditionalField
	withPassword *PasswordQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AdditionalFieldQuery builder.
func (afq *AdditionalFieldQuery) Where(ps ...predicate.AdditionalField) *AdditionalFieldQuery {
	afq.predicates = append(afq.predicates, ps...)
	return afq
}

// Limit the number of records to be returned by this query.
func (afq *AdditionalFieldQuery) Limit(limit int) *AdditionalFieldQuery {
	afq.ctx.Limit = &limit
	return afq
}

// Offset to start from.
func (afq *AdditionalFieldQuery) Offset(offset int) *AdditionalFieldQuery {
	afq.ctx.Offset = &offset
	return afq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (afq *AdditionalFieldQuery) Unique(unique bool) *AdditionalFieldQuery {
	afq.ctx.Unique = &unique
	return afq
}

// Order specifies how the records should be ordered.
func (afq *AdditionalFieldQuery) Order(o ...OrderFunc) *AdditionalFieldQuery {
	afq.order = append(afq.order, o...)
	return afq
}

// QueryPassword chains the current query on the "password" edge.
func (afq *AdditionalFieldQuery) QueryPassword() *PasswordQuery {
	query := (&PasswordClient{config: afq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := afq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := afq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(additionalfield.Table, additionalfield.FieldID, selector),
			sqlgraph.To(password.Table, password.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, additionalfield.PasswordTable, additionalfield.PasswordColumn),
		)
		fromU = sqlgraph.SetNeighbors(afq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AdditionalField entity from the query.
// Returns a *NotFoundError when no AdditionalField was found.
func (afq *AdditionalFieldQuery) First(ctx context.Context) (*AdditionalField, error) {
	nodes, err := afq.Limit(1).All(setContextOp(ctx, afq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{additionalfield.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (afq *AdditionalFieldQuery) FirstX(ctx context.Context) *AdditionalField {
	node, err := afq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AdditionalField ID from the query.
// Returns a *NotFoundError when no AdditionalField ID was found.
func (afq *AdditionalFieldQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = afq.Limit(1).IDs(setContextOp(ctx, afq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{additionalfield.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (afq *AdditionalFieldQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := afq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AdditionalField entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AdditionalField entity is found.
// Returns a *NotFoundError when no AdditionalField entities are found.
func (afq *AdditionalFieldQuery) Only(ctx context.Context) (*AdditionalField, error) {
	nodes, err := afq.Limit(2).All(setContextOp(ctx, afq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{additionalfield.Label}
	default:
		return nil, &NotSingularError{additionalfield.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (afq *AdditionalFieldQuery) OnlyX(ctx context.Context) *AdditionalField {
	node, err := afq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AdditionalField ID in the query.
// Returns a *NotSingularError when more than one AdditionalField ID is found.
// Returns a *NotFoundError when no entities are found.
func (afq *AdditionalFieldQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = afq.Limit(2).IDs(setContextOp(ctx, afq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{additionalfield.Label}
	default:
		err = &NotSingularError{additionalfield.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (afq *AdditionalFieldQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := afq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AdditionalFields.
func (afq *AdditionalFieldQuery) All(ctx context.Context) ([]*AdditionalField, error) {
	ctx = setContextOp(ctx, afq.ctx, "All")
	if err := afq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AdditionalField, *AdditionalFieldQuery]()
	return withInterceptors[[]*AdditionalField](ctx, afq, qr, afq.inters)
}

// AllX is like All, but panics if an error occurs.
func (afq *AdditionalFieldQuery) AllX(ctx context.Context) []*AdditionalField {
	nodes, err := afq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AdditionalField IDs.
func (afq *AdditionalFieldQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if afq.ctx.Unique == nil && afq.path != nil {
		afq.Unique(true)
	}
	ctx = setContextOp(ctx, afq.ctx, "IDs")
	if err = afq.Select(additionalfield.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (afq *AdditionalFieldQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := afq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (afq *AdditionalFieldQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, afq.ctx, "Count")
	if err := afq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, afq, querierCount[*AdditionalFieldQuery](), afq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (afq *AdditionalFieldQuery) CountX(ctx context.Context) int {
	count, err := afq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (afq *AdditionalFieldQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, afq.ctx, "Exist")
	switch _, err := afq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (afq *AdditionalFieldQuery) ExistX(ctx context.Context) bool {
	exist, err := afq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AdditionalFieldQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (afq *AdditionalFieldQuery) Clone() *AdditionalFieldQuery {
	if afq == nil {
		return nil
	}
	return &AdditionalFieldQuery{
		config:       afq.config,
		ctx:          afq.ctx.Clone(),
		order:        append([]OrderFunc{}, afq.order...),
		inters:       append([]Interceptor{}, afq.inters...),
		predicates:   append([]predicate.AdditionalField{}, afq.predicates...),
		withPassword: afq.withPassword.Clone(),
		// clone intermediate query.
		sql:  afq.sql.Clone(),
		path: afq.path,
	}
}

// WithPassword tells the query-builder to eager-load the nodes that are connected to
// the "password" edge. The optional arguments are used to configure the query builder of the edge.
func (afq *AdditionalFieldQuery) WithPassword(opts ...func(*PasswordQuery)) *AdditionalFieldQuery {
	query := (&PasswordClient{config: afq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	afq.withPassword = query
	return afq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key []byte `form:"key" json:"key" xml:"key"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AdditionalField.Query().
//		GroupBy(additionalfield.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (afq *AdditionalFieldQuery) GroupBy(field string, fields ...string) *AdditionalFieldGroupBy {
	afq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AdditionalFieldGroupBy{build: afq}
	grbuild.flds = &afq.ctx.Fields
	grbuild.label = additionalfield.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Key []byte `form:"key" json:"key" xml:"key"`
//	}
//
//	client.AdditionalField.Query().
//		Select(additionalfield.FieldKey).
//		Scan(ctx, &v)
func (afq *AdditionalFieldQuery) Select(fields ...string) *AdditionalFieldSelect {
	afq.ctx.Fields = append(afq.ctx.Fields, fields...)
	sbuild := &AdditionalFieldSelect{AdditionalFieldQuery: afq}
	sbuild.label = additionalfield.Label
	sbuild.flds, sbuild.scan = &afq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AdditionalFieldSelect configured with the given aggregations.
func (afq *AdditionalFieldQuery) Aggregate(fns ...AggregateFunc) *AdditionalFieldSelect {
	return afq.Select().Aggregate(fns...)
}

func (afq *AdditionalFieldQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range afq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, afq); err != nil {
				return err
			}
		}
	}
	for _, f := range afq.ctx.Fields {
		if !additionalfield.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if afq.path != nil {
		prev, err := afq.path(ctx)
		if err != nil {
			return err
		}
		afq.sql = prev
	}
	return nil
}

func (afq *AdditionalFieldQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AdditionalField, error) {
	var (
		nodes       = []*AdditionalField{}
		withFKs     = afq.withFKs
		_spec       = afq.querySpec()
		loadedTypes = [1]bool{
			afq.withPassword != nil,
		}
	)
	if afq.withPassword != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, additionalfield.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AdditionalField).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AdditionalField{config: afq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, afq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := afq.withPassword; query != nil {
		if err := afq.loadPassword(ctx, query, nodes, nil,
			func(n *AdditionalField, e *Password) { n.Edges.Password = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (afq *AdditionalFieldQuery) loadPassword(ctx context.Context, query *PasswordQuery, nodes []*AdditionalField, init func(*AdditionalField), assign func(*AdditionalField, *Password)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*AdditionalField)
	for i := range nodes {
		if nodes[i].password_additional_fields == nil {
			continue
		}
		fk := *nodes[i].password_additional_fields
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(password.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "password_additional_fields" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (afq *AdditionalFieldQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := afq.querySpec()
	_spec.Node.Columns = afq.ctx.Fields
	if len(afq.ctx.Fields) > 0 {
		_spec.Unique = afq.ctx.Unique != nil && *afq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, afq.driver, _spec)
}

func (afq *AdditionalFieldQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(additionalfield.Table, additionalfield.Columns, sqlgraph.NewFieldSpec(additionalfield.FieldID, field.TypeUUID))
	_spec.From = afq.sql
	if unique := afq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if afq.path != nil {
		_spec.Unique = true
	}
	if fields := afq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, additionalfield.FieldID)
		for i := range fields {
			if fields[i] != additionalfield.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := afq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := afq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := afq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := afq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (afq *AdditionalFieldQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(afq.driver.Dialect())
	t1 := builder.Table(additionalfield.Table)
	columns := afq.ctx.Fields
	if len(columns) == 0 {
		columns = additionalfield.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if afq.sql != nil {
		selector = afq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if afq.ctx.Unique != nil && *afq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range afq.predicates {
		p(selector)
	}
	for _, p := range afq.order {
		p(selector)
	}
	if offset := afq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := afq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AdditionalFieldGroupBy is the group-by builder for AdditionalField entities.
type AdditionalFieldGroupBy struct {
	selector
	build *AdditionalFieldQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (afgb *AdditionalFieldGroupBy) Aggregate(fns ...AggregateFunc) *AdditionalFieldGroupBy {
	afgb.fns = append(afgb.fns, fns...)
	return afgb
}

// Scan applies the selector query and scans the result into the given value.
func (afgb *AdditionalFieldGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, afgb.build.ctx, "GroupBy")
	if err := afgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdditionalFieldQuery, *AdditionalFieldGroupBy](ctx, afgb.build, afgb, afgb.build.inters, v)
}

func (afgb *AdditionalFieldGroupBy) sqlScan(ctx context.Context, root *AdditionalFieldQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(afgb.fns))
	for _, fn := range afgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*afgb.flds)+len(afgb.fns))
		for _, f := range *afgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*afgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := afgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AdditionalFieldSelect is the builder for selecting fields of AdditionalField entities.
type AdditionalFieldSelect struct {
	*AdditionalFieldQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (afs *AdditionalFieldSelect) Aggregate(fns ...AggregateFunc) *AdditionalFieldSelect {
	afs.fns = append(afs.fns, fns...)
	return afs
}

// Scan applies the selector query and scans the result into the given value.
func (afs *AdditionalFieldSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, afs.ctx, "Select")
	if err := afs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AdditionalFieldQuery, *AdditionalFieldSelect](ctx, afs.AdditionalFieldQuery, afs, afs.inters, v)
}

func (afs *AdditionalFieldSelect) sqlScan(ctx context.Context, root *AdditionalFieldQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(afs.fns))
	for _, fn := range afs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*afs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := afs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
