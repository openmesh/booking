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
	"github.com/openmesh/booking/ent/organization"
	"github.com/openmesh/booking/ent/predicate"
	"github.com/openmesh/booking/ent/resource"
	"github.com/openmesh/booking/ent/slot"
	"github.com/openmesh/booking/ent/unavailability"
)

// ResourceUpdate is the builder for updating Resource entities.
type ResourceUpdate struct {
	config
	hooks    []Hook
	mutation *ResourceMutation
}

// Where appends a list predicates to the ResourceUpdate builder.
func (ru *ResourceUpdate) Where(ps ...predicate.Resource) *ResourceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUpdatedAt sets the "updatedAt" field.
func (ru *ResourceUpdate) SetUpdatedAt(t time.Time) *ResourceUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetName sets the "name" field.
func (ru *ResourceUpdate) SetName(s string) *ResourceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetDescription sets the "description" field.
func (ru *ResourceUpdate) SetDescription(s string) *ResourceUpdate {
	ru.mutation.SetDescription(s)
	return ru
}

// SetTimezone sets the "timezone" field.
func (ru *ResourceUpdate) SetTimezone(s string) *ResourceUpdate {
	ru.mutation.SetTimezone(s)
	return ru
}

// SetPassword sets the "password" field.
func (ru *ResourceUpdate) SetPassword(s string) *ResourceUpdate {
	ru.mutation.SetPassword(s)
	return ru
}

// SetPrice sets the "price" field.
func (ru *ResourceUpdate) SetPrice(i int) *ResourceUpdate {
	ru.mutation.ResetPrice()
	ru.mutation.SetPrice(i)
	return ru
}

// AddPrice adds i to the "price" field.
func (ru *ResourceUpdate) AddPrice(i int) *ResourceUpdate {
	ru.mutation.AddPrice(i)
	return ru
}

// SetBookingPrice sets the "bookingPrice" field.
func (ru *ResourceUpdate) SetBookingPrice(i int) *ResourceUpdate {
	ru.mutation.ResetBookingPrice()
	ru.mutation.SetBookingPrice(i)
	return ru
}

// AddBookingPrice adds i to the "bookingPrice" field.
func (ru *ResourceUpdate) AddBookingPrice(i int) *ResourceUpdate {
	ru.mutation.AddBookingPrice(i)
	return ru
}

// SetOrganizationId sets the "organizationId" field.
func (ru *ResourceUpdate) SetOrganizationId(i int) *ResourceUpdate {
	ru.mutation.SetOrganizationId(i)
	return ru
}

// SetQuantityAvailable sets the "quantityAvailable" field.
func (ru *ResourceUpdate) SetQuantityAvailable(i int) *ResourceUpdate {
	ru.mutation.ResetQuantityAvailable()
	ru.mutation.SetQuantityAvailable(i)
	return ru
}

// SetNillableQuantityAvailable sets the "quantityAvailable" field if the given value is not nil.
func (ru *ResourceUpdate) SetNillableQuantityAvailable(i *int) *ResourceUpdate {
	if i != nil {
		ru.SetQuantityAvailable(*i)
	}
	return ru
}

// AddQuantityAvailable adds i to the "quantityAvailable" field.
func (ru *ResourceUpdate) AddQuantityAvailable(i int) *ResourceUpdate {
	ru.mutation.AddQuantityAvailable(i)
	return ru
}

// ClearQuantityAvailable clears the value of the "quantityAvailable" field.
func (ru *ResourceUpdate) ClearQuantityAvailable() *ResourceUpdate {
	ru.mutation.ClearQuantityAvailable()
	return ru
}

// AddSlotIDs adds the "slots" edge to the Slot entity by IDs.
func (ru *ResourceUpdate) AddSlotIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddSlotIDs(ids...)
	return ru
}

// AddSlots adds the "slots" edges to the Slot entity.
func (ru *ResourceUpdate) AddSlots(s ...*Slot) *ResourceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddSlotIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (ru *ResourceUpdate) AddBookingIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddBookingIDs(ids...)
	return ru
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (ru *ResourceUpdate) AddBookings(b ...*Booking) *ResourceUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.AddBookingIDs(ids...)
}

// AddUnavailabilityIDs adds the "unavailabilities" edge to the Unavailability entity by IDs.
func (ru *ResourceUpdate) AddUnavailabilityIDs(ids ...int) *ResourceUpdate {
	ru.mutation.AddUnavailabilityIDs(ids...)
	return ru
}

