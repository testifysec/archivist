// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/in-toto/archivista/ent/attestationcollection"
	"github.com/in-toto/archivista/ent/attestationpolicy"
	"github.com/in-toto/archivista/ent/dsse"
	"github.com/in-toto/archivista/ent/predicate"
	"github.com/in-toto/archivista/ent/statement"
	"github.com/in-toto/archivista/ent/subject"
)

// StatementQuery is the builder for querying Statement entities.
type StatementQuery struct {
	config
	ctx                        *QueryContext
	order                      []statement.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.Statement
	withSubjects               *SubjectQuery
	withPolicy                 *AttestationPolicyQuery
	withAttestationCollections *AttestationCollectionQuery
	withDsse                   *DsseQuery
	modifiers                  []func(*sql.Selector)
	loadTotal                  []func(context.Context, []*Statement) error
	withNamedSubjects          map[string]*SubjectQuery
	withNamedDsse              map[string]*DsseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StatementQuery builder.
func (sq *StatementQuery) Where(ps ...predicate.Statement) *StatementQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *StatementQuery) Limit(limit int) *StatementQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *StatementQuery) Offset(offset int) *StatementQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *StatementQuery) Unique(unique bool) *StatementQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *StatementQuery) Order(o ...statement.OrderOption) *StatementQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QuerySubjects chains the current query on the "subjects" edge.
func (sq *StatementQuery) QuerySubjects() *SubjectQuery {
	query := (&SubjectClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statement.Table, statement.FieldID, selector),
			sqlgraph.To(subject.Table, subject.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, statement.SubjectsTable, statement.SubjectsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPolicy chains the current query on the "policy" edge.
func (sq *StatementQuery) QueryPolicy() *AttestationPolicyQuery {
	query := (&AttestationPolicyClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statement.Table, statement.FieldID, selector),
			sqlgraph.To(attestationpolicy.Table, attestationpolicy.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, statement.PolicyTable, statement.PolicyColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAttestationCollections chains the current query on the "attestation_collections" edge.
func (sq *StatementQuery) QueryAttestationCollections() *AttestationCollectionQuery {
	query := (&AttestationCollectionClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statement.Table, statement.FieldID, selector),
			sqlgraph.To(attestationcollection.Table, attestationcollection.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, statement.AttestationCollectionsTable, statement.AttestationCollectionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDsse chains the current query on the "dsse" edge.
func (sq *StatementQuery) QueryDsse() *DsseQuery {
	query := (&DsseClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(statement.Table, statement.FieldID, selector),
			sqlgraph.To(dsse.Table, dsse.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, statement.DsseTable, statement.DsseColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Statement entity from the query.
// Returns a *NotFoundError when no Statement was found.
func (sq *StatementQuery) First(ctx context.Context) (*Statement, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{statement.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *StatementQuery) FirstX(ctx context.Context) *Statement {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Statement ID from the query.
// Returns a *NotFoundError when no Statement ID was found.
func (sq *StatementQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{statement.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *StatementQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Statement entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Statement entity is found.
// Returns a *NotFoundError when no Statement entities are found.
func (sq *StatementQuery) Only(ctx context.Context) (*Statement, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{statement.Label}
	default:
		return nil, &NotSingularError{statement.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *StatementQuery) OnlyX(ctx context.Context) *Statement {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Statement ID in the query.
// Returns a *NotSingularError when more than one Statement ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *StatementQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{statement.Label}
	default:
		err = &NotSingularError{statement.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *StatementQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Statements.
func (sq *StatementQuery) All(ctx context.Context) ([]*Statement, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Statement, *StatementQuery]()
	return withInterceptors[[]*Statement](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *StatementQuery) AllX(ctx context.Context) []*Statement {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Statement IDs.
func (sq *StatementQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sq.ctx.Unique == nil && sq.path != nil {
		sq.Unique(true)
	}
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err = sq.Select(statement.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *StatementQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *StatementQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*StatementQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *StatementQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *StatementQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *StatementQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StatementQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *StatementQuery) Clone() *StatementQuery {
	if sq == nil {
		return nil
	}
	return &StatementQuery{
		config:                     sq.config,
		ctx:                        sq.ctx.Clone(),
		order:                      append([]statement.OrderOption{}, sq.order...),
		inters:                     append([]Interceptor{}, sq.inters...),
		predicates:                 append([]predicate.Statement{}, sq.predicates...),
		withSubjects:               sq.withSubjects.Clone(),
		withPolicy:                 sq.withPolicy.Clone(),
		withAttestationCollections: sq.withAttestationCollections.Clone(),
		withDsse:                   sq.withDsse.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithSubjects tells the query-builder to eager-load the nodes that are connected to
// the "subjects" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithSubjects(opts ...func(*SubjectQuery)) *StatementQuery {
	query := (&SubjectClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withSubjects = query
	return sq
}

// WithPolicy tells the query-builder to eager-load the nodes that are connected to
// the "policy" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithPolicy(opts ...func(*AttestationPolicyQuery)) *StatementQuery {
	query := (&AttestationPolicyClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withPolicy = query
	return sq
}

// WithAttestationCollections tells the query-builder to eager-load the nodes that are connected to
// the "attestation_collections" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithAttestationCollections(opts ...func(*AttestationCollectionQuery)) *StatementQuery {
	query := (&AttestationCollectionClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withAttestationCollections = query
	return sq
}

// WithDsse tells the query-builder to eager-load the nodes that are connected to
// the "dsse" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithDsse(opts ...func(*DsseQuery)) *StatementQuery {
	query := (&DsseClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withDsse = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Predicate string `json:"predicate,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Statement.Query().
//		GroupBy(statement.FieldPredicate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *StatementQuery) GroupBy(field string, fields ...string) *StatementGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StatementGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = statement.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Predicate string `json:"predicate,omitempty"`
//	}
//
//	client.Statement.Query().
//		Select(statement.FieldPredicate).
//		Scan(ctx, &v)
func (sq *StatementQuery) Select(fields ...string) *StatementSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &StatementSelect{StatementQuery: sq}
	sbuild.label = statement.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StatementSelect configured with the given aggregations.
func (sq *StatementQuery) Aggregate(fns ...AggregateFunc) *StatementSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *StatementQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !statement.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *StatementQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Statement, error) {
	var (
		nodes       = []*Statement{}
		_spec       = sq.querySpec()
		loadedTypes = [4]bool{
			sq.withSubjects != nil,
			sq.withPolicy != nil,
			sq.withAttestationCollections != nil,
			sq.withDsse != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Statement).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Statement{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withSubjects; query != nil {
		if err := sq.loadSubjects(ctx, query, nodes,
			func(n *Statement) { n.Edges.Subjects = []*Subject{} },
			func(n *Statement, e *Subject) { n.Edges.Subjects = append(n.Edges.Subjects, e) }); err != nil {
			return nil, err
		}
	}
	if query := sq.withPolicy; query != nil {
		if err := sq.loadPolicy(ctx, query, nodes, nil,
			func(n *Statement, e *AttestationPolicy) { n.Edges.Policy = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withAttestationCollections; query != nil {
		if err := sq.loadAttestationCollections(ctx, query, nodes, nil,
			func(n *Statement, e *AttestationCollection) { n.Edges.AttestationCollections = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withDsse; query != nil {
		if err := sq.loadDsse(ctx, query, nodes,
			func(n *Statement) { n.Edges.Dsse = []*Dsse{} },
			func(n *Statement, e *Dsse) { n.Edges.Dsse = append(n.Edges.Dsse, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range sq.withNamedSubjects {
		if err := sq.loadSubjects(ctx, query, nodes,
			func(n *Statement) { n.appendNamedSubjects(name) },
			func(n *Statement, e *Subject) { n.appendNamedSubjects(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range sq.withNamedDsse {
		if err := sq.loadDsse(ctx, query, nodes,
			func(n *Statement) { n.appendNamedDsse(name) },
			func(n *Statement, e *Dsse) { n.appendNamedDsse(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range sq.loadTotal {
		if err := sq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *StatementQuery) loadSubjects(ctx context.Context, query *SubjectQuery, nodes []*Statement, init func(*Statement), assign func(*Statement, *Subject)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Statement)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Subject(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(statement.SubjectsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.statement_subjects
		if fk == nil {
			return fmt.Errorf(`foreign-key "statement_subjects" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "statement_subjects" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *StatementQuery) loadPolicy(ctx context.Context, query *AttestationPolicyQuery, nodes []*Statement, init func(*Statement), assign func(*Statement, *AttestationPolicy)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Statement)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.AttestationPolicy(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(statement.PolicyColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.statement_policy
		if fk == nil {
			return fmt.Errorf(`foreign-key "statement_policy" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "statement_policy" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *StatementQuery) loadAttestationCollections(ctx context.Context, query *AttestationCollectionQuery, nodes []*Statement, init func(*Statement), assign func(*Statement, *AttestationCollection)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Statement)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.AttestationCollection(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(statement.AttestationCollectionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.statement_attestation_collections
		if fk == nil {
			return fmt.Errorf(`foreign-key "statement_attestation_collections" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "statement_attestation_collections" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (sq *StatementQuery) loadDsse(ctx context.Context, query *DsseQuery, nodes []*Statement, init func(*Statement), assign func(*Statement, *Dsse)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Statement)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Dsse(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(statement.DsseColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.dsse_statement
		if fk == nil {
			return fmt.Errorf(`foreign-key "dsse_statement" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "dsse_statement" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (sq *StatementQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	if len(sq.modifiers) > 0 {
		_spec.Modifiers = sq.modifiers
	}
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *StatementQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(statement.Table, statement.Columns, sqlgraph.NewFieldSpec(statement.FieldID, field.TypeInt))
	_spec.From = sq.sql
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sq.path != nil {
		_spec.Unique = true
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statement.FieldID)
		for i := range fields {
			if fields[i] != statement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *StatementQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(statement.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = statement.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedSubjects tells the query-builder to eager-load the nodes that are connected to the "subjects"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithNamedSubjects(name string, opts ...func(*SubjectQuery)) *StatementQuery {
	query := (&SubjectClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if sq.withNamedSubjects == nil {
		sq.withNamedSubjects = make(map[string]*SubjectQuery)
	}
	sq.withNamedSubjects[name] = query
	return sq
}

// WithNamedDsse tells the query-builder to eager-load the nodes that are connected to the "dsse"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (sq *StatementQuery) WithNamedDsse(name string, opts ...func(*DsseQuery)) *StatementQuery {
	query := (&DsseClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if sq.withNamedDsse == nil {
		sq.withNamedDsse = make(map[string]*DsseQuery)
	}
	sq.withNamedDsse[name] = query
	return sq
}

// StatementGroupBy is the group-by builder for Statement entities.
type StatementGroupBy struct {
	selector
	build *StatementQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *StatementGroupBy) Aggregate(fns ...AggregateFunc) *StatementGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *StatementGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StatementQuery, *StatementGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *StatementGroupBy) sqlScan(ctx context.Context, root *StatementQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StatementSelect is the builder for selecting fields of Statement entities.
type StatementSelect struct {
	*StatementQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *StatementSelect) Aggregate(fns ...AggregateFunc) *StatementSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *StatementSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StatementQuery, *StatementSelect](ctx, ss.StatementQuery, ss, ss.inters, v)
}

func (ss *StatementSelect) sqlScan(ctx context.Context, root *StatementQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
