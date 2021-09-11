// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/predicate"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/unavailability"
)

// UnavailabilityUpdate is the builder for updating Unavailability entities.
type UnavailabilityUpdate struct {
	config
	hooks    []Hook
	mutation *UnavailabilityMutation
}

// Where appends a list predicates to the UnavailabilityUpdate builder.
func (uu *UnavailabilityUpdate) Where(ps ...predicate.Unavailability) *UnavailabilityUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updatedAt" field.
func (uu *UnavailabilityUpdate) SetUpdatedAt(t time.Time) *UnavailabilityUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetStartTime sets the "startTime" field.
func (uu *UnavailabilityUpdate) SetStartTime(t time.Time) *UnavailabilityUpdate {
	uu.mutation.SetStartTime(t)
	return uu
}

// SetEndTime sets the "endTime" field.
func (uu *UnavailabilityUpdate) SetEndTime(t time.Time) *UnavailabilityUpdate {
	uu.mutation.SetEndTime(t)
	return uu
}

// SetResourceId sets the "resourceId" field.
func (uu *UnavailabilityUpdate) SetResourceId(i int) *UnavailabilityUpdate {
	uu.mutation.SetResourceId(i)
	return uu
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (uu *UnavailabilityUpdate) SetResourceID(id int) *UnavailabilityUpdate {
	uu.mutation.SetResourceID(id)
	return uu
}

// SetResource sets the "resource" edge to the Resource entity.
func (uu *UnavailabilityUpdate) SetResource(r *Resource) *UnavailabilityUpdate {
	return uu.SetResourceID(r.ID)
}

// Mutation returns the UnavailabilityMutation object of the builder.
func (uu *UnavailabilityUpdate) Mutation() *UnavailabilityMutation {
	return uu.mutation
}

// ClearResource clears the "resource" edge to the Resource entity.
func (uu *UnavailabilityUpdate) ClearResource() *UnavailabilityUpdate {
	uu.mutation.ClearResource()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UnavailabilityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnavailabilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UnavailabilityUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UnavailabilityUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UnavailabilityUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UnavailabilityUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := unavailability.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UnavailabilityUpdate) check() error {
	if _, ok := uu.mutation.ResourceID(); uu.mutation.ResourceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"resource\"")
	}
	return nil
}

func (uu *UnavailabilityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unavailability.Table,
			Columns: unavailability.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unavailability.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldStartTime,
		})
	}
	if value, ok := uu.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldEndTime,
		})
	}
	if uu.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unavailability.ResourceTable,
			Columns: []string{unavailability.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unavailability.ResourceTable,
			Columns: []string{unavailability.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unavailability.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UnavailabilityUpdateOne is the builder for updating a single Unavailability entity.
type UnavailabilityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UnavailabilityMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (uuo *UnavailabilityUpdateOne) SetUpdatedAt(t time.Time) *UnavailabilityUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetStartTime sets the "startTime" field.
func (uuo *UnavailabilityUpdateOne) SetStartTime(t time.Time) *UnavailabilityUpdateOne {
	uuo.mutation.SetStartTime(t)
	return uuo
}

// SetEndTime sets the "endTime" field.
func (uuo *UnavailabilityUpdateOne) SetEndTime(t time.Time) *UnavailabilityUpdateOne {
	uuo.mutation.SetEndTime(t)
	return uuo
}

// SetResourceId sets the "resourceId" field.
func (uuo *UnavailabilityUpdateOne) SetResourceId(i int) *UnavailabilityUpdateOne {
	uuo.mutation.SetResourceId(i)
	return uuo
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (uuo *UnavailabilityUpdateOne) SetResourceID(id int) *UnavailabilityUpdateOne {
	uuo.mutation.SetResourceID(id)
	return uuo
}

// SetResource sets the "resource" edge to the Resource entity.
func (uuo *UnavailabilityUpdateOne) SetResource(r *Resource) *UnavailabilityUpdateOne {
	return uuo.SetResourceID(r.ID)
}

// Mutation returns the UnavailabilityMutation object of the builder.
func (uuo *UnavailabilityUpdateOne) Mutation() *UnavailabilityMutation {
	return uuo.mutation
}

// ClearResource clears the "resource" edge to the Resource entity.
func (uuo *UnavailabilityUpdateOne) ClearResource() *UnavailabilityUpdateOne {
	uuo.mutation.ClearResource()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UnavailabilityUpdateOne) Select(field string, fields ...string) *UnavailabilityUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated Unavailability entity.
func (uuo *UnavailabilityUpdateOne) Save(ctx context.Context) (*Unavailability, error) {
	var (
		err  error
		node *Unavailability
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnavailabilityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UnavailabilityUpdateOne) SaveX(ctx context.Context) *Unavailability {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UnavailabilityUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UnavailabilityUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UnavailabilityUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := unavailability.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UnavailabilityUpdateOne) check() error {
	if _, ok := uuo.mutation.ResourceID(); uuo.mutation.ResourceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"resource\"")
	}
	return nil
}

func (uuo *UnavailabilityUpdateOne) sqlSave(ctx context.Context) (_node *Unavailability, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unavailability.Table,
			Columns: unavailability.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unavailability.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Unavailability.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, unavailability.FieldID)
		for _, f := range fields {
			if !unavailability.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != unavailability.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldStartTime,
		})
	}
	if value, ok := uuo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: unavailability.FieldEndTime,
		})
	}
	if uuo.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unavailability.ResourceTable,
			Columns: []string{unavailability.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   unavailability.ResourceTable,
			Columns: []string{unavailability.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: resource.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Unavailability{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unavailability.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
