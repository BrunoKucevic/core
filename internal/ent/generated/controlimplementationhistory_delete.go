// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/controlimplementationhistory"
	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// ControlImplementationHistoryDelete is the builder for deleting a ControlImplementationHistory entity.
type ControlImplementationHistoryDelete struct {
	config
	hooks    []Hook
	mutation *ControlImplementationHistoryMutation
}

// Where appends a list predicates to the ControlImplementationHistoryDelete builder.
func (cihd *ControlImplementationHistoryDelete) Where(ps ...predicate.ControlImplementationHistory) *ControlImplementationHistoryDelete {
	cihd.mutation.Where(ps...)
	return cihd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cihd *ControlImplementationHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cihd.sqlExec, cihd.mutation, cihd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cihd *ControlImplementationHistoryDelete) ExecX(ctx context.Context) int {
	n, err := cihd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cihd *ControlImplementationHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(controlimplementationhistory.Table, sqlgraph.NewFieldSpec(controlimplementationhistory.FieldID, field.TypeString))
	_spec.Node.Schema = cihd.schemaConfig.ControlImplementationHistory
	ctx = internal.NewSchemaConfigContext(ctx, cihd.schemaConfig)
	if ps := cihd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cihd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cihd.mutation.done = true
	return affected, err
}

// ControlImplementationHistoryDeleteOne is the builder for deleting a single ControlImplementationHistory entity.
type ControlImplementationHistoryDeleteOne struct {
	cihd *ControlImplementationHistoryDelete
}

// Where appends a list predicates to the ControlImplementationHistoryDelete builder.
func (cihdo *ControlImplementationHistoryDeleteOne) Where(ps ...predicate.ControlImplementationHistory) *ControlImplementationHistoryDeleteOne {
	cihdo.cihd.mutation.Where(ps...)
	return cihdo
}

// Exec executes the deletion query.
func (cihdo *ControlImplementationHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := cihdo.cihd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{controlimplementationhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cihdo *ControlImplementationHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := cihdo.Exec(ctx); err != nil {
		panic(err)
	}
}
