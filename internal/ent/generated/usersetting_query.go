// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/file"
	"github.com/theopenlane/core/internal/ent/generated/organization"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/internal/ent/generated/user"
	"github.com/theopenlane/core/internal/ent/generated/usersetting"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// UserSettingQuery is the builder for querying UserSetting entities.
type UserSettingQuery struct {
	config
	ctx            *QueryContext
	order          []usersetting.OrderOption
	inters         []Interceptor
	predicates     []predicate.UserSetting
	withUser       *UserQuery
	withDefaultOrg *OrganizationQuery
	withFiles      *FileQuery
	withFKs        bool
	loadTotal      []func(context.Context, []*UserSetting) error
	modifiers      []func(*sql.Selector)
	withNamedFiles map[string]*FileQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserSettingQuery builder.
func (usq *UserSettingQuery) Where(ps ...predicate.UserSetting) *UserSettingQuery {
	usq.predicates = append(usq.predicates, ps...)
	return usq
}

// Limit the number of records to be returned by this query.
func (usq *UserSettingQuery) Limit(limit int) *UserSettingQuery {
	usq.ctx.Limit = &limit
	return usq
}

// Offset to start from.
func (usq *UserSettingQuery) Offset(offset int) *UserSettingQuery {
	usq.ctx.Offset = &offset
	return usq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (usq *UserSettingQuery) Unique(unique bool) *UserSettingQuery {
	usq.ctx.Unique = &unique
	return usq
}

// Order specifies how the records should be ordered.
func (usq *UserSettingQuery) Order(o ...usersetting.OrderOption) *UserSettingQuery {
	usq.order = append(usq.order, o...)
	return usq
}

// QueryUser chains the current query on the "user" edge.
func (usq *UserSettingQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: usq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := usq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersetting.Table, usersetting.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, usersetting.UserTable, usersetting.UserColumn),
		)
		schemaConfig := usq.schemaConfig
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserSetting
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDefaultOrg chains the current query on the "default_org" edge.
func (usq *UserSettingQuery) QueryDefaultOrg() *OrganizationQuery {
	query := (&OrganizationClient{config: usq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := usq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersetting.Table, usersetting.FieldID, selector),
			sqlgraph.To(organization.Table, organization.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, usersetting.DefaultOrgTable, usersetting.DefaultOrgColumn),
		)
		schemaConfig := usq.schemaConfig
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.UserSetting
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFiles chains the current query on the "files" edge.
func (usq *UserSettingQuery) QueryFiles() *FileQuery {
	query := (&FileClient{config: usq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := usq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := usq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usersetting.Table, usersetting.FieldID, selector),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, usersetting.FilesTable, usersetting.FilesPrimaryKey...),
		)
		schemaConfig := usq.schemaConfig
		step.To.Schema = schemaConfig.File
		step.Edge.Schema = schemaConfig.UserSettingFiles
		fromU = sqlgraph.SetNeighbors(usq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserSetting entity from the query.
// Returns a *NotFoundError when no UserSetting was found.
func (usq *UserSettingQuery) First(ctx context.Context) (*UserSetting, error) {
	nodes, err := usq.Limit(1).All(setContextOp(ctx, usq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usersetting.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (usq *UserSettingQuery) FirstX(ctx context.Context) *UserSetting {
	node, err := usq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserSetting ID from the query.
// Returns a *NotFoundError when no UserSetting ID was found.
func (usq *UserSettingQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = usq.Limit(1).IDs(setContextOp(ctx, usq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usersetting.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (usq *UserSettingQuery) FirstIDX(ctx context.Context) string {
	id, err := usq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserSetting entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserSetting entity is found.
// Returns a *NotFoundError when no UserSetting entities are found.
func (usq *UserSettingQuery) Only(ctx context.Context) (*UserSetting, error) {
	nodes, err := usq.Limit(2).All(setContextOp(ctx, usq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usersetting.Label}
	default:
		return nil, &NotSingularError{usersetting.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (usq *UserSettingQuery) OnlyX(ctx context.Context) *UserSetting {
	node, err := usq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserSetting ID in the query.
// Returns a *NotSingularError when more than one UserSetting ID is found.
// Returns a *NotFoundError when no entities are found.
func (usq *UserSettingQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = usq.Limit(2).IDs(setContextOp(ctx, usq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usersetting.Label}
	default:
		err = &NotSingularError{usersetting.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (usq *UserSettingQuery) OnlyIDX(ctx context.Context) string {
	id, err := usq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserSettings.
func (usq *UserSettingQuery) All(ctx context.Context) ([]*UserSetting, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryAll)
	if err := usq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserSetting, *UserSettingQuery]()
	return withInterceptors[[]*UserSetting](ctx, usq, qr, usq.inters)
}

// AllX is like All, but panics if an error occurs.
func (usq *UserSettingQuery) AllX(ctx context.Context) []*UserSetting {
	nodes, err := usq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserSetting IDs.
func (usq *UserSettingQuery) IDs(ctx context.Context) (ids []string, err error) {
	if usq.ctx.Unique == nil && usq.path != nil {
		usq.Unique(true)
	}
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryIDs)
	if err = usq.Select(usersetting.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (usq *UserSettingQuery) IDsX(ctx context.Context) []string {
	ids, err := usq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (usq *UserSettingQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryCount)
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, usq, querierCount[*UserSettingQuery](), usq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (usq *UserSettingQuery) CountX(ctx context.Context) int {
	count, err := usq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (usq *UserSettingQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryExist)
	switch _, err := usq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (usq *UserSettingQuery) ExistX(ctx context.Context) bool {
	exist, err := usq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserSettingQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (usq *UserSettingQuery) Clone() *UserSettingQuery {
	if usq == nil {
		return nil
	}
	return &UserSettingQuery{
		config:         usq.config,
		ctx:            usq.ctx.Clone(),
		order:          append([]usersetting.OrderOption{}, usq.order...),
		inters:         append([]Interceptor{}, usq.inters...),
		predicates:     append([]predicate.UserSetting{}, usq.predicates...),
		withUser:       usq.withUser.Clone(),
		withDefaultOrg: usq.withDefaultOrg.Clone(),
		withFiles:      usq.withFiles.Clone(),
		// clone intermediate query.
		sql:       usq.sql.Clone(),
		path:      usq.path,
		modifiers: append([]func(*sql.Selector){}, usq.modifiers...),
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (usq *UserSettingQuery) WithUser(opts ...func(*UserQuery)) *UserSettingQuery {
	query := (&UserClient{config: usq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	usq.withUser = query
	return usq
}

// WithDefaultOrg tells the query-builder to eager-load the nodes that are connected to
// the "default_org" edge. The optional arguments are used to configure the query builder of the edge.
func (usq *UserSettingQuery) WithDefaultOrg(opts ...func(*OrganizationQuery)) *UserSettingQuery {
	query := (&OrganizationClient{config: usq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	usq.withDefaultOrg = query
	return usq
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (usq *UserSettingQuery) WithFiles(opts ...func(*FileQuery)) *UserSettingQuery {
	query := (&FileClient{config: usq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	usq.withFiles = query
	return usq
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
//	client.UserSetting.Query().
//		GroupBy(usersetting.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (usq *UserSettingQuery) GroupBy(field string, fields ...string) *UserSettingGroupBy {
	usq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserSettingGroupBy{build: usq}
	grbuild.flds = &usq.ctx.Fields
	grbuild.label = usersetting.Label
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
//	client.UserSetting.Query().
//		Select(usersetting.FieldCreatedAt).
//		Scan(ctx, &v)
func (usq *UserSettingQuery) Select(fields ...string) *UserSettingSelect {
	usq.ctx.Fields = append(usq.ctx.Fields, fields...)
	sbuild := &UserSettingSelect{UserSettingQuery: usq}
	sbuild.label = usersetting.Label
	sbuild.flds, sbuild.scan = &usq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserSettingSelect configured with the given aggregations.
func (usq *UserSettingQuery) Aggregate(fns ...AggregateFunc) *UserSettingSelect {
	return usq.Select().Aggregate(fns...)
}

func (usq *UserSettingQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range usq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, usq); err != nil {
				return err
			}
		}
	}
	for _, f := range usq.ctx.Fields {
		if !usersetting.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if usq.path != nil {
		prev, err := usq.path(ctx)
		if err != nil {
			return err
		}
		usq.sql = prev
	}
	if usersetting.Policy == nil {
		return errors.New("generated: uninitialized usersetting.Policy (forgotten import generated/runtime?)")
	}
	if err := usersetting.Policy.EvalQuery(ctx, usq); err != nil {
		return err
	}
	return nil
}

func (usq *UserSettingQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserSetting, error) {
	var (
		nodes       = []*UserSetting{}
		withFKs     = usq.withFKs
		_spec       = usq.querySpec()
		loadedTypes = [3]bool{
			usq.withUser != nil,
			usq.withDefaultOrg != nil,
			usq.withFiles != nil,
		}
	)
	if usq.withDefaultOrg != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, usersetting.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserSetting).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserSetting{config: usq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = usq.schemaConfig.UserSetting
	ctx = internal.NewSchemaConfigContext(ctx, usq.schemaConfig)
	if len(usq.modifiers) > 0 {
		_spec.Modifiers = usq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, usq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := usq.withUser; query != nil {
		if err := usq.loadUser(ctx, query, nodes, nil,
			func(n *UserSetting, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := usq.withDefaultOrg; query != nil {
		if err := usq.loadDefaultOrg(ctx, query, nodes, nil,
			func(n *UserSetting, e *Organization) { n.Edges.DefaultOrg = e }); err != nil {
			return nil, err
		}
	}
	if query := usq.withFiles; query != nil {
		if err := usq.loadFiles(ctx, query, nodes,
			func(n *UserSetting) { n.Edges.Files = []*File{} },
			func(n *UserSetting, e *File) { n.Edges.Files = append(n.Edges.Files, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range usq.withNamedFiles {
		if err := usq.loadFiles(ctx, query, nodes,
			func(n *UserSetting) { n.appendNamedFiles(name) },
			func(n *UserSetting, e *File) { n.appendNamedFiles(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range usq.loadTotal {
		if err := usq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (usq *UserSettingQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserSetting, init func(*UserSetting), assign func(*UserSetting, *User)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserSetting)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (usq *UserSettingQuery) loadDefaultOrg(ctx context.Context, query *OrganizationQuery, nodes []*UserSetting, init func(*UserSetting), assign func(*UserSetting, *Organization)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*UserSetting)
	for i := range nodes {
		if nodes[i].user_setting_default_org == nil {
			continue
		}
		fk := *nodes[i].user_setting_default_org
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(organization.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_setting_default_org" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (usq *UserSettingQuery) loadFiles(ctx context.Context, query *FileQuery, nodes []*UserSetting, init func(*UserSetting), assign func(*UserSetting, *File)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*UserSetting)
	nids := make(map[string]map[*UserSetting]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(usersetting.FilesTable)
		joinT.Schema(usq.schemaConfig.UserSettingFiles)
		s.Join(joinT).On(s.C(file.FieldID), joinT.C(usersetting.FilesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(usersetting.FilesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(usersetting.FilesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*UserSetting]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*File](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "files" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (usq *UserSettingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := usq.querySpec()
	_spec.Node.Schema = usq.schemaConfig.UserSetting
	ctx = internal.NewSchemaConfigContext(ctx, usq.schemaConfig)
	if len(usq.modifiers) > 0 {
		_spec.Modifiers = usq.modifiers
	}
	_spec.Node.Columns = usq.ctx.Fields
	if len(usq.ctx.Fields) > 0 {
		_spec.Unique = usq.ctx.Unique != nil && *usq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, usq.driver, _spec)
}

func (usq *UserSettingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usersetting.Table, usersetting.Columns, sqlgraph.NewFieldSpec(usersetting.FieldID, field.TypeString))
	_spec.From = usq.sql
	if unique := usq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if usq.path != nil {
		_spec.Unique = true
	}
	if fields := usq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usersetting.FieldID)
		for i := range fields {
			if fields[i] != usersetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if usq.withUser != nil {
			_spec.Node.AddColumnOnce(usersetting.FieldUserID)
		}
	}
	if ps := usq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := usq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := usq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := usq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (usq *UserSettingQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(usq.driver.Dialect())
	t1 := builder.Table(usersetting.Table)
	columns := usq.ctx.Fields
	if len(columns) == 0 {
		columns = usersetting.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if usq.sql != nil {
		selector = usq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if usq.ctx.Unique != nil && *usq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(usq.schemaConfig.UserSetting)
	ctx = internal.NewSchemaConfigContext(ctx, usq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range usq.modifiers {
		m(selector)
	}
	for _, p := range usq.predicates {
		p(selector)
	}
	for _, p := range usq.order {
		p(selector)
	}
	if offset := usq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := usq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (usq *UserSettingQuery) Modify(modifiers ...func(s *sql.Selector)) *UserSettingSelect {
	usq.modifiers = append(usq.modifiers, modifiers...)
	return usq.Select()
}

// WithNamedFiles tells the query-builder to eager-load the nodes that are connected to the "files"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (usq *UserSettingQuery) WithNamedFiles(name string, opts ...func(*FileQuery)) *UserSettingQuery {
	query := (&FileClient{config: usq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if usq.withNamedFiles == nil {
		usq.withNamedFiles = make(map[string]*FileQuery)
	}
	usq.withNamedFiles[name] = query
	return usq
}

// CountIDs returns the count of ids and allows for filtering of the query post retrieval by IDs
func (usq *UserSettingQuery) CountIDs(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, usq.ctx, ent.OpQueryIDs)
	if err := usq.prepareQuery(ctx); err != nil {
		return 0, err
	}

	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return usq.IDs(ctx)
	})

	ids, err := withInterceptors[[]string](ctx, usq, qr, usq.inters)
	if err != nil {
		return 0, err
	}

	return len(ids), nil
}

// UserSettingGroupBy is the group-by builder for UserSetting entities.
type UserSettingGroupBy struct {
	selector
	build *UserSettingQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (usgb *UserSettingGroupBy) Aggregate(fns ...AggregateFunc) *UserSettingGroupBy {
	usgb.fns = append(usgb.fns, fns...)
	return usgb
}

// Scan applies the selector query and scans the result into the given value.
func (usgb *UserSettingGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, usgb.build.ctx, ent.OpQueryGroupBy)
	if err := usgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSettingQuery, *UserSettingGroupBy](ctx, usgb.build, usgb, usgb.build.inters, v)
}

func (usgb *UserSettingGroupBy) sqlScan(ctx context.Context, root *UserSettingQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(usgb.fns))
	for _, fn := range usgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*usgb.flds)+len(usgb.fns))
		for _, f := range *usgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*usgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := usgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserSettingSelect is the builder for selecting fields of UserSetting entities.
type UserSettingSelect struct {
	*UserSettingQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uss *UserSettingSelect) Aggregate(fns ...AggregateFunc) *UserSettingSelect {
	uss.fns = append(uss.fns, fns...)
	return uss
}

// Scan applies the selector query and scans the result into the given value.
func (uss *UserSettingSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uss.ctx, ent.OpQuerySelect)
	if err := uss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserSettingQuery, *UserSettingSelect](ctx, uss.UserSettingQuery, uss, uss.inters, v)
}

func (uss *UserSettingSelect) sqlScan(ctx context.Context, root *UserSettingQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uss.fns))
	for _, fn := range uss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (uss *UserSettingSelect) Modify(modifiers ...func(s *sql.Selector)) *UserSettingSelect {
	uss.modifiers = append(uss.modifiers, modifiers...)
	return uss
}