// AddUnavailabilities adds the "unavailabilities" edges to the Unavailability entity.
func (ru *ResourceUpdate) AddUnavailabilities(u ...*Unavailability) *ResourceUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.AddUnavailabilityIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ru *ResourceUpdate) SetOrganizationID(id int) *ResourceUpdate {
	ru.mutation.SetOrganizationID(id)
	return ru
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ru *ResourceUpdate) SetOrganization(o *Organization) *ResourceUpdate {
	return ru.SetOrganizationID(o.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ru *ResourceUpdate) Mutation() *ResourceMutation {
	return ru.mutation
}

// ClearSlots clears all "slots" edges to the Slot entity.
func (ru *ResourceUpdate) ClearSlots() *ResourceUpdate {
	ru.mutation.ClearSlots()
	return ru
}

// RemoveSlotIDs removes the "slots" edge to Slot entities by IDs.
func (ru *ResourceUpdate) RemoveSlotIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemoveSlotIDs(ids...)
	return ru
}

// RemoveSlots removes "slots" edges to Slot entities.
func (ru *ResourceUpdate) RemoveSlots(s ...*Slot) *ResourceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveSlotIDs(ids...)
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (ru *ResourceUpdate) ClearBookings() *ResourceUpdate {
	ru.mutation.ClearBookings()
	return ru
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (ru *ResourceUpdate) RemoveBookingIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemoveBookingIDs(ids...)
	return ru
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (ru *ResourceUpdate) RemoveBookings(b ...*Booking) *ResourceUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ru.RemoveBookingIDs(ids...)
}

// ClearUnavailabilities clears all "unavailabilities" edges to the Unavailability entity.
func (ru *ResourceUpdate) ClearUnavailabilities() *ResourceUpdate {
	ru.mutation.ClearUnavailabilities()
	return ru
}

// RemoveUnavailabilityIDs removes the "unavailabilities" edge to Unavailability entities by IDs.
func (ru *ResourceUpdate) RemoveUnavailabilityIDs(ids ...int) *ResourceUpdate {
	ru.mutation.RemoveUnavailabilityIDs(ids...)
	return ru
}

// RemoveUnavailabilities removes "unavailabilities" edges to Unavailability entities.
func (ru *ResourceUpdate) RemoveUnavailabilities(u ...*Unavailability) *ResourceUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ru.RemoveUnavailabilityIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ru *ResourceUpdate) ClearOrganization() *ResourceUpdate {
	ru.mutation.ClearOrganization()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ResourceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ResourceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ResourceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ResourceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ResourceUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := resource.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ResourceUpdate) check() error {
	if _, ok := ru.mutation.OrganizationID(); ru.mutation.OrganizationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"organization\"")
	}
	return nil
}

