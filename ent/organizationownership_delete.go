// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/organizationownership"
	"github.com/openmesh/booking/ent/predicate"
)

// OrganizationOwnershipDelete is the builder for deleting a OrganizationOwnership entity.
type OrganizationOwnershipDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationOwnershipMutation
}

// Where appends a list predicates to the OrganizationOwnershipDelete builder.
func (ood *OrganizationOwnershipDelete) Where(ps ...predicate.OrganizationOwnership) *OrganizationOwnershipDelete {
	ood.mutation.Where(ps...)
	return ood
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ood *OrganizationOwnershipDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ood.hooks) == 0 {
		affected, err = ood.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrganizationOwnershipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ood.mutation = mutation
			affected, err = ood.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ood.hooks) - 1; i >= 0; i-- {
			if ood.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ood.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ood.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ood *OrganizationOwnershipDelete) ExecX(ctx context.Context) int {
	n, err := ood.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ood *OrganizationOwnershipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: organizationownership.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: organizationownership.FieldID,
			},
		},
	}
	if ps := ood.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ood.driver, _spec)
}

// OrganizationOwnershipDeleteOne is the builder for deleting a single OrganizationOwnership entity.
type OrganizationOwnershipDeleteOne struct {
	ood *OrganizationOwnershipDelete
}

// Exec executes the deletion query.
func (oodo *OrganizationOwnershipDeleteOne) Exec(ctx context.Context) error {
	n, err := oodo.ood.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationownership.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (oodo *OrganizationOwnershipDeleteOne) ExecX(ctx context.Context) {
	oodo.ood.ExecX(ctx)
}
