// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"magnum/ent/formsubmissiondata"
	"magnum/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FormSubmissionDataQuery is the builder for querying FormSubmissionData entities.
type FormSubmissionDataQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.FormSubmissionData
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FormSubmissionDataQuery builder.
func (fsdq *FormSubmissionDataQuery) Where(ps ...predicate.FormSubmissionData) *FormSubmissionDataQuery {
	fsdq.predicates = append(fsdq.predicates, ps...)
	return fsdq
}

// Limit the number of records to be returned by this query.
func (fsdq *FormSubmissionDataQuery) Limit(limit int) *FormSubmissionDataQuery {
	fsdq.ctx.Limit = &limit
	return fsdq
}

// Offset to start from.
func (fsdq *FormSubmissionDataQuery) Offset(offset int) *FormSubmissionDataQuery {
	fsdq.ctx.Offset = &offset
	return fsdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fsdq *FormSubmissionDataQuery) Unique(unique bool) *FormSubmissionDataQuery {
	fsdq.ctx.Unique = &unique
	return fsdq
}

// Order specifies how the records should be ordered.
func (fsdq *FormSubmissionDataQuery) Order(o ...OrderFunc) *FormSubmissionDataQuery {
	fsdq.order = append(fsdq.order, o...)
	return fsdq
}

