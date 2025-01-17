// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/auth"
	"github.com/openmesh/booking/ent/organization"
	"github.com/openmesh/booking/ent/predicate"
	"github.com/openmesh/booking/ent/token"
	"github.com/openmesh/booking/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updatedAt" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetOrganizationId sets the "organizationId" field.
func (uu *UserUpdate) SetOrganizationId(i int) *UserUpdate {
	uu.mutation.SetOrganizationId(i)
	return uu
}

// SetNillableOrganizationId sets the "organizationId" field if the given value is not nil.
func (uu *UserUpdate) SetNillableOrganizationId(i *int) *UserUpdate {
	if i != nil {
		uu.SetOrganizationId(*i)
	}
	return uu
}

// ClearOrganizationId clears the value of the "organizationId" field.
func (uu *UserUpdate) ClearOrganizationId() *UserUpdate {
	uu.mutation.ClearOrganizationId()
	return uu
}

// AddAuthIDs adds the "auths" edge to the Auth entity by IDs.
func (uu *UserUpdate) AddAuthIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAuthIDs(ids...)
	return uu
}

// AddAuths adds the "auths" edges to the Auth entity.
func (uu *UserUpdate) AddAuths(a ...*Auth) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.AddAuthIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uu *UserUpdate) AddTokenIDs(ids ...string) *UserUpdate {
	uu.mutation.AddTokenIDs(ids...)
	return uu
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uu *UserUpdate) AddTokens(t ...*Token) *UserUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTokenIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (uu *UserUpdate) SetOrganizationID(id int) *UserUpdate {
	uu.mutation.SetOrganizationID(id)
	return uu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableOrganizationID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetOrganizationID(*id)
	}
	return uu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (uu *UserUpdate) SetOrganization(o *Organization) *UserUpdate {
	return uu.SetOrganizationID(o.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearAuths clears all "auths" edges to the Auth entity.
func (uu *UserUpdate) ClearAuths() *UserUpdate {
	uu.mutation.ClearAuths()
	return uu
}

// RemoveAuthIDs removes the "auths" edge to Auth entities by IDs.
func (uu *UserUpdate) RemoveAuthIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAuthIDs(ids...)
	return uu
}

// RemoveAuths removes "auths" edges to Auth entities.
func (uu *UserUpdate) RemoveAuths(a ...*Auth) *UserUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uu.RemoveAuthIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uu *UserUpdate) ClearTokens() *UserUpdate {
	uu.mutation.ClearTokens()
	return uu
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uu *UserUpdate) RemoveTokenIDs(ids ...string) *UserUpdate {
	uu.mutation.RemoveTokenIDs(ids...)
	return uu
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uu *UserUpdate) RemoveTokens(t ...*Token) *UserUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTokenIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (uu *UserUpdate) ClearOrganization() *UserUpdate {
	uu.mutation.ClearOrganization()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
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
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uu.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAuthsIDs(); len(nodes) > 0 && !uu.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AuthsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
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
	if nodes := uu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updatedAt" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetOrganizationId sets the "organizationId" field.
func (uuo *UserUpdateOne) SetOrganizationId(i int) *UserUpdateOne {
	uuo.mutation.SetOrganizationId(i)
	return uuo
}

// SetNillableOrganizationId sets the "organizationId" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOrganizationId(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetOrganizationId(*i)
	}
	return uuo
}

// ClearOrganizationId clears the value of the "organizationId" field.
func (uuo *UserUpdateOne) ClearOrganizationId() *UserUpdateOne {
	uuo.mutation.ClearOrganizationId()
	return uuo
}

// AddAuthIDs adds the "auths" edge to the Auth entity by IDs.
func (uuo *UserUpdateOne) AddAuthIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAuthIDs(ids...)
	return uuo
}

// AddAuths adds the "auths" edges to the Auth entity.
func (uuo *UserUpdateOne) AddAuths(a ...*Auth) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.AddAuthIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uuo *UserUpdateOne) AddTokenIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddTokenIDs(ids...)
	return uuo
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) AddTokens(t ...*Token) *UserUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTokenIDs(ids...)
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (uuo *UserUpdateOne) SetOrganizationID(id int) *UserUpdateOne {
	uuo.mutation.SetOrganizationID(id)
	return uuo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableOrganizationID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetOrganizationID(*id)
	}
	return uuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (uuo *UserUpdateOne) SetOrganization(o *Organization) *UserUpdateOne {
	return uuo.SetOrganizationID(o.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearAuths clears all "auths" edges to the Auth entity.
func (uuo *UserUpdateOne) ClearAuths() *UserUpdateOne {
	uuo.mutation.ClearAuths()
	return uuo
}

// RemoveAuthIDs removes the "auths" edge to Auth entities by IDs.
func (uuo *UserUpdateOne) RemoveAuthIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAuthIDs(ids...)
	return uuo
}

// RemoveAuths removes "auths" edges to Auth entities.
func (uuo *UserUpdateOne) RemoveAuths(a ...*Auth) *UserUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return uuo.RemoveAuthIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) ClearTokens() *UserUpdateOne {
	uuo.mutation.ClearTokens()
	return uuo
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uuo *UserUpdateOne) RemoveTokenIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemoveTokenIDs(ids...)
	return uuo
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uuo *UserUpdateOne) RemoveTokens(t ...*Token) *UserUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTokenIDs(ids...)
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (uuo *UserUpdateOne) ClearOrganization() *UserUpdateOne {
	uuo.mutation.ClearOrganization()
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing User.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
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
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldEmail,
		})
	}
	if uuo.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAuthsIDs(); len(nodes) > 0 && !uuo.mutation.AuthsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AuthsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AuthsTable,
			Columns: []string{user.AuthsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: auth.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: token.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
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
	if nodes := uuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.OrganizationTable,
			Columns: []string{user.OrganizationColumn},
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
