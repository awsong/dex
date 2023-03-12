// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/awsong/dex/storage/ent/db/keys"
	"github.com/awsong/dex/storage/ent/db/predicate"
)

// KeysQuery is the builder for querying Keys entities.
type KeysQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Keys
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeysQuery builder.
func (kq *KeysQuery) Where(ps ...predicate.Keys) *KeysQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit the number of records to be returned by this query.
func (kq *KeysQuery) Limit(limit int) *KeysQuery {
	kq.ctx.Limit = &limit
	return kq
}

// Offset to start from.
func (kq *KeysQuery) Offset(offset int) *KeysQuery {
	kq.ctx.Offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KeysQuery) Unique(unique bool) *KeysQuery {
	kq.ctx.Unique = &unique
	return kq
}

// Order specifies how the records should be ordered.
func (kq *KeysQuery) Order(o ...OrderFunc) *KeysQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// First returns the first Keys entity from the query.
// Returns a *NotFoundError when no Keys was found.
func (kq *KeysQuery) First(ctx context.Context) (*Keys, error) {
	nodes, err := kq.Limit(1).All(setContextOp(ctx, kq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keys.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KeysQuery) FirstX(ctx context.Context) *Keys {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Keys ID from the query.
// Returns a *NotFoundError when no Keys ID was found.
func (kq *KeysQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(1).IDs(setContextOp(ctx, kq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keys.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KeysQuery) FirstIDX(ctx context.Context) string {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Keys entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Keys entity is found.
// Returns a *NotFoundError when no Keys entities are found.
func (kq *KeysQuery) Only(ctx context.Context) (*Keys, error) {
	nodes, err := kq.Limit(2).All(setContextOp(ctx, kq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keys.Label}
	default:
		return nil, &NotSingularError{keys.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KeysQuery) OnlyX(ctx context.Context) *Keys {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Keys ID in the query.
// Returns a *NotSingularError when more than one Keys ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KeysQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = kq.Limit(2).IDs(setContextOp(ctx, kq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keys.Label}
	default:
		err = &NotSingularError{keys.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KeysQuery) OnlyIDX(ctx context.Context) string {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeysSlice.
func (kq *KeysQuery) All(ctx context.Context) ([]*Keys, error) {
	ctx = setContextOp(ctx, kq.ctx, "All")
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Keys, *KeysQuery]()
	return withInterceptors[[]*Keys](ctx, kq, qr, kq.inters)
}

// AllX is like All, but panics if an error occurs.
func (kq *KeysQuery) AllX(ctx context.Context) []*Keys {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Keys IDs.
func (kq *KeysQuery) IDs(ctx context.Context) (ids []string, err error) {
	if kq.ctx.Unique == nil && kq.path != nil {
		kq.Unique(true)
	}
	ctx = setContextOp(ctx, kq.ctx, "IDs")
	if err = kq.Select(keys.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KeysQuery) IDsX(ctx context.Context) []string {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KeysQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, kq.ctx, "Count")
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, kq, querierCount[*KeysQuery](), kq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KeysQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KeysQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, kq.ctx, "Exist")
	switch _, err := kq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KeysQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeysQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KeysQuery) Clone() *KeysQuery {
	if kq == nil {
		return nil
	}
	return &KeysQuery{
		config:     kq.config,
		ctx:        kq.ctx.Clone(),
		order:      append([]OrderFunc{}, kq.order...),
		inters:     append([]Interceptor{}, kq.inters...),
		predicates: append([]predicate.Keys{}, kq.predicates...),
		// clone intermediate query.
		sql:  kq.sql.Clone(),
		path: kq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VerificationKeys []storage.VerificationKey `json:"verification_keys,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Keys.Query().
//		GroupBy(keys.FieldVerificationKeys).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (kq *KeysQuery) GroupBy(field string, fields ...string) *KeysGroupBy {
	kq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &KeysGroupBy{build: kq}
	grbuild.flds = &kq.ctx.Fields
	grbuild.label = keys.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VerificationKeys []storage.VerificationKey `json:"verification_keys,omitempty"`
//	}
//
//	client.Keys.Query().
//		Select(keys.FieldVerificationKeys).
//		Scan(ctx, &v)
func (kq *KeysQuery) Select(fields ...string) *KeysSelect {
	kq.ctx.Fields = append(kq.ctx.Fields, fields...)
	sbuild := &KeysSelect{KeysQuery: kq}
	sbuild.label = keys.Label
	sbuild.flds, sbuild.scan = &kq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a KeysSelect configured with the given aggregations.
func (kq *KeysQuery) Aggregate(fns ...AggregateFunc) *KeysSelect {
	return kq.Select().Aggregate(fns...)
}

func (kq *KeysQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range kq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, kq); err != nil {
				return err
			}
		}
	}
	for _, f := range kq.ctx.Fields {
		if !keys.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	return nil
}

func (kq *KeysQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Keys, error) {
	var (
		nodes = []*Keys{}
		_spec = kq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Keys).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Keys{config: kq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kq *KeysQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	_spec.Node.Columns = kq.ctx.Fields
	if len(kq.ctx.Fields) > 0 {
		_spec.Unique = kq.ctx.Unique != nil && *kq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KeysQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(keys.Table, keys.Columns, sqlgraph.NewFieldSpec(keys.FieldID, field.TypeString))
	_spec.From = kq.sql
	if unique := kq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if kq.path != nil {
		_spec.Unique = true
	}
	if fields := kq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keys.FieldID)
		for i := range fields {
			if fields[i] != keys.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KeysQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(keys.Table)
	columns := kq.ctx.Fields
	if len(columns) == 0 {
		columns = keys.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.ctx.Unique != nil && *kq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeysGroupBy is the group-by builder for Keys entities.
type KeysGroupBy struct {
	selector
	build *KeysQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KeysGroupBy) Aggregate(fns ...AggregateFunc) *KeysGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the selector query and scans the result into the given value.
func (kgb *KeysGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kgb.build.ctx, "GroupBy")
	if err := kgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeysQuery, *KeysGroupBy](ctx, kgb.build, kgb, kgb.build.inters, v)
}

func (kgb *KeysGroupBy) sqlScan(ctx context.Context, root *KeysQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*kgb.flds)+len(kgb.fns))
		for _, f := range *kgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*kgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// KeysSelect is the builder for selecting fields of Keys entities.
type KeysSelect struct {
	*KeysQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ks *KeysSelect) Aggregate(fns ...AggregateFunc) *KeysSelect {
	ks.fns = append(ks.fns, fns...)
	return ks
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KeysSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ks.ctx, "Select")
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeysQuery, *KeysSelect](ctx, ks.KeysQuery, ks, ks.inters, v)
}

func (ks *KeysSelect) sqlScan(ctx context.Context, root *KeysQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ks.fns))
	for _, fn := range ks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
