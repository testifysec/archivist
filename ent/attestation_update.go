// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/attestation"
	"github.com/in-toto/archivista/ent/attestationcollection"
	"github.com/in-toto/archivista/ent/gitattestation"
	"github.com/in-toto/archivista/ent/predicate"
)

// AttestationUpdate is the builder for updating Attestation entities.
type AttestationUpdate struct {
	config
	hooks    []Hook
	mutation *AttestationMutation
}

// Where appends a list predicates to the AttestationUpdate builder.
func (au *AttestationUpdate) Where(ps ...predicate.Attestation) *AttestationUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetType sets the "type" field.
func (au *AttestationUpdate) SetType(s string) *AttestationUpdate {
	au.mutation.SetType(s)
	return au
}

// SetNillableType sets the "type" field if the given value is not nil.
func (au *AttestationUpdate) SetNillableType(s *string) *AttestationUpdate {
	if s != nil {
		au.SetType(*s)
	}
	return au
}

// SetAttestationCollectionID sets the "attestation_collection" edge to the AttestationCollection entity by ID.
func (au *AttestationUpdate) SetAttestationCollectionID(id uuid.UUID) *AttestationUpdate {
	au.mutation.SetAttestationCollectionID(id)
	return au
}

// SetAttestationCollection sets the "attestation_collection" edge to the AttestationCollection entity.
func (au *AttestationUpdate) SetAttestationCollection(a *AttestationCollection) *AttestationUpdate {
	return au.SetAttestationCollectionID(a.ID)
}

// SetGitAttestationID sets the "git_attestation" edge to the GitAttestation entity by ID.
func (au *AttestationUpdate) SetGitAttestationID(id uuid.UUID) *AttestationUpdate {
	au.mutation.SetGitAttestationID(id)
	return au
}

// SetNillableGitAttestationID sets the "git_attestation" edge to the GitAttestation entity by ID if the given value is not nil.
func (au *AttestationUpdate) SetNillableGitAttestationID(id *uuid.UUID) *AttestationUpdate {
	if id != nil {
		au = au.SetGitAttestationID(*id)
	}
	return au
}

// SetGitAttestation sets the "git_attestation" edge to the GitAttestation entity.
func (au *AttestationUpdate) SetGitAttestation(g *GitAttestation) *AttestationUpdate {
	return au.SetGitAttestationID(g.ID)
}

// Mutation returns the AttestationMutation object of the builder.
func (au *AttestationUpdate) Mutation() *AttestationMutation {
	return au.mutation
}

// ClearAttestationCollection clears the "attestation_collection" edge to the AttestationCollection entity.
func (au *AttestationUpdate) ClearAttestationCollection() *AttestationUpdate {
	au.mutation.ClearAttestationCollection()
	return au
}

// ClearGitAttestation clears the "git_attestation" edge to the GitAttestation entity.
func (au *AttestationUpdate) ClearGitAttestation() *AttestationUpdate {
	au.mutation.ClearGitAttestation()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttestationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttestationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttestationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttestationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttestationUpdate) check() error {
	if v, ok := au.mutation.GetType(); ok {
		if err := attestation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Attestation.type": %w`, err)}
		}
	}
	if _, ok := au.mutation.AttestationCollectionID(); au.mutation.AttestationCollectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attestation.attestation_collection"`)
	}
	return nil
}

