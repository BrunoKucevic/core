// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/internal/ent/generated/tfasetting"
	"github.com/theopenlane/core/internal/ent/generated/user"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// TFASettingQuery is the builder for querying TFASetting entities.
type TFASettingQuery struct {
	config
	ctx        *QueryContext
	order      []tfasetting.OrderOption
	inters     []Interceptor
	predicates []predicate.TFASetting
	withOwner  *UserQuery
	loadTotal  []func(context.Context, []*TFASetting) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TFASettingQuery builder.
func (tsq *TFASettingQuery) Where(ps ...predicate.TFASetting) *TFASettingQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TFASettingQuery) Limit(limit int) *TFASettingQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TFASettingQuery) Offset(offset int) *TFASettingQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TFASettingQuery) Unique(unique bool) *TFASettingQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TFASettingQuery) Order(o ...tfasetting.OrderOption) *TFASettingQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryOwner chains the current query on the "owner" edge.
func (tsq *TFASettingQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tfasetting.Table, tfasetting.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tfasetting.OwnerTable, tfasetting.OwnerColumn),
		)
		schemaConfig := tsq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.TFASetting
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TFASetting entity from the query.
// Returns a *NotFoundError when no TFASetting was found.
func (tsq *TFASettingQuery) First(ctx context.Context) (*TFASetting, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tfasetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TFASettingQuery) FirstX(ctx context.Context) *TFASetting {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TFASetting ID from the query.
// Returns a *NotFoundError when no TFASetting ID was found.
func (tsq *TFASettingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tfasetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TFASettingQuery) FirstIDX(ctx context.Context) string {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TFASetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TFASetting entity is found.
// Returns a *NotFoundError when no TFASetting entities are found.
func (tsq *TFASettingQuery) Only(ctx context.Context) (*TFASetting, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tfasetting.Label}
	default:
		return nil, &NotSingularError{tfasetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TFASettingQuery) OnlyX(ctx context.Context) *TFASetting {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TFASetting ID in the query.
// Returns a *NotSingularError when more than one TFASetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TFASettingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tfasetting.Label}
	default:
		err = &NotSingularError{tfasetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TFASettingQuery) OnlyIDX(ctx context.Context) string {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TFASettings.
func (tsq *TFASettingQuery) All(ctx context.Context) ([]*TFASetting, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryAll)
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TFASetting, *TFASettingQuery]()
	return withInterceptors[[]*TFASetting](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TFASettingQuery) AllX(ctx context.Context) []*TFASetting {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TFASetting IDs.
func (tsq *TFASettingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryIDs)
	if err = tsq.Select(tfasetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TFASettingQuery) IDsX(ctx context.Context) []string {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TFASettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryCount)
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TFASettingQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TFASettingQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TFASettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryExist)
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TFASettingQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TFASettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TFASettingQuery) Clone() *TFASettingQuery {
	if tsq == nil {
		return nil
	}
	return &TFASettingQuery{
		config:     tsq.config,
		ctx:        tsq.ctx.Clone(),
		order:      append([]tfasetting.OrderOption{}, tsq.order...),
		inters:     append([]Interceptor{}, tsq.inters...),
		predicates: append([]predicate.TFASetting{}, tsq.predicates...),
		withOwner:  tsq.withOwner.Clone(),
		// clone intermediate query.
		sql:       tsq.sql.Clone(),
		path:      tsq.path,
		modifiers: append([]func(*sql.Selector){}, tsq.modifiers...),
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TFASettingQuery) WithOwner(opts ...func(*UserQuery)) *TFASettingQuery {
	query := (&UserClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withOwner = query
	return tsq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TFASetting.Query().
//		GroupBy(tfasetting.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (tsq *TFASettingQuery) GroupBy(field string, fields ...string) *TFASettingGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TFASettingGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = tfasetting.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.TFASetting.Query().
//		Select(tfasetting.FieldCreatedAt).
//		Scan(ctx, &v)
func (tsq *TFASettingQuery) Select(fields ...string) *TFASettingSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TFASettingSelect{TFASettingQuery: tsq}
	sbuild.label = tfasetting.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TFASettingSelect configured with the given aggregations.
func (tsq *TFASettingQuery) Aggregate(fns ...AggregateFunc) *TFASettingSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TFASettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !tfasetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	if tfasetting.Policy == nil {
		return errors.New("generated: uninitialized tfasetting.Policy (forgotten import generated/runtime?)")
	}
	if err := tfasetting.Policy.EvalQuery(ctx, tsq); err != nil {
		return err
	}
	return nil
}

func (tsq *TFASettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TFASetting, error) {
	var (
		nodes       = []*TFASetting{}
		_spec       = tsq.querySpec()
		loadedTypes = [1]bool{
			tsq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TFASetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TFASetting{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tsq.schemaConfig.TFASetting
	ctx = internal.NewSchemaConfigContext(ctx, tsq.schemaConfig)
	if len(tsq.modifiers) > 0 {
		_spec.Modifiers = tsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tsq.withOwner; query != nil {
		if err := tsq.loadOwner(ctx, query, nodes, nil,
			func(n *TFASetting, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	for i := range tsq.loadTotal {
		if err := tsq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tsq *TFASettingQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*TFASetting, init func(*TFASetting), assign func(*TFASetting, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TFASetting)
	for i := range nodes {
		fk := nodes[i].OwnerID
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
			return fmt.Errorf(`unexpected foreign-key "owner_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tsq *TFASettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Schema = tsq.schemaConfig.TFASetting
	ctx = internal.NewSchemaConfigContext(ctx, tsq.schemaConfig)
	if len(tsq.modifiers) > 0 {
		_spec.Modifiers = tsq.modifiers
	}
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TFASettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tfasetting.Table, tfasetting.Columns, sqlgraph.NewFieldSpec(tfasetting.FieldID, field.TypeString))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tfasetting.FieldID)
		for i := range fields {
			if fields[i] != tfasetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tsq.withOwner != nil {
			_spec.Node.AddColumnOnce(tfasetting.FieldOwnerID)
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TFASettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(tfasetting.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = tfasetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tsq.schemaConfig.TFASetting)
	ctx = internal.NewSchemaConfigContext(ctx, tsq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tsq.modifiers {
		m(selector)
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tsq *TFASettingQuery) Modify(modifiers ...func(s *sql.Selector)) *TFASettingSelect {
	tsq.modifiers = append(tsq.modifiers, modifiers...)
	return tsq.Select()
}

// CountIDs returns the count of ids and allows for filtering of the query post retrieval by IDs
func (tsq *TFASettingQuery) CountIDs(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, ent.OpQueryIDs)
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}

	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return tsq.IDs(ctx)
	})

	ids, err := withInterceptors[[]string](ctx, tsq, qr, tsq.inters)
	if err != nil {
		return 0, err
	}

	return len(ids), nil
}

// TFASettingGroupBy is the group-by builder for TFASetting entities.
type TFASettingGroupBy struct {
	selector
	build *TFASettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TFASettingGroupBy) Aggregate(fns ...AggregateFunc) *TFASettingGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TFASettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, ent.OpQueryGroupBy)
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TFASettingQuery, *TFASettingGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TFASettingGroupBy) sqlScan(ctx context.Context, root *TFASettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TFASettingSelect is the builder for selecting fields of TFASetting entities.
type TFASettingSelect struct {
	*TFASettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TFASettingSelect) Aggregate(fns ...AggregateFunc) *TFASettingSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TFASettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, ent.OpQuerySelect)
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TFASettingQuery, *TFASettingSelect](ctx, tss.TFASettingQuery, tss, tss.inters, v)
}

func (tss *TFASettingSelect) sqlScan(ctx context.Context, root *TFASettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tss *TFASettingSelect) Modify(modifiers ...func(s *sql.Selector)) *TFASettingSelect {
	tss.modifiers = append(tss.modifiers, modifiers...)
	return tss
}