// First returns the first FormSubmissionData entity from the query.
// Returns a *NotFoundError when no FormSubmissionData was found.
func (fsdq *FormSubmissionDataQuery) First(ctx context.Context) (*FormSubmissionData, error) {
	nodes, err := fsdq.Limit(1).All(setContextOp(ctx, fsdq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{formsubmissiondata.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) FirstX(ctx context.Context) *FormSubmissionData {
	node, err := fsdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FormSubmissionData ID from the query.
// Returns a *NotFoundError when no FormSubmissionData ID was found.
func (fsdq *FormSubmissionDataQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fsdq.Limit(1).IDs(setContextOp(ctx, fsdq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{formsubmissiondata.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) FirstIDX(ctx context.Context) string {
	id, err := fsdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FormSubmissionData entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FormSubmissionData entity is found.
// Returns a *NotFoundError when no FormSubmissionData entities are found.
func (fsdq *FormSubmissionDataQuery) Only(ctx context.Context) (*FormSubmissionData, error) {
	nodes, err := fsdq.Limit(2).All(setContextOp(ctx, fsdq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{formsubmissiondata.Label}
	default:
		return nil, &NotSingularError{formsubmissiondata.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) OnlyX(ctx context.Context) *FormSubmissionData {
	node, err := fsdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FormSubmissionData ID in the query.
// Returns a *NotSingularError when more than one FormSubmissionData ID is found.
// Returns a *NotFoundError when no entities are found.
func (fsdq *FormSubmissionDataQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fsdq.Limit(2).IDs(setContextOp(ctx, fsdq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{formsubmissiondata.Label}
	default:
		err = &NotSingularError{formsubmissiondata.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) OnlyIDX(ctx context.Context) string {
	id, err := fsdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FormSubmissionDataSlice.
func (fsdq *FormSubmissionDataQuery) All(ctx context.Context) ([]*FormSubmissionData, error) {
	ctx = setContextOp(ctx, fsdq.ctx, "All")
	if err := fsdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FormSubmissionData, *FormSubmissionDataQuery]()
	return withInterceptors[[]*FormSubmissionData](ctx, fsdq, qr, fsdq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) AllX(ctx context.Context) []*FormSubmissionData {
	nodes, err := fsdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FormSubmissionData IDs.
func (fsdq *FormSubmissionDataQuery) IDs(ctx context.Context) (ids []string, err error) {
	if fsdq.ctx.Unique == nil && fsdq.path != nil {
		fsdq.Unique(true)
	}
	ctx = setContextOp(ctx, fsdq.ctx, "IDs")
	if err = fsdq.Select(formsubmissiondata.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) IDsX(ctx context.Context) []string {
	ids, err := fsdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fsdq *FormSubmissionDataQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fsdq.ctx, "Count")
	if err := fsdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fsdq, querierCount[*FormSubmissionDataQuery](), fsdq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) CountX(ctx context.Context) int {
	count, err := fsdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fsdq *FormSubmissionDataQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fsdq.ctx, "Exist")
	switch _, err := fsdq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fsdq *FormSubmissionDataQuery) ExistX(ctx context.Context) bool {
	exist, err := fsdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FormSubmissionDataQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fsdq *FormSubmissionDataQuery) Clone() *FormSubmissionDataQuery {
	if fsdq == nil {
		return nil
	}
	return &FormSubmissionDataQuery{
		config:     fsdq.config,
		ctx:        fsdq.ctx.Clone(),
		order:      append([]OrderFunc{}, fsdq.order...),
		inters:     append([]Interceptor{}, fsdq.inters...),
		predicates: append([]predicate.FormSubmissionData{}, fsdq.predicates...),
		// clone intermediate query.
		sql:  fsdq.sql.Clone(),
		path: fsdq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		SubmissionID string `json:"submission_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FormSubmissionData.Query().
//		GroupBy(formsubmissiondata.FieldSubmissionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fsdq *FormSubmissionDataQuery) GroupBy(field string, fields ...string) *FormSubmissionDataGroupBy {
	fsdq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FormSubmissionDataGroupBy{build: fsdq}
	grbuild.flds = &fsdq.ctx.Fields
	grbuild.label = formsubmissiondata.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		SubmissionID string `json:"submission_id,omitempty"`
//	}
//
//	client.FormSubmissionData.Query().
//		Select(formsubmissiondata.FieldSubmissionID).
//		Scan(ctx, &v)
func (fsdq *FormSubmissionDataQuery) Select(fields ...string) *FormSubmissionDataSelect {
	fsdq.ctx.Fields = append(fsdq.ctx.Fields, fields...)
	sbuild := &FormSubmissionDataSelect{FormSubmissionDataQuery: fsdq}
	sbuild.label = formsubmissiondata.Label
	sbuild.flds, sbuild.scan = &fsdq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FormSubmissionDataSelect configured with the given aggregations.
func (fsdq *FormSubmissionDataQuery) Aggregate(fns ...AggregateFunc) *FormSubmissionDataSelect {
	return fsdq.Select().Aggregate(fns...)
}

func (fsdq *FormSubmissionDataQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fsdq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fsdq); err != nil {
				return err
			}
		}
	}
	for _, f := range fsdq.ctx.Fields {
		if !formsubmissiondata.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fsdq.path != nil {
		prev, err := fsdq.path(ctx)
		if err != nil {
			return err
		}
		fsdq.sql = prev
	}
	return nil
}

func (fsdq *FormSubmissionDataQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FormSubmissionData, error) {
	var (
		nodes = []*FormSubmissionData{}
		_spec = fsdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FormSubmissionData).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FormSubmissionData{config: fsdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fsdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (fsdq *FormSubmissionDataQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fsdq.querySpec()
	_spec.Node.Columns = fsdq.ctx.Fields
	if len(fsdq.ctx.Fields) > 0 {
		_spec.Unique = fsdq.ctx.Unique != nil && *fsdq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fsdq.driver, _spec)
}

func (fsdq *FormSubmissionDataQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(formsubmissiondata.Table, formsubmissiondata.Columns, sqlgraph.NewFieldSpec(formsubmissiondata.FieldID, field.TypeString))
	_spec.From = fsdq.sql
	if unique := fsdq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fsdq.path != nil {
		_spec.Unique = true
	}
	if fields := fsdq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, formsubmissiondata.FieldID)
		for i := range fields {
			if fields[i] != formsubmissiondata.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fsdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fsdq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fsdq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fsdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fsdq *FormSubmissionDataQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fsdq.driver.Dialect())
	t1 := builder.Table(formsubmissiondata.Table)
	columns := fsdq.ctx.Fields
	if len(columns) == 0 {
		columns = formsubmissiondata.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fsdq.sql != nil {
		selector = fsdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fsdq.ctx.Unique != nil && *fsdq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fsdq.predicates {
		p(selector)
	}
	for _, p := range fsdq.order {
		p(selector)
	}
	if offset := fsdq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fsdq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FormSubmissionDataGroupBy is the group-by builder for FormSubmissionData entities.
type FormSubmissionDataGroupBy struct {
	selector
	build *FormSubmissionDataQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fsdgb *FormSubmissionDataGroupBy) Aggregate(fns ...AggregateFunc) *FormSubmissionDataGroupBy {
	fsdgb.fns = append(fsdgb.fns, fns...)
	return fsdgb
}

// Scan applies the selector query and scans the result into the given value.
func (fsdgb *FormSubmissionDataGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fsdgb.build.ctx, "GroupBy")
	if err := fsdgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FormSubmissionDataQuery, *FormSubmissionDataGroupBy](ctx, fsdgb.build, fsdgb, fsdgb.build.inters, v)
}

func (fsdgb *FormSubmissionDataGroupBy) sqlScan(ctx context.Context, root *FormSubmissionDataQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fsdgb.fns))
	for _, fn := range fsdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fsdgb.flds)+len(fsdgb.fns))
		for _, f := range *fsdgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fsdgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fsdgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FormSubmissionDataSelect is the builder for selecting fields of FormSubmissionData entities.
type FormSubmissionDataSelect struct {
	*FormSubmissionDataQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fsds *FormSubmissionDataSelect) Aggregate(fns ...AggregateFunc) *FormSubmissionDataSelect {
	fsds.fns = append(fsds.fns, fns...)
	return fsds
}

// Scan applies the selector query and scans the result into the given value.
func (fsds *FormSubmissionDataSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fsds.ctx, "Select")
	if err := fsds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FormSubmissionDataQuery, *FormSubmissionDataSelect](ctx, fsds.FormSubmissionDataQuery, fsds, fsds.inters, v)
}

func (fsds *FormSubmissionDataSelect) sqlScan(ctx context.Context, root *FormSubmissionDataQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fsds.fns))
	for _, fn := range fsds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fsds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fsds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