func (au *AttestationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attestation.Table, attestation.Columns, sqlgraph.NewFieldSpec(attestation.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.SetField(attestation.FieldType, field.TypeString, value)
	}
	if au.mutation.AttestationCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attestation.AttestationCollectionTable,
			Columns: []string{attestation.AttestationCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attestationcollection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AttestationCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attestation.AttestationCollectionTable,
			Columns: []string{attestation.AttestationCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attestationcollection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.GitAttestationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   attestation.GitAttestationTable,
			Columns: []string{attestation.GitAttestationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitattestation.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.GitAttestationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   attestation.GitAttestationTable,
			Columns: []string{attestation.GitAttestationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitattestation.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attestation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttestationUpdateOne is the builder for updating a single Attestation entity.
type AttestationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttestationMutation
}

// SetType sets the "type" field.
func (auo *AttestationUpdateOne) SetType(s string) *AttestationUpdateOne {
	auo.mutation.SetType(s)
	return auo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (auo *AttestationUpdateOne) SetNillableType(s *string) *AttestationUpdateOne {
	if s != nil {
		auo.SetType(*s)
	}
	return auo
}

// SetAttestationCollectionID sets the "attestation_collection" edge to the AttestationCollection entity by ID.
func (auo *AttestationUpdateOne) SetAttestationCollectionID(id uuid.UUID) *AttestationUpdateOne {
	auo.mutation.SetAttestationCollectionID(id)
	return auo
}

// SetAttestationCollection sets the "attestation_collection" edge to the AttestationCollection entity.
func (auo *AttestationUpdateOne) SetAttestationCollection(a *AttestationCollection) *AttestationUpdateOne {
	return auo.SetAttestationCollectionID(a.ID)
}

// SetGitAttestationID sets the "git_attestation" edge to the GitAttestation entity by ID.
func (auo *AttestationUpdateOne) SetGitAttestationID(id uuid.UUID) *AttestationUpdateOne {
	auo.mutation.SetGitAttestationID(id)
	return auo
}

// SetNillableGitAttestationID sets the "git_attestation" edge to the GitAttestation entity by ID if the given value is not nil.
func (auo *AttestationUpdateOne) SetNillableGitAttestationID(id *uuid.UUID) *AttestationUpdateOne {
	if id != nil {
		auo = auo.SetGitAttestationID(*id)
	}
	return auo
}

// SetGitAttestation sets the "git_attestation" edge to the GitAttestation entity.
func (auo *AttestationUpdateOne) SetGitAttestation(g *GitAttestation) *AttestationUpdateOne {
	return auo.SetGitAttestationID(g.ID)
}

// Mutation returns the AttestationMutation object of the builder.
func (auo *AttestationUpdateOne) Mutation() *AttestationMutation {
	return auo.mutation
}

// ClearAttestationCollection clears the "attestation_collection" edge to the AttestationCollection entity.
func (auo *AttestationUpdateOne) ClearAttestationCollection() *AttestationUpdateOne {
	auo.mutation.ClearAttestationCollection()
	return auo
}

// ClearGitAttestation clears the "git_attestation" edge to the GitAttestation entity.
func (auo *AttestationUpdateOne) ClearGitAttestation() *AttestationUpdateOne {
	auo.mutation.ClearGitAttestation()
	return auo
}

// Where appends a list predicates to the AttestationUpdate builder.
func (auo *AttestationUpdateOne) Where(ps ...predicate.Attestation) *AttestationUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttestationUpdateOne) Select(field string, fields ...string) *AttestationUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attestation entity.
func (auo *AttestationUpdateOne) Save(ctx context.Context) (*Attestation, error) {
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttestationUpdateOne) SaveX(ctx context.Context) *Attestation {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttestationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttestationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttestationUpdateOne) check() error {
	if v, ok := auo.mutation.GetType(); ok {
		if err := attestation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Attestation.type": %w`, err)}
		}
	}
	if _, ok := auo.mutation.AttestationCollectionID(); auo.mutation.AttestationCollectionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attestation.attestation_collection"`)
	}
	return nil
}

func (auo *AttestationUpdateOne) sqlSave(ctx context.Context) (_node *Attestation, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attestation.Table, attestation.Columns, sqlgraph.NewFieldSpec(attestation.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attestation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attestation.FieldID)
		for _, f := range fields {
			if !attestation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attestation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.SetField(attestation.FieldType, field.TypeString, value)
	}
	if auo.mutation.AttestationCollectionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attestation.AttestationCollectionTable,
			Columns: []string{attestation.AttestationCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attestationcollection.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AttestationCollectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attestation.AttestationCollectionTable,
			Columns: []string{attestation.AttestationCollectionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attestationcollection.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.GitAttestationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   attestation.GitAttestationTable,
			Columns: []string{attestation.GitAttestationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitattestation.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.GitAttestationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   attestation.GitAttestationTable,
			Columns: []string{attestation.GitAttestationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(gitattestation.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attestation{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attestation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