func (ru *ResourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resource.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldName,
		})
	}
	if value, ok := ru.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldDescription,
		})
	}
	if value, ok := ru.mutation.Timezone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldTimezone,
		})
	}
	if value, ok := ru.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldPassword,
		})
	}
	if value, ok := ru.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldPrice,
		})
	}
	if value, ok := ru.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldPrice,
		})
	}
	if value, ok := ru.mutation.BookingPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldBookingPrice,
		})
	}
	if value, ok := ru.mutation.AddedBookingPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldBookingPrice,
		})
	}
	if value, ok := ru.mutation.QuantityAvailable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if value, ok := ru.mutation.AddedQuantityAvailable(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if ru.mutation.QuantityAvailableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if ru.mutation.SlotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedSlotsIDs(); len(nodes) > 0 && !ru.mutation.SlotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.SlotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !ru.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UnavailabilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUnavailabilitiesIDs(); len(nodes) > 0 && !ru.mutation.UnavailabilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UnavailabilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.OrganizationTable,
			Columns: []string{resource.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.OrganizationTable,
			Columns: []string{resource.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ResourceUpdateOne is the builder for updating a single Resource entity.
type ResourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ResourceMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (ruo *ResourceUpdateOne) SetUpdatedAt(t time.Time) *ResourceUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetName sets the "name" field.
func (ruo *ResourceUpdateOne) SetName(s string) *ResourceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetDescription sets the "description" field.
func (ruo *ResourceUpdateOne) SetDescription(s string) *ResourceUpdateOne {
	ruo.mutation.SetDescription(s)
	return ruo
}

// SetTimezone sets the "timezone" field.
func (ruo *ResourceUpdateOne) SetTimezone(s string) *ResourceUpdateOne {
	ruo.mutation.SetTimezone(s)
	return ruo
}

// SetPassword sets the "password" field.
func (ruo *ResourceUpdateOne) SetPassword(s string) *ResourceUpdateOne {
	ruo.mutation.SetPassword(s)
	return ruo
}

// SetPrice sets the "price" field.
func (ruo *ResourceUpdateOne) SetPrice(i int) *ResourceUpdateOne {
	ruo.mutation.ResetPrice()
	ruo.mutation.SetPrice(i)
	return ruo
}

// AddPrice adds i to the "price" field.
func (ruo *ResourceUpdateOne) AddPrice(i int) *ResourceUpdateOne {
	ruo.mutation.AddPrice(i)
	return ruo
}

// SetBookingPrice sets the "bookingPrice" field.
func (ruo *ResourceUpdateOne) SetBookingPrice(i int) *ResourceUpdateOne {
	ruo.mutation.ResetBookingPrice()
	ruo.mutation.SetBookingPrice(i)
	return ruo
}

// AddBookingPrice adds i to the "bookingPrice" field.
func (ruo *ResourceUpdateOne) AddBookingPrice(i int) *ResourceUpdateOne {
	ruo.mutation.AddBookingPrice(i)
	return ruo
}

// SetOrganizationId sets the "organizationId" field.
func (ruo *ResourceUpdateOne) SetOrganizationId(i int) *ResourceUpdateOne {
	ruo.mutation.SetOrganizationId(i)
	return ruo
}

// SetQuantityAvailable sets the "quantityAvailable" field.
func (ruo *ResourceUpdateOne) SetQuantityAvailable(i int) *ResourceUpdateOne {
	ruo.mutation.ResetQuantityAvailable()
	ruo.mutation.SetQuantityAvailable(i)
	return ruo
}

// SetNillableQuantityAvailable sets the "quantityAvailable" field if the given value is not nil.
func (ruo *ResourceUpdateOne) SetNillableQuantityAvailable(i *int) *ResourceUpdateOne {
	if i != nil {
		ruo.SetQuantityAvailable(*i)
	}
	return ruo
}

// AddQuantityAvailable adds i to the "quantityAvailable" field.
func (ruo *ResourceUpdateOne) AddQuantityAvailable(i int) *ResourceUpdateOne {
	ruo.mutation.AddQuantityAvailable(i)
	return ruo
}

// ClearQuantityAvailable clears the value of the "quantityAvailable" field.
func (ruo *ResourceUpdateOne) ClearQuantityAvailable() *ResourceUpdateOne {
	ruo.mutation.ClearQuantityAvailable()
	return ruo
}

// AddSlotIDs adds the "slots" edge to the Slot entity by IDs.
func (ruo *ResourceUpdateOne) AddSlotIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddSlotIDs(ids...)
	return ruo
}

// AddSlots adds the "slots" edges to the Slot entity.
func (ruo *ResourceUpdateOne) AddSlots(s ...*Slot) *ResourceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddSlotIDs(ids...)
}

// AddBookingIDs adds the "bookings" edge to the Booking entity by IDs.
func (ruo *ResourceUpdateOne) AddBookingIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddBookingIDs(ids...)
	return ruo
}

// AddBookings adds the "bookings" edges to the Booking entity.
func (ruo *ResourceUpdateOne) AddBookings(b ...*Booking) *ResourceUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.AddBookingIDs(ids...)
}

// AddUnavailabilityIDs adds the "unavailabilities" edge to the Unavailability entity by IDs.
func (ruo *ResourceUpdateOne) AddUnavailabilityIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.AddUnavailabilityIDs(ids...)
	return ruo
}

// AddUnavailabilities adds the "unavailabilities" edges to the Unavailability entity.
func (ruo *ResourceUpdateOne) AddUnavailabilities(u ...*Unavailability) *ResourceUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.AddUnavailabilityIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (ruo *ResourceUpdateOne) SetOrganizationID(id int) *ResourceUpdateOne {
	ruo.mutation.SetOrganizationID(id)
	return ruo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (ruo *ResourceUpdateOne) SetOrganization(o *Organization) *ResourceUpdateOne {
	return ruo.SetOrganizationID(o.ID)
}

// Mutation returns the ResourceMutation object of the builder.
func (ruo *ResourceUpdateOne) Mutation() *ResourceMutation {
	return ruo.mutation
}

// ClearSlots clears all "slots" edges to the Slot entity.
func (ruo *ResourceUpdateOne) ClearSlots() *ResourceUpdateOne {
	ruo.mutation.ClearSlots()
	return ruo
}

// RemoveSlotIDs removes the "slots" edge to Slot entities by IDs.
func (ruo *ResourceUpdateOne) RemoveSlotIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemoveSlotIDs(ids...)
	return ruo
}

