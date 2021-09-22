// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/bookingmetadatum"
	"github.com/openmesh/booking/ent/organization"
	"github.com/openmesh/booking/ent/resource"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (bc *BookingCreate) SetCreatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bc *BookingCreate) SetNillableCreatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updatedAt" field.
func (bc *BookingCreate) SetUpdatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bc *BookingCreate) SetNillableUpdatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BookingCreate) SetStatus(s string) *BookingCreate {
	bc.mutation.SetStatus(s)
	return bc
}

// SetStartTime sets the "startTime" field.
func (bc *BookingCreate) SetStartTime(t time.Time) *BookingCreate {
	bc.mutation.SetStartTime(t)
	return bc
}

// SetEndTime sets the "endTime" field.
func (bc *BookingCreate) SetEndTime(t time.Time) *BookingCreate {
	bc.mutation.SetEndTime(t)
	return bc
}

// SetResourceId sets the "resourceId" field.
func (bc *BookingCreate) SetResourceId(i int) *BookingCreate {
	bc.mutation.SetResourceId(i)
	return bc
}

// SetOrganizationId sets the "organizationId" field.
func (bc *BookingCreate) SetOrganizationId(i int) *BookingCreate {
	bc.mutation.SetOrganizationId(i)
	return bc
}

// AddMetadatumIDs adds the "metadata" edge to the BookingMetadatum entity by IDs.
func (bc *BookingCreate) AddMetadatumIDs(ids ...int) *BookingCreate {
	bc.mutation.AddMetadatumIDs(ids...)
	return bc
}

// AddMetadata adds the "metadata" edges to the BookingMetadatum entity.
func (bc *BookingCreate) AddMetadata(b ...*BookingMetadatum) *BookingCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddMetadatumIDs(ids...)
}

// SetResourceID sets the "resource" edge to the Resource entity by ID.
func (bc *BookingCreate) SetResourceID(id int) *BookingCreate {
	bc.mutation.SetResourceID(id)
	return bc
}

// SetResource sets the "resource" edge to the Resource entity.
func (bc *BookingCreate) SetResource(r *Resource) *BookingCreate {
	return bc.SetResourceID(r.ID)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (bc *BookingCreate) SetOrganizationID(id int) *BookingCreate {
	bc.mutation.SetOrganizationID(id)
	return bc
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (bc *BookingCreate) SetOrganization(o *Organization) *BookingCreate {
	return bc.SetOrganizationID(o.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	var (
		err  error
		node *Booking
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BookingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookingCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookingCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookingCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := booking.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := booking.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookingCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "createdAt"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "updatedAt"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if _, ok := bc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "startTime", err: errors.New(`ent: missing required field "startTime"`)}
	}
	if _, ok := bc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "endTime", err: errors.New(`ent: missing required field "endTime"`)}
	}
	if _, ok := bc.mutation.ResourceId(); !ok {
		return &ValidationError{Name: "resourceId", err: errors.New(`ent: missing required field "resourceId"`)}
	}
	if _, ok := bc.mutation.OrganizationId(); !ok {
		return &ValidationError{Name: "organizationId", err: errors.New(`ent: missing required field "organizationId"`)}
	}
	if _, ok := bc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource", err: errors.New("ent: missing required edge \"resource\"")}
	}
	if _, ok := bc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization", err: errors.New("ent: missing required edge \"organization\"")}
	}
	return nil
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		_node = &Booking{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: booking.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: booking.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: booking.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := bc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := bc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: booking.FieldEndTime,
		})
		_node.EndTime = value
	}
	if nodes := bc.mutation.MetadataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   booking.MetadataTable,
			Columns: []string{booking.MetadataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bookingmetadatum.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.ResourceIDs(); len(nodes) > 0 {
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
		_node.ResourceId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.OrganizationTable,
			Columns: []string{booking.OrganizationColumn},
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
		_node.OrganizationId = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookingCreateBulk is the builder for creating many Booking entities in bulk.
type BookingCreateBulk struct {
	config
	builders []*BookingCreate
}

// Save creates the Booking entities in the database.
func (bcb *BookingCreateBulk) Save(ctx context.Context) ([]*Booking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Booking, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookingCreateBulk) SaveX(ctx context.Context) []*Booking {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookingCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookingCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
