// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/enuesaa/pinit/internal/ent/binder"
	"github.com/enuesaa/pinit/internal/ent/predicate"
)

// BinderQuery is the builder for querying Binder entities.
type BinderQuery struct {
	config
	ctx        *QueryContext
	order      []binder.OrderOption
	inters     []Interceptor
	predicates []predicate.Binder
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BinderQuery builder.
func (bq *BinderQuery) Where(ps ...predicate.Binder) *BinderQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BinderQuery) Limit(limit int) *BinderQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BinderQuery) Offset(offset int) *BinderQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BinderQuery) Unique(unique bool) *BinderQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BinderQuery) Order(o ...binder.OrderOption) *BinderQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Binder entity from the query.
// Returns a *NotFoundError when no Binder was found.
func (bq *BinderQuery) First(ctx context.Context) (*Binder, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{binder.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BinderQuery) FirstX(ctx context.Context) *Binder {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Binder ID from the query.
// Returns a *NotFoundError when no Binder ID was found.
func (bq *BinderQuery) FirstID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{binder.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BinderQuery) FirstIDX(ctx context.Context) uint {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Binder entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Binder entity is found.
// Returns a *NotFoundError when no Binder entities are found.
func (bq *BinderQuery) Only(ctx context.Context) (*Binder, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{binder.Label}
	default:
		return nil, &NotSingularError{binder.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BinderQuery) OnlyX(ctx context.Context) *Binder {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Binder ID in the query.
// Returns a *NotSingularError when more than one Binder ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BinderQuery) OnlyID(ctx context.Context) (id uint, err error) {
	var ids []uint
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{binder.Label}
	default:
		err = &NotSingularError{binder.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BinderQuery) OnlyIDX(ctx context.Context) uint {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Binders.
func (bq *BinderQuery) All(ctx context.Context) ([]*Binder, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Binder, *BinderQuery]()
	return withInterceptors[[]*Binder](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BinderQuery) AllX(ctx context.Context) []*Binder {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Binder IDs.
func (bq *BinderQuery) IDs(ctx context.Context) (ids []uint, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryIDs)
	if err = bq.Select(binder.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BinderQuery) IDsX(ctx context.Context) []uint {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BinderQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BinderQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BinderQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BinderQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryExist)
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BinderQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BinderQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BinderQuery) Clone() *BinderQuery {
	if bq == nil {
		return nil
	}
	return &BinderQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]binder.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Binder{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Binder.Query().
//		GroupBy(binder.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BinderQuery) GroupBy(field string, fields ...string) *BinderGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BinderGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = binder.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Binder.Query().
//		Select(binder.FieldName).
//		Scan(ctx, &v)
func (bq *BinderQuery) Select(fields ...string) *BinderSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BinderSelect{BinderQuery: bq}
	sbuild.label = binder.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BinderSelect configured with the given aggregations.
func (bq *BinderQuery) Aggregate(fns ...AggregateFunc) *BinderSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BinderQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !binder.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BinderQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Binder, error) {
	var (
		nodes = []*Binder{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Binder).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Binder{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BinderQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BinderQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(binder.Table, binder.Columns, sqlgraph.NewFieldSpec(binder.FieldID, field.TypeUint))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, binder.FieldID)
		for i := range fields {
			if fields[i] != binder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BinderQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(binder.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = binder.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BinderGroupBy is the group-by builder for Binder entities.
type BinderGroupBy struct {
	selector
	build *BinderQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BinderGroupBy) Aggregate(fns ...AggregateFunc) *BinderGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BinderGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BinderQuery, *BinderGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BinderGroupBy) sqlScan(ctx context.Context, root *BinderQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BinderSelect is the builder for selecting fields of Binder entities.
type BinderSelect struct {
	*BinderQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BinderSelect) Aggregate(fns ...AggregateFunc) *BinderSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BinderSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BinderQuery, *BinderSelect](ctx, bs.BinderQuery, bs, bs.inters, v)
}

func (bs *BinderSelect) sqlScan(ctx context.Context, root *BinderQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