// RemoveSlots removes "slots" edges to Slot entities.
func (ruo *ResourceUpdateOne) RemoveSlots(s ...*Slot) *ResourceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveSlotIDs(ids...)
}

// ClearBookings clears all "bookings" edges to the Booking entity.
func (ruo *ResourceUpdateOne) ClearBookings() *ResourceUpdateOne {
	ruo.mutation.ClearBookings()
	return ruo
}

// RemoveBookingIDs removes the "bookings" edge to Booking entities by IDs.
func (ruo *ResourceUpdateOne) RemoveBookingIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemoveBookingIDs(ids...)
	return ruo
}

// RemoveBookings removes "bookings" edges to Booking entities.
func (ruo *ResourceUpdateOne) RemoveBookings(b ...*Booking) *ResourceUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return ruo.RemoveBookingIDs(ids...)
}

// ClearUnavailabilities clears all "unavailabilities" edges to the Unavailability entity.
func (ruo *ResourceUpdateOne) ClearUnavailabilities() *ResourceUpdateOne {
	ruo.mutation.ClearUnavailabilities()
	return ruo
}

// RemoveUnavailabilityIDs removes the "unavailabilities" edge to Unavailability entities by IDs.
func (ruo *ResourceUpdateOne) RemoveUnavailabilityIDs(ids ...int) *ResourceUpdateOne {
	ruo.mutation.RemoveUnavailabilityIDs(ids...)
	return ruo
}

// RemoveUnavailabilities removes "unavailabilities" edges to Unavailability entities.
func (ruo *ResourceUpdateOne) RemoveUnavailabilities(u ...*Unavailability) *ResourceUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ruo.RemoveUnavailabilityIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (ruo *ResourceUpdateOne) ClearOrganization() *ResourceUpdateOne {
	ruo.mutation.ClearOrganization()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ResourceUpdateOne) Select(field string, fields ...string) *ResourceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Resource entity.
func (ruo *ResourceUpdateOne) Save(ctx context.Context) (*Resource, error) {
	var (
		err  error
		node *Resource
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ResourceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ResourceUpdateOne) SaveX(ctx context.Context) *Resource {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ResourceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ResourceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ResourceUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := resource.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ResourceUpdateOne) check() error {
	if _, ok := ruo.mutation.OrganizationID(); ruo.mutation.OrganizationCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"organization\"")
	}
	return nil
}

func (ruo *ResourceUpdateOne) sqlSave(ctx context.Context) (_node *Resource, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   resource.Table,
			Columns: resource.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: resource.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Resource.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resource.FieldID)
		for _, f := range fields {
			if !resource.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != resource.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: resource.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldName,
		})
	}
	if value, ok := ruo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldDescription,
		})
	}
	if value, ok := ruo.mutation.Timezone(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldTimezone,
		})
	}
	if value, ok := ruo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: resource.FieldPassword,
		})
	}
	if value, ok := ruo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldPrice,
		})
	}
	if value, ok := ruo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldPrice,
		})
	}
	if value, ok := ruo.mutation.BookingPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldBookingPrice,
		})
	}
	if value, ok := ruo.mutation.AddedBookingPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldBookingPrice,
		})
	}
	if value, ok := ruo.mutation.QuantityAvailable(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if value, ok := ruo.mutation.AddedQuantityAvailable(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if ruo.mutation.QuantityAvailableCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: resource.FieldQuantityAvailable,
		})
	}
	if ruo.mutation.SlotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedSlotsIDs(); len(nodes) > 0 && !ruo.mutation.SlotsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.SlotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.SlotsTable,
			Columns: []string{resource.SlotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: slot.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedBookingsIDs(); len(nodes) > 0 && !ruo.mutation.BookingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.BookingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.BookingsTable,
			Columns: []string{resource.BookingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: booking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UnavailabilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUnavailabilitiesIDs(); len(nodes) > 0 && !ruo.mutation.UnavailabilitiesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UnavailabilitiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   resource.UnavailabilitiesTable,
			Columns: []string{resource.UnavailabilitiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unavailability.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.OrganizationTable,
			Columns: []string{resource.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   resource.OrganizationTable,
			Columns: []string{resource.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: organization.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Resource{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{resource.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
