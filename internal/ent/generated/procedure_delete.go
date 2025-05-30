// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"

	"github.com/theopenlane/core/internal/ent/generated/internal"
	"github.com/theopenlane/core/internal/ent/generated/procedure"
)

// ProcedureDelete is the builder for deleting a Procedure entity.
type ProcedureDelete struct {
	config
	hooks    []Hook
	mutation *ProcedureMutation
}

// Where appends a list predicates to the ProcedureDelete builder.
func (pd *ProcedureDelete) Where(ps ...predicate.Procedure) *ProcedureDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *ProcedureDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *ProcedureDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *ProcedureDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(procedure.Table, sqlgraph.NewFieldSpec(procedure.FieldID, field.TypeString))
	_spec.Node.Schema = pd.schemaConfig.Procedure
	ctx = internal.NewSchemaConfigContext(ctx, pd.schemaConfig)
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// ProcedureDeleteOne is the builder for deleting a single Procedure entity.
type ProcedureDeleteOne struct {
	pd *ProcedureDelete
}

// Where appends a list predicates to the ProcedureDelete builder.
func (pdo *ProcedureDeleteOne) Where(ps ...predicate.Procedure) *ProcedureDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *ProcedureDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{procedure.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *ProcedureDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
