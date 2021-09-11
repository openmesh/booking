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
	"github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/predicate"
	"github.com/openmesh/booking/ent/resource"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updatedAt" field.
func (bu *BookingUpdate) SetUpdatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(s string) *BookingUpdate {
	bu.mutation.SetStatus(s)
	return bu
}

// SetStartTime sets the "startTime" field.
func (bu *BookingUpdate) SetStartTime(t time.Time) *BookingUpdate {
	bu.mutation.SetStartTime(t)
	return bu
}

// SetEndTime sets the "endTime" field.
func (bu *BookingUpdate) SetEndTime(t time.Time) *BookingUpdate {
	bu.mutation.SetEndTime(t)
	return bu
}

// SetResourceId sets the "resourceId" field.
func (bu *BookingUpdate) SetResourceId(i int) *BookingUpdate {
	bu.mutation.SetResourceId(i)
	return bu
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (bu *BookingUpdate) SetResourceID(id int) *BookingUpdate {
	bu.mutation.SetResourceID(id)
	return bu
}

// SetResource sets the "resource" edge to the Resource entity.
func (bu *BookingUpdate) SetResource(r *Resource) *BookingUpdate {
	return bu.SetResourceID(r.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearResource clears the "resource" edge to the Resource entity.
func (bu *BookingUpdate) ClearResource() *BookingUpdate {
	bu.mutation.ClearResource()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	bu.defaults()
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookingUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if _, ok := bu.mutation.ResourceID(); bu.mutation.ResourceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"resource\"")
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldUpdatedAt,
		})
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: booking.FieldStatus,
		})
	}
	if value, ok := bu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldStartTime,
		})
	}
	if value, ok := bu.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldEndTime,
		})
	}
	if bu.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.ResourceTable,
			Columns: []string{booking.ResourceColumn},
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
	if nodes := bu.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.ResourceTable,
			Columns: []string{booking.ResourceColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (buo *BookingUpdateOne) SetUpdatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(s string) *BookingUpdateOne {
	buo.mutation.SetStatus(s)
	return buo
}

// SetStartTime sets the "startTime" field.
func (buo *BookingUpdateOne) SetStartTime(t time.Time) *BookingUpdateOne {
	buo.mutation.SetStartTime(t)
	return buo
}

// SetEndTime sets the "endTime" field.
func (buo *BookingUpdateOne) SetEndTime(t time.Time) *BookingUpdateOne {
	buo.mutation.SetEndTime(t)
	return buo
}

// SetResourceId sets the "resourceId" field.
func (buo *BookingUpdateOne) SetResourceId(i int) *BookingUpdateOne {
	buo.mutation.SetResourceId(i)
	return buo
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (buo *BookingUpdateOne) SetResourceID(id int) *BookingUpdateOne {
	buo.mutation.SetResourceID(id)
	return buo
}

// SetResource sets the "resource" edge to the Resource entity.
func (buo *BookingUpdateOne) SetResource(r *Resource) *BookingUpdateOne {
	return buo.SetResourceID(r.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearResource clears the "resource" edge to the Resource entity.
func (buo *BookingUpdateOne) ClearResource() *BookingUpdateOne {
	buo.mutation.ClearResource()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	var (
		err  error
		node *Booking
	)
	buo.defaults()
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookingUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if _, ok := buo.mutation.ResourceID(); buo.mutation.ResourceCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"resource\"")
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   booking.Table,
			Columns: booking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Booking.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldUpdatedAt,
		})
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: booking.FieldStatus,
		})
	}
	if value, ok := buo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldStartTime,
		})
	}
	if value, ok := buo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldEndTime,
		})
	}
	if buo.mutation.ResourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.ResourceTable,
			Columns: []string{booking.ResourceColumn},
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
	if nodes := buo.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.ResourceTable,
			Columns: []string{booking.ResourceColumn},
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
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}