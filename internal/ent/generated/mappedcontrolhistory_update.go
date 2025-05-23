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
	"github.com/theopenlane/core/internal/ent/generated/mappedcontrolhistory"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// MappedControlHistoryUpdate is the builder for updating MappedControlHistory entities.
type MappedControlHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *MappedControlHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MappedControlHistoryUpdate builder.
func (mchu *MappedControlHistoryUpdate) Where(ps ...predicate.MappedControlHistory) *MappedControlHistoryUpdate {
	mchu.mutation.Where(ps...)
	return mchu
}

// SetUpdatedAt sets the "updated_at" field.
func (mchu *MappedControlHistoryUpdate) SetUpdatedAt(t time.Time) *MappedControlHistoryUpdate {
	mchu.mutation.SetUpdatedAt(t)
	return mchu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mchu *MappedControlHistoryUpdate) ClearUpdatedAt() *MappedControlHistoryUpdate {
	mchu.mutation.ClearUpdatedAt()
	return mchu
}

// SetUpdatedBy sets the "updated_by" field.
func (mchu *MappedControlHistoryUpdate) SetUpdatedBy(s string) *MappedControlHistoryUpdate {
	mchu.mutation.SetUpdatedBy(s)
	return mchu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mchu *MappedControlHistoryUpdate) SetNillableUpdatedBy(s *string) *MappedControlHistoryUpdate {
	if s != nil {
		mchu.SetUpdatedBy(*s)
	}
	return mchu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mchu *MappedControlHistoryUpdate) ClearUpdatedBy() *MappedControlHistoryUpdate {
	mchu.mutation.ClearUpdatedBy()
	return mchu
}

// SetDeletedAt sets the "deleted_at" field.
func (mchu *MappedControlHistoryUpdate) SetDeletedAt(t time.Time) *MappedControlHistoryUpdate {
	mchu.mutation.SetDeletedAt(t)
	return mchu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mchu *MappedControlHistoryUpdate) SetNillableDeletedAt(t *time.Time) *MappedControlHistoryUpdate {
	if t != nil {
		mchu.SetDeletedAt(*t)
	}
	return mchu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mchu *MappedControlHistoryUpdate) ClearDeletedAt() *MappedControlHistoryUpdate {
	mchu.mutation.ClearDeletedAt()
	return mchu
}

