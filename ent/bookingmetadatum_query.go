// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/openmesh/booking/ent/booking"
	"github.com/openmesh/booking/ent/bookingmetadatum"
	"github.com/openmesh/booking/ent/predicate"
)

// BookingMetadatumQuery is the builder for querying BookingMetadatum entities.
type BookingMetadatumQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BookingMetadatum
	// eager-loading edges.
	withBooking *BookingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BookingMetadatumQuery builder.
func (bmq *BookingMetadatumQuery) Where(ps ...predicate.BookingMetadatum) *BookingMetadatumQuery {
	bmq.predicates = append(bmq.predicates, ps...)
	return bmq
}

// Limit adds a limit step to the query.
func (bmq *BookingMetadatumQuery) Limit(limit int) *BookingMetadatumQuery {
	bmq.limit = &limit
	return bmq
}

// Offset adds an offset step to the query.
func (bmq *BookingMetadatumQuery) Offset(offset int) *BookingMetadatumQuery {
	bmq.offset = &offset
	return bmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bmq *BookingMetadatumQuery) Unique(unique bool) *BookingMetadatumQuery {
	bmq.unique = &unique
	return bmq
}

// Order adds an order step to the query.
func (bmq *BookingMetadatumQuery) Order(o ...OrderFunc) *BookingMetadatumQuery {
	bmq.order = append(bmq.order, o...)
	return bmq
}

