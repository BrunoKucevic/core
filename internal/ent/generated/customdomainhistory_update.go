// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/customdomainhistory"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/pkg/enums"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// CustomDomainHistoryUpdate is the builder for updating CustomDomainHistory entities.
type CustomDomainHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *CustomDomainHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CustomDomainHistoryUpdate builder.
func (cdhu *CustomDomainHistoryUpdate) Where(ps ...predicate.CustomDomainHistory) *CustomDomainHistoryUpdate {
	cdhu.mutation.Where(ps...)
	return cdhu
}

// SetUpdatedAt sets the "updated_at" field.
func (cdhu *CustomDomainHistoryUpdate) SetUpdatedAt(t time.Time) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetUpdatedAt(t)
	return cdhu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cdhu *CustomDomainHistoryUpdate) ClearUpdatedAt() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearUpdatedAt()
	return cdhu
}

// SetUpdatedBy sets the "updated_by" field.
func (cdhu *CustomDomainHistoryUpdate) SetUpdatedBy(s string) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetUpdatedBy(s)
	return cdhu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cdhu *CustomDomainHistoryUpdate) SetNillableUpdatedBy(s *string) *CustomDomainHistoryUpdate {
	if s != nil {
		cdhu.SetUpdatedBy(*s)
	}
	return cdhu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cdhu *CustomDomainHistoryUpdate) ClearUpdatedBy() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearUpdatedBy()
	return cdhu
}

// SetDeletedAt sets the "deleted_at" field.
func (cdhu *CustomDomainHistoryUpdate) SetDeletedAt(t time.Time) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetDeletedAt(t)
	return cdhu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cdhu *CustomDomainHistoryUpdate) SetNillableDeletedAt(t *time.Time) *CustomDomainHistoryUpdate {
	if t != nil {
		cdhu.SetDeletedAt(*t)
	}
	return cdhu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cdhu *CustomDomainHistoryUpdate) ClearDeletedAt() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearDeletedAt()
	return cdhu
}

// SetDeletedBy sets the "deleted_by" field.
func (cdhu *CustomDomainHistoryUpdate) SetDeletedBy(s string) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetDeletedBy(s)
	return cdhu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cdhu *CustomDomainHistoryUpdate) SetNillableDeletedBy(s *string) *CustomDomainHistoryUpdate {
	if s != nil {
		cdhu.SetDeletedBy(*s)
	}
	return cdhu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (cdhu *CustomDomainHistoryUpdate) ClearDeletedBy() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearDeletedBy()
	return cdhu
}

// SetTags sets the "tags" field.
func (cdhu *CustomDomainHistoryUpdate) SetTags(s []string) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetTags(s)
	return cdhu
}

// AppendTags appends s to the "tags" field.
func (cdhu *CustomDomainHistoryUpdate) AppendTags(s []string) *CustomDomainHistoryUpdate {
	cdhu.mutation.AppendTags(s)
	return cdhu
}

// ClearTags clears the value of the "tags" field.
func (cdhu *CustomDomainHistoryUpdate) ClearTags() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearTags()
	return cdhu
}

// SetOwnerID sets the "owner_id" field.
func (cdhu *CustomDomainHistoryUpdate) SetOwnerID(s string) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetOwnerID(s)
	return cdhu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cdhu *CustomDomainHistoryUpdate) SetNillableOwnerID(s *string) *CustomDomainHistoryUpdate {
	if s != nil {
		cdhu.SetOwnerID(*s)
	}
	return cdhu
}

// ClearOwnerID clears the value of the "owner_id" field.
func (cdhu *CustomDomainHistoryUpdate) ClearOwnerID() *CustomDomainHistoryUpdate {
	cdhu.mutation.ClearOwnerID()
	return cdhu
}

// SetStatus sets the "status" field.
func (cdhu *CustomDomainHistoryUpdate) SetStatus(eds enums.CustomDomainStatus) *CustomDomainHistoryUpdate {
	cdhu.mutation.SetStatus(eds)
	return cdhu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cdhu *CustomDomainHistoryUpdate) SetNillableStatus(eds *enums.CustomDomainStatus) *CustomDomainHistoryUpdate {
	if eds != nil {
		cdhu.SetStatus(*eds)
	}
	return cdhu
}

