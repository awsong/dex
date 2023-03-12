// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/awsong/dex/storage/ent/db/oauth2client"
	"github.com/awsong/dex/storage/ent/db/predicate"
)

// OAuth2ClientDelete is the builder for deleting a OAuth2Client entity.
type OAuth2ClientDelete struct {
	config
	hooks    []Hook
	mutation *OAuth2ClientMutation
}

// Where appends a list predicates to the OAuth2ClientDelete builder.
func (od *OAuth2ClientDelete) Where(ps ...predicate.OAuth2Client) *OAuth2ClientDelete {
	od.mutation.Where(ps...)
	return od
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (od *OAuth2ClientDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OAuth2ClientMutation](ctx, od.sqlExec, od.mutation, od.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (od *OAuth2ClientDelete) ExecX(ctx context.Context) int {
	n, err := od.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (od *OAuth2ClientDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(oauth2client.Table, sqlgraph.NewFieldSpec(oauth2client.FieldID, field.TypeString))
	if ps := od.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, od.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	od.mutation.done = true
	return affected, err
}

// OAuth2ClientDeleteOne is the builder for deleting a single OAuth2Client entity.
type OAuth2ClientDeleteOne struct {
	od *OAuth2ClientDelete
}

// Where appends a list predicates to the OAuth2ClientDelete builder.
func (odo *OAuth2ClientDeleteOne) Where(ps ...predicate.OAuth2Client) *OAuth2ClientDeleteOne {
	odo.od.mutation.Where(ps...)
	return odo
}

// Exec executes the deletion query.
func (odo *OAuth2ClientDeleteOne) Exec(ctx context.Context) error {
	n, err := odo.od.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{oauth2client.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (odo *OAuth2ClientDeleteOne) ExecX(ctx context.Context) {
	if err := odo.Exec(ctx); err != nil {
		panic(err)
	}
}