// QueryBooking chains the current query on the "booking" edge.
func (bmq *BookingMetadatumQuery) QueryBooking() *BookingQuery {
	query := &BookingQuery{config: bmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bookingmetadatum.Table, bookingmetadatum.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bookingmetadatum.BookingTable, bookingmetadatum.BookingColumn),
		)
		fromU = sqlgraph.SetNeighbors(bmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BookingMetadatum entity from the query.
// Returns a *NotFoundError when no BookingMetadatum was found.
func (bmq *BookingMetadatumQuery) First(ctx context.Context) (*BookingMetadatum, error) {
	nodes, err := bmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bookingmetadatum.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) FirstX(ctx context.Context) *BookingMetadatum {
	node, err := bmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BookingMetadatum ID from the query.
// Returns a *NotFoundError when no BookingMetadatum ID was found.
func (bmq *BookingMetadatumQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bookingmetadatum.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) FirstIDX(ctx context.Context) int {
	id, err := bmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BookingMetadatum entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one BookingMetadatum entity is not found.
// Returns a *NotFoundError when no BookingMetadatum entities are found.
func (bmq *BookingMetadatumQuery) Only(ctx context.Context) (*BookingMetadatum, error) {
	nodes, err := bmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bookingmetadatum.Label}
	default:
		return nil, &NotSingularError{bookingmetadatum.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) OnlyX(ctx context.Context) *BookingMetadatum {
	node, err := bmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BookingMetadatum ID in the query.
// Returns a *NotSingularError when exactly one BookingMetadatum ID is not found.
// Returns a *NotFoundError when no entities are found.
func (bmq *BookingMetadatumQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = &NotSingularError{bookingmetadatum.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) OnlyIDX(ctx context.Context) int {
	id, err := bmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BookingMetadata.
func (bmq *BookingMetadatumQuery) All(ctx context.Context) ([]*BookingMetadatum, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return bmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) AllX(ctx context.Context) []*BookingMetadatum {
	nodes, err := bmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BookingMetadatum IDs.
func (bmq *BookingMetadatumQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := bmq.Select(bookingmetadatum.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) IDsX(ctx context.Context) []int {
	ids, err := bmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bmq *BookingMetadatumQuery) Count(ctx context.Context) (int, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return bmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) CountX(ctx context.Context) int {
	count, err := bmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bmq *BookingMetadatumQuery) Exist(ctx context.Context) (bool, error) {
	if err := bmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return bmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (bmq *BookingMetadatumQuery) ExistX(ctx context.Context) bool {
	exist, err := bmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BookingMetadatumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bmq *BookingMetadatumQuery) Clone() *BookingMetadatumQuery {
	if bmq == nil {
		return nil
	}
	return &BookingMetadatumQuery{
		config:      bmq.config,
		limit:       bmq.limit,
		offset:      bmq.offset,
		order:       append([]OrderFunc{}, bmq.order...),
		predicates:  append([]predicate.BookingMetadatum{}, bmq.predicates...),
		withBooking: bmq.withBooking.Clone(),
		// clone intermediate query.
		sql:  bmq.sql.Clone(),
		path: bmq.path,
	}
}

// WithBooking tells the query-builder to eager-load the nodes that are connected to
// the "booking" edge. The optional arguments are used to configure the query builder of the edge.
func (bmq *BookingMetadatumQuery) WithBooking(opts ...func(*BookingQuery)) *BookingMetadatumQuery {
	query := &BookingQuery{config: bmq.config}
	for _, opt := range opts {
		opt(query)
	}
	bmq.withBooking = query
	return bmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BookingMetadatum.Query().
//		GroupBy(bookingmetadatum.FieldKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (bmq *BookingMetadatumQuery) GroupBy(field string, fields ...string) *BookingMetadatumGroupBy {
	group := &BookingMetadatumGroupBy{config: bmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := bmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return bmq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Key string `json:"key,omitempty"`
//	}
//
//	client.BookingMetadatum.Query().
//		Select(bookingmetadatum.FieldKey).
//		Scan(ctx, &v)
//
func (bmq *BookingMetadatumQuery) Select(fields ...string) *BookingMetadatumSelect {
	bmq.fields = append(bmq.fields, fields...)
	return &BookingMetadatumSelect{BookingMetadatumQuery: bmq}
}

func (bmq *BookingMetadatumQuery) prepareQuery(ctx context.Context) error {
	for _, f := range bmq.fields {
		if !bookingmetadatum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bmq.path != nil {
		prev, err := bmq.path(ctx)
		if err != nil {
			return err
		}
		bmq.sql = prev
	}
	if bookingmetadatum.Policy == nil {
		return errors.New("ent: uninitialized bookingmetadatum.Policy (forgotten import ent/runtime?)")
	}
	if err := bookingmetadatum.Policy.EvalQuery(ctx, bmq); err != nil {
		return err
	}
	return nil
}

func (bmq *BookingMetadatumQuery) sqlAll(ctx context.Context) ([]*BookingMetadatum, error) {
	var (
		nodes       = []*BookingMetadatum{}
		_spec       = bmq.querySpec()
		loadedTypes = [1]bool{
			bmq.withBooking != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &BookingMetadatum{config: bmq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, bmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := bmq.withBooking; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BookingMetadatum)
		for i := range nodes {
			fk := nodes[i].BookingId
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(booking.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "bookingId" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Booking = n
			}
		}
	}

	return nodes, nil
}

func (bmq *BookingMetadatumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bmq.querySpec()
	return sqlgraph.CountNodes(ctx, bmq.driver, _spec)
}

func (bmq *BookingMetadatumQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := bmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (bmq *BookingMetadatumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bookingmetadatum.Table,
			Columns: bookingmetadatum.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bookingmetadatum.FieldID,
			},
		},
		From:   bmq.sql,
		Unique: true,
	}
	if unique := bmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := bmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bookingmetadatum.FieldID)
		for i := range fields {
			if fields[i] != bookingmetadatum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bmq *BookingMetadatumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bmq.driver.Dialect())
	t1 := builder.Table(bookingmetadatum.Table)
	columns := bmq.fields
	if len(columns) == 0 {
		columns = bookingmetadatum.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bmq.sql != nil {
		selector = bmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range bmq.predicates {
		p(selector)
	}
	for _, p := range bmq.order {
		p(selector)
	}
	if offset := bmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BookingMetadatumGroupBy is the group-by builder for BookingMetadatum entities.
type BookingMetadatumGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bmgb *BookingMetadatumGroupBy) Aggregate(fns ...AggregateFunc) *BookingMetadatumGroupBy {
	bmgb.fns = append(bmgb.fns, fns...)
	return bmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (bmgb *BookingMetadatumGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bmgb.path(ctx)
	if err != nil {
		return err
	}
	bmgb.sql = query
	return bmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := bmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(bmgb.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := bmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) StringsX(ctx context.Context) []string {
	v, err := bmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) StringX(ctx context.Context) string {
	v, err := bmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(bmgb.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := bmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) IntsX(ctx context.Context) []int {
	v, err := bmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) IntX(ctx context.Context) int {
	v, err := bmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(bmgb.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := bmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := bmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) Float64X(ctx context.Context) float64 {
	v, err := bmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(bmgb.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := bmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := bmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (bmgb *BookingMetadatumGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bmgb *BookingMetadatumGroupBy) BoolX(ctx context.Context) bool {
	v, err := bmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bmgb *BookingMetadatumGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bmgb.fields {
		if !bookingmetadatum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bmgb *BookingMetadatumGroupBy) sqlQuery() *sql.Selector {
	selector := bmgb.sql.Select()
	aggregation := make([]string, 0, len(bmgb.fns))
	for _, fn := range bmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bmgb.fields)+len(bmgb.fns))
		for _, f := range bmgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bmgb.fields...)...)
}

// BookingMetadatumSelect is the builder for selecting fields of BookingMetadatum entities.
type BookingMetadatumSelect struct {
	*BookingMetadatumQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bms *BookingMetadatumSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bms.prepareQuery(ctx); err != nil {
		return err
	}
	bms.sql = bms.BookingMetadatumQuery.sqlQuery(ctx)
	return bms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (bms *BookingMetadatumSelect) ScanX(ctx context.Context, v interface{}) {
	if err := bms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Strings(ctx context.Context) ([]string, error) {
	if len(bms.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := bms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (bms *BookingMetadatumSelect) StringsX(ctx context.Context) []string {
	v, err := bms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = bms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (bms *BookingMetadatumSelect) StringX(ctx context.Context) string {
	v, err := bms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Ints(ctx context.Context) ([]int, error) {
	if len(bms.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := bms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (bms *BookingMetadatumSelect) IntsX(ctx context.Context) []int {
	v, err := bms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = bms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (bms *BookingMetadatumSelect) IntX(ctx context.Context) int {
	v, err := bms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(bms.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := bms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (bms *BookingMetadatumSelect) Float64sX(ctx context.Context) []float64 {
	v, err := bms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = bms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (bms *BookingMetadatumSelect) Float64X(ctx context.Context) float64 {
	v, err := bms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(bms.fields) > 1 {
		return nil, errors.New("ent: BookingMetadatumSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := bms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (bms *BookingMetadatumSelect) BoolsX(ctx context.Context) []bool {
	v, err := bms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (bms *BookingMetadatumSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = bms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{bookingmetadatum.Label}
	default:
		err = fmt.Errorf("ent: BookingMetadatumSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (bms *BookingMetadatumSelect) BoolX(ctx context.Context) bool {
	v, err := bms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bms *BookingMetadatumSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bms.sql.Query()
	if err := bms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
