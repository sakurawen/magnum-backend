// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"magnum/ent/formsubmission"
	"magnum/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FormSubmissionDelete is the builder for deleting a FormSubmission entity.
type FormSubmissionDelete struct {
	config
	hooks    []Hook
	mutation *FormSubmissionMutation
}

// Where appends a list predicates to the FormSubmissionDelete builder.
func (fsd *FormSubmissionDelete) Where(ps ...predicate.FormSubmission) *FormSubmissionDelete {
	fsd.mutation.Where(ps...)
	return fsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fsd *FormSubmissionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FormSubmissionMutation](ctx, fsd.sqlExec, fsd.mutation, fsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fsd *FormSubmissionDelete) ExecX(ctx context.Context) int {
	n, err := fsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fsd *FormSubmissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(formsubmission.Table, sqlgraph.NewFieldSpec(formsubmission.FieldID, field.TypeString))
	if ps := fsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fsd.mutation.done = true
	return affected, err
}

// FormSubmissionDeleteOne is the builder for deleting a single FormSubmission entity.
type FormSubmissionDeleteOne struct {
	fsd *FormSubmissionDelete
}

// Where appends a list predicates to the FormSubmissionDelete builder.
func (fsdo *FormSubmissionDeleteOne) Where(ps ...predicate.FormSubmission) *FormSubmissionDeleteOne {
	fsdo.fsd.mutation.Where(ps...)
	return fsdo
}

// Exec executes the deletion query.
func (fsdo *FormSubmissionDeleteOne) Exec(ctx context.Context) error {
	n, err := fsdo.fsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{formsubmission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fsdo *FormSubmissionDeleteOne) ExecX(ctx context.Context) {
	if err := fsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