// SetDeletedBy sets the "deleted_by" field.
func (mchu *MappedControlHistoryUpdate) SetDeletedBy(s string) *MappedControlHistoryUpdate {
	mchu.mutation.SetDeletedBy(s)
	return mchu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (mchu *MappedControlHistoryUpdate) SetNillableDeletedBy(s *string) *MappedControlHistoryUpdate {
	if s != nil {
		mchu.SetDeletedBy(*s)
	}
	return mchu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (mchu *MappedControlHistoryUpdate) ClearDeletedBy() *MappedControlHistoryUpdate {
	mchu.mutation.ClearDeletedBy()
	return mchu
}

// SetTags sets the "tags" field.
func (mchu *MappedControlHistoryUpdate) SetTags(s []string) *MappedControlHistoryUpdate {
	mchu.mutation.SetTags(s)
	return mchu
}

// AppendTags appends s to the "tags" field.
func (mchu *MappedControlHistoryUpdate) AppendTags(s []string) *MappedControlHistoryUpdate {
	mchu.mutation.AppendTags(s)
	return mchu
}

// ClearTags clears the value of the "tags" field.
func (mchu *MappedControlHistoryUpdate) ClearTags() *MappedControlHistoryUpdate {
	mchu.mutation.ClearTags()
	return mchu
}

// SetMappingType sets the "mapping_type" field.
func (mchu *MappedControlHistoryUpdate) SetMappingType(s string) *MappedControlHistoryUpdate {
	mchu.mutation.SetMappingType(s)
	return mchu
}

// SetNillableMappingType sets the "mapping_type" field if the given value is not nil.
func (mchu *MappedControlHistoryUpdate) SetNillableMappingType(s *string) *MappedControlHistoryUpdate {
	if s != nil {
		mchu.SetMappingType(*s)
	}
	return mchu
}

// ClearMappingType clears the value of the "mapping_type" field.
func (mchu *MappedControlHistoryUpdate) ClearMappingType() *MappedControlHistoryUpdate {
	mchu.mutation.ClearMappingType()
	return mchu
}

// SetRelation sets the "relation" field.
func (mchu *MappedControlHistoryUpdate) SetRelation(s string) *MappedControlHistoryUpdate {
	mchu.mutation.SetRelation(s)
	return mchu
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (mchu *MappedControlHistoryUpdate) SetNillableRelation(s *string) *MappedControlHistoryUpdate {
	if s != nil {
		mchu.SetRelation(*s)
	}
	return mchu
}

// ClearRelation clears the value of the "relation" field.
func (mchu *MappedControlHistoryUpdate) ClearRelation() *MappedControlHistoryUpdate {
	mchu.mutation.ClearRelation()
	return mchu
}

// Mutation returns the MappedControlHistoryMutation object of the builder.
func (mchu *MappedControlHistoryUpdate) Mutation() *MappedControlHistoryMutation {
	return mchu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mchu *MappedControlHistoryUpdate) Save(ctx context.Context) (int, error) {
	mchu.defaults()
	return withHooks(ctx, mchu.sqlSave, mchu.mutation, mchu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mchu *MappedControlHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := mchu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mchu *MappedControlHistoryUpdate) Exec(ctx context.Context) error {
	_, err := mchu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mchu *MappedControlHistoryUpdate) ExecX(ctx context.Context) {
	if err := mchu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mchu *MappedControlHistoryUpdate) defaults() {
	if _, ok := mchu.mutation.UpdatedAt(); !ok && !mchu.mutation.UpdatedAtCleared() {
		v := mappedcontrolhistory.UpdateDefaultUpdatedAt()
		mchu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mchu *MappedControlHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MappedControlHistoryUpdate {
	mchu.modifiers = append(mchu.modifiers, modifiers...)
	return mchu
}

func (mchu *MappedControlHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(mappedcontrolhistory.Table, mappedcontrolhistory.Columns, sqlgraph.NewFieldSpec(mappedcontrolhistory.FieldID, field.TypeString))
	if ps := mchu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mchu.mutation.RefCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldRef, field.TypeString)
	}
	if mchu.mutation.CreatedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mchu.mutation.UpdatedAt(); ok {
		_spec.SetField(mappedcontrolhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if mchu.mutation.UpdatedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldUpdatedAt, field.TypeTime)
	}
	if mchu.mutation.CreatedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := mchu.mutation.UpdatedBy(); ok {
		_spec.SetField(mappedcontrolhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if mchu.mutation.UpdatedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := mchu.mutation.DeletedAt(); ok {
		_spec.SetField(mappedcontrolhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if mchu.mutation.DeletedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := mchu.mutation.DeletedBy(); ok {
		_spec.SetField(mappedcontrolhistory.FieldDeletedBy, field.TypeString, value)
	}
	if mchu.mutation.DeletedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := mchu.mutation.Tags(); ok {
		_spec.SetField(mappedcontrolhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := mchu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, mappedcontrolhistory.FieldTags, value)
		})
	}
	if mchu.mutation.TagsCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := mchu.mutation.MappingType(); ok {
		_spec.SetField(mappedcontrolhistory.FieldMappingType, field.TypeString, value)
	}
	if mchu.mutation.MappingTypeCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldMappingType, field.TypeString)
	}
	if value, ok := mchu.mutation.Relation(); ok {
		_spec.SetField(mappedcontrolhistory.FieldRelation, field.TypeString, value)
	}
	if mchu.mutation.RelationCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldRelation, field.TypeString)
	}
	_spec.Node.Schema = mchu.schemaConfig.MappedControlHistory
	ctx = internal.NewSchemaConfigContext(ctx, mchu.schemaConfig)
	_spec.AddModifiers(mchu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mchu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mappedcontrolhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mchu.mutation.done = true
	return n, nil
}

// MappedControlHistoryUpdateOne is the builder for updating a single MappedControlHistory entity.
type MappedControlHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MappedControlHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (mchuo *MappedControlHistoryUpdateOne) SetUpdatedAt(t time.Time) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetUpdatedAt(t)
	return mchuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearUpdatedAt() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearUpdatedAt()
	return mchuo
}

// SetUpdatedBy sets the "updated_by" field.
func (mchuo *MappedControlHistoryUpdateOne) SetUpdatedBy(s string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetUpdatedBy(s)
	return mchuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mchuo *MappedControlHistoryUpdateOne) SetNillableUpdatedBy(s *string) *MappedControlHistoryUpdateOne {
	if s != nil {
		mchuo.SetUpdatedBy(*s)
	}
	return mchuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearUpdatedBy() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearUpdatedBy()
	return mchuo
}

// SetDeletedAt sets the "deleted_at" field.
func (mchuo *MappedControlHistoryUpdateOne) SetDeletedAt(t time.Time) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetDeletedAt(t)
	return mchuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (mchuo *MappedControlHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *MappedControlHistoryUpdateOne {
	if t != nil {
		mchuo.SetDeletedAt(*t)
	}
	return mchuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearDeletedAt() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearDeletedAt()
	return mchuo
}

// SetDeletedBy sets the "deleted_by" field.
func (mchuo *MappedControlHistoryUpdateOne) SetDeletedBy(s string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetDeletedBy(s)
	return mchuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (mchuo *MappedControlHistoryUpdateOne) SetNillableDeletedBy(s *string) *MappedControlHistoryUpdateOne {
	if s != nil {
		mchuo.SetDeletedBy(*s)
	}
	return mchuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearDeletedBy() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearDeletedBy()
	return mchuo
}

// SetTags sets the "tags" field.
func (mchuo *MappedControlHistoryUpdateOne) SetTags(s []string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetTags(s)
	return mchuo
}

// AppendTags appends s to the "tags" field.
func (mchuo *MappedControlHistoryUpdateOne) AppendTags(s []string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.AppendTags(s)
	return mchuo
}

// ClearTags clears the value of the "tags" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearTags() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearTags()
	return mchuo
}

// SetMappingType sets the "mapping_type" field.
func (mchuo *MappedControlHistoryUpdateOne) SetMappingType(s string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetMappingType(s)
	return mchuo
}

// SetNillableMappingType sets the "mapping_type" field if the given value is not nil.
func (mchuo *MappedControlHistoryUpdateOne) SetNillableMappingType(s *string) *MappedControlHistoryUpdateOne {
	if s != nil {
		mchuo.SetMappingType(*s)
	}
	return mchuo
}

// ClearMappingType clears the value of the "mapping_type" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearMappingType() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearMappingType()
	return mchuo
}

// SetRelation sets the "relation" field.
func (mchuo *MappedControlHistoryUpdateOne) SetRelation(s string) *MappedControlHistoryUpdateOne {
	mchuo.mutation.SetRelation(s)
	return mchuo
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (mchuo *MappedControlHistoryUpdateOne) SetNillableRelation(s *string) *MappedControlHistoryUpdateOne {
	if s != nil {
		mchuo.SetRelation(*s)
	}
	return mchuo
}

// ClearRelation clears the value of the "relation" field.
func (mchuo *MappedControlHistoryUpdateOne) ClearRelation() *MappedControlHistoryUpdateOne {
	mchuo.mutation.ClearRelation()
	return mchuo
}

// Mutation returns the MappedControlHistoryMutation object of the builder.
func (mchuo *MappedControlHistoryUpdateOne) Mutation() *MappedControlHistoryMutation {
	return mchuo.mutation
}

// Where appends a list predicates to the MappedControlHistoryUpdate builder.
func (mchuo *MappedControlHistoryUpdateOne) Where(ps ...predicate.MappedControlHistory) *MappedControlHistoryUpdateOne {
	mchuo.mutation.Where(ps...)
	return mchuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mchuo *MappedControlHistoryUpdateOne) Select(field string, fields ...string) *MappedControlHistoryUpdateOne {
	mchuo.fields = append([]string{field}, fields...)
	return mchuo
}

// Save executes the query and returns the updated MappedControlHistory entity.
func (mchuo *MappedControlHistoryUpdateOne) Save(ctx context.Context) (*MappedControlHistory, error) {
	mchuo.defaults()
	return withHooks(ctx, mchuo.sqlSave, mchuo.mutation, mchuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mchuo *MappedControlHistoryUpdateOne) SaveX(ctx context.Context) *MappedControlHistory {
	node, err := mchuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mchuo *MappedControlHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := mchuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mchuo *MappedControlHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := mchuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mchuo *MappedControlHistoryUpdateOne) defaults() {
	if _, ok := mchuo.mutation.UpdatedAt(); !ok && !mchuo.mutation.UpdatedAtCleared() {
		v := mappedcontrolhistory.UpdateDefaultUpdatedAt()
		mchuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mchuo *MappedControlHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MappedControlHistoryUpdateOne {
	mchuo.modifiers = append(mchuo.modifiers, modifiers...)
	return mchuo
}

func (mchuo *MappedControlHistoryUpdateOne) sqlSave(ctx context.Context) (_node *MappedControlHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(mappedcontrolhistory.Table, mappedcontrolhistory.Columns, sqlgraph.NewFieldSpec(mappedcontrolhistory.FieldID, field.TypeString))
	id, ok := mchuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "MappedControlHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mchuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mappedcontrolhistory.FieldID)
		for _, f := range fields {
			if !mappedcontrolhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != mappedcontrolhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mchuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mchuo.mutation.RefCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldRef, field.TypeString)
	}
	if mchuo.mutation.CreatedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mchuo.mutation.UpdatedAt(); ok {
		_spec.SetField(mappedcontrolhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if mchuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldUpdatedAt, field.TypeTime)
	}
	if mchuo.mutation.CreatedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := mchuo.mutation.UpdatedBy(); ok {
		_spec.SetField(mappedcontrolhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if mchuo.mutation.UpdatedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := mchuo.mutation.DeletedAt(); ok {
		_spec.SetField(mappedcontrolhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if mchuo.mutation.DeletedAtCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := mchuo.mutation.DeletedBy(); ok {
		_spec.SetField(mappedcontrolhistory.FieldDeletedBy, field.TypeString, value)
	}
	if mchuo.mutation.DeletedByCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := mchuo.mutation.Tags(); ok {
		_spec.SetField(mappedcontrolhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := mchuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, mappedcontrolhistory.FieldTags, value)
		})
	}
	if mchuo.mutation.TagsCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := mchuo.mutation.MappingType(); ok {
		_spec.SetField(mappedcontrolhistory.FieldMappingType, field.TypeString, value)
	}
	if mchuo.mutation.MappingTypeCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldMappingType, field.TypeString)
	}
	if value, ok := mchuo.mutation.Relation(); ok {
		_spec.SetField(mappedcontrolhistory.FieldRelation, field.TypeString, value)
	}
	if mchuo.mutation.RelationCleared() {
		_spec.ClearField(mappedcontrolhistory.FieldRelation, field.TypeString)
	}
	_spec.Node.Schema = mchuo.schemaConfig.MappedControlHistory
	ctx = internal.NewSchemaConfigContext(ctx, mchuo.schemaConfig)
	_spec.AddModifiers(mchuo.modifiers...)
	_node = &MappedControlHistory{config: mchuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mchuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mappedcontrolhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mchuo.mutation.done = true
	return _node, nil
}