// Mutation returns the CustomDomainHistoryMutation object of the builder.
func (cdhu *CustomDomainHistoryUpdate) Mutation() *CustomDomainHistoryMutation {
	return cdhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cdhu *CustomDomainHistoryUpdate) Save(ctx context.Context) (int, error) {
	cdhu.defaults()
	return withHooks(ctx, cdhu.sqlSave, cdhu.mutation, cdhu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cdhu *CustomDomainHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cdhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cdhu *CustomDomainHistoryUpdate) Exec(ctx context.Context) error {
	_, err := cdhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdhu *CustomDomainHistoryUpdate) ExecX(ctx context.Context) {
	if err := cdhu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdhu *CustomDomainHistoryUpdate) defaults() {
	if _, ok := cdhu.mutation.UpdatedAt(); !ok && !cdhu.mutation.UpdatedAtCleared() {
		v := customdomainhistory.UpdateDefaultUpdatedAt()
		cdhu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cdhu *CustomDomainHistoryUpdate) check() error {
	if v, ok := cdhu.mutation.Status(); ok {
		if err := customdomainhistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "CustomDomainHistory.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cdhu *CustomDomainHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomDomainHistoryUpdate {
	cdhu.modifiers = append(cdhu.modifiers, modifiers...)
	return cdhu
}

func (cdhu *CustomDomainHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cdhu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(customdomainhistory.Table, customdomainhistory.Columns, sqlgraph.NewFieldSpec(customdomainhistory.FieldID, field.TypeString))
	if ps := cdhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cdhu.mutation.RefCleared() {
		_spec.ClearField(customdomainhistory.FieldRef, field.TypeString)
	}
	if cdhu.mutation.CreatedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cdhu.mutation.UpdatedAt(); ok {
		_spec.SetField(customdomainhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if cdhu.mutation.UpdatedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldUpdatedAt, field.TypeTime)
	}
	if cdhu.mutation.CreatedByCleared() {
		_spec.ClearField(customdomainhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := cdhu.mutation.UpdatedBy(); ok {
		_spec.SetField(customdomainhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if cdhu.mutation.UpdatedByCleared() {
		_spec.ClearField(customdomainhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := cdhu.mutation.DeletedAt(); ok {
		_spec.SetField(customdomainhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if cdhu.mutation.DeletedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cdhu.mutation.DeletedBy(); ok {
		_spec.SetField(customdomainhistory.FieldDeletedBy, field.TypeString, value)
	}
	if cdhu.mutation.DeletedByCleared() {
		_spec.ClearField(customdomainhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := cdhu.mutation.Tags(); ok {
		_spec.SetField(customdomainhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := cdhu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, customdomainhistory.FieldTags, value)
		})
	}
	if cdhu.mutation.TagsCleared() {
		_spec.ClearField(customdomainhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := cdhu.mutation.OwnerID(); ok {
		_spec.SetField(customdomainhistory.FieldOwnerID, field.TypeString, value)
	}
	if cdhu.mutation.OwnerIDCleared() {
		_spec.ClearField(customdomainhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := cdhu.mutation.Status(); ok {
		_spec.SetField(customdomainhistory.FieldStatus, field.TypeEnum, value)
	}
	_spec.Node.Schema = cdhu.schemaConfig.CustomDomainHistory
	ctx = internal.NewSchemaConfigContext(ctx, cdhu.schemaConfig)
	_spec.AddModifiers(cdhu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, cdhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customdomainhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cdhu.mutation.done = true
	return n, nil
}

// CustomDomainHistoryUpdateOne is the builder for updating a single CustomDomainHistory entity.
type CustomDomainHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CustomDomainHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetUpdatedAt(t time.Time) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetUpdatedAt(t)
	return cdhuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearUpdatedAt() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearUpdatedAt()
	return cdhuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetUpdatedBy(s string) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetUpdatedBy(s)
	return cdhuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cdhuo *CustomDomainHistoryUpdateOne) SetNillableUpdatedBy(s *string) *CustomDomainHistoryUpdateOne {
	if s != nil {
		cdhuo.SetUpdatedBy(*s)
	}
	return cdhuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearUpdatedBy() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearUpdatedBy()
	return cdhuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetDeletedAt(t time.Time) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetDeletedAt(t)
	return cdhuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cdhuo *CustomDomainHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *CustomDomainHistoryUpdateOne {
	if t != nil {
		cdhuo.SetDeletedAt(*t)
	}
	return cdhuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearDeletedAt() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearDeletedAt()
	return cdhuo
}

// SetDeletedBy sets the "deleted_by" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetDeletedBy(s string) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetDeletedBy(s)
	return cdhuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cdhuo *CustomDomainHistoryUpdateOne) SetNillableDeletedBy(s *string) *CustomDomainHistoryUpdateOne {
	if s != nil {
		cdhuo.SetDeletedBy(*s)
	}
	return cdhuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearDeletedBy() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearDeletedBy()
	return cdhuo
}

// SetTags sets the "tags" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetTags(s []string) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetTags(s)
	return cdhuo
}

// AppendTags appends s to the "tags" field.
func (cdhuo *CustomDomainHistoryUpdateOne) AppendTags(s []string) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.AppendTags(s)
	return cdhuo
}

// ClearTags clears the value of the "tags" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearTags() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearTags()
	return cdhuo
}

// SetOwnerID sets the "owner_id" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetOwnerID(s string) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetOwnerID(s)
	return cdhuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (cdhuo *CustomDomainHistoryUpdateOne) SetNillableOwnerID(s *string) *CustomDomainHistoryUpdateOne {
	if s != nil {
		cdhuo.SetOwnerID(*s)
	}
	return cdhuo
}

// ClearOwnerID clears the value of the "owner_id" field.
func (cdhuo *CustomDomainHistoryUpdateOne) ClearOwnerID() *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.ClearOwnerID()
	return cdhuo
}

// SetStatus sets the "status" field.
func (cdhuo *CustomDomainHistoryUpdateOne) SetStatus(eds enums.CustomDomainStatus) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.SetStatus(eds)
	return cdhuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cdhuo *CustomDomainHistoryUpdateOne) SetNillableStatus(eds *enums.CustomDomainStatus) *CustomDomainHistoryUpdateOne {
	if eds != nil {
		cdhuo.SetStatus(*eds)
	}
	return cdhuo
}

// Mutation returns the CustomDomainHistoryMutation object of the builder.
func (cdhuo *CustomDomainHistoryUpdateOne) Mutation() *CustomDomainHistoryMutation {
	return cdhuo.mutation
}

// Where appends a list predicates to the CustomDomainHistoryUpdate builder.
func (cdhuo *CustomDomainHistoryUpdateOne) Where(ps ...predicate.CustomDomainHistory) *CustomDomainHistoryUpdateOne {
	cdhuo.mutation.Where(ps...)
	return cdhuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cdhuo *CustomDomainHistoryUpdateOne) Select(field string, fields ...string) *CustomDomainHistoryUpdateOne {
	cdhuo.fields = append([]string{field}, fields...)
	return cdhuo
}

// Save executes the query and returns the updated CustomDomainHistory entity.
func (cdhuo *CustomDomainHistoryUpdateOne) Save(ctx context.Context) (*CustomDomainHistory, error) {
	cdhuo.defaults()
	return withHooks(ctx, cdhuo.sqlSave, cdhuo.mutation, cdhuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cdhuo *CustomDomainHistoryUpdateOne) SaveX(ctx context.Context) *CustomDomainHistory {
	node, err := cdhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cdhuo *CustomDomainHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cdhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cdhuo *CustomDomainHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := cdhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cdhuo *CustomDomainHistoryUpdateOne) defaults() {
	if _, ok := cdhuo.mutation.UpdatedAt(); !ok && !cdhuo.mutation.UpdatedAtCleared() {
		v := customdomainhistory.UpdateDefaultUpdatedAt()
		cdhuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cdhuo *CustomDomainHistoryUpdateOne) check() error {
	if v, ok := cdhuo.mutation.Status(); ok {
		if err := customdomainhistory.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "CustomDomainHistory.status": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cdhuo *CustomDomainHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CustomDomainHistoryUpdateOne {
	cdhuo.modifiers = append(cdhuo.modifiers, modifiers...)
	return cdhuo
}

func (cdhuo *CustomDomainHistoryUpdateOne) sqlSave(ctx context.Context) (_node *CustomDomainHistory, err error) {
	if err := cdhuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(customdomainhistory.Table, customdomainhistory.Columns, sqlgraph.NewFieldSpec(customdomainhistory.FieldID, field.TypeString))
	id, ok := cdhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "CustomDomainHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cdhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customdomainhistory.FieldID)
		for _, f := range fields {
			if !customdomainhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != customdomainhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cdhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if cdhuo.mutation.RefCleared() {
		_spec.ClearField(customdomainhistory.FieldRef, field.TypeString)
	}
	if cdhuo.mutation.CreatedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := cdhuo.mutation.UpdatedAt(); ok {
		_spec.SetField(customdomainhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if cdhuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldUpdatedAt, field.TypeTime)
	}
	if cdhuo.mutation.CreatedByCleared() {
		_spec.ClearField(customdomainhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := cdhuo.mutation.UpdatedBy(); ok {
		_spec.SetField(customdomainhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if cdhuo.mutation.UpdatedByCleared() {
		_spec.ClearField(customdomainhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := cdhuo.mutation.DeletedAt(); ok {
		_spec.SetField(customdomainhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if cdhuo.mutation.DeletedAtCleared() {
		_spec.ClearField(customdomainhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := cdhuo.mutation.DeletedBy(); ok {
		_spec.SetField(customdomainhistory.FieldDeletedBy, field.TypeString, value)
	}
	if cdhuo.mutation.DeletedByCleared() {
		_spec.ClearField(customdomainhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := cdhuo.mutation.Tags(); ok {
		_spec.SetField(customdomainhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := cdhuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, customdomainhistory.FieldTags, value)
		})
	}
	if cdhuo.mutation.TagsCleared() {
		_spec.ClearField(customdomainhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := cdhuo.mutation.OwnerID(); ok {
		_spec.SetField(customdomainhistory.FieldOwnerID, field.TypeString, value)
	}
	if cdhuo.mutation.OwnerIDCleared() {
		_spec.ClearField(customdomainhistory.FieldOwnerID, field.TypeString)
	}
	if value, ok := cdhuo.mutation.Status(); ok {
		_spec.SetField(customdomainhistory.FieldStatus, field.TypeEnum, value)
	}
	_spec.Node.Schema = cdhuo.schemaConfig.CustomDomainHistory
	ctx = internal.NewSchemaConfigContext(ctx, cdhuo.schemaConfig)
	_spec.AddModifiers(cdhuo.modifiers...)
	_node = &CustomDomainHistory{config: cdhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cdhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customdomainhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cdhuo.mutation.done = true
	return _node, nil
}
