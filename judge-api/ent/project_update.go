// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/testifysec/judge/judge-api/ent/predicate"
	"github.com/testifysec/judge/judge-api/ent/project"
	"github.com/testifysec/judge/judge-api/ent/user"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProjectUpdate) SetUpdatedAt(t time.Time) *ProjectUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetRepoID sets the "repo_id" field.
func (pu *ProjectUpdate) SetRepoID(s string) *ProjectUpdate {
	pu.mutation.SetRepoID(s)
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetProjecturl sets the "projecturl" field.
func (pu *ProjectUpdate) SetProjecturl(s string) *ProjectUpdate {
	pu.mutation.SetProjecturl(s)
	return pu
}

// SetModifiedByID sets the "modified_by" edge to the User entity by ID.
func (pu *ProjectUpdate) SetModifiedByID(id uuid.UUID) *ProjectUpdate {
	pu.mutation.SetModifiedByID(id)
	return pu
}

// SetNillableModifiedByID sets the "modified_by" edge to the User entity by ID if the given value is not nil.
func (pu *ProjectUpdate) SetNillableModifiedByID(id *uuid.UUID) *ProjectUpdate {
	if id != nil {
		pu = pu.SetModifiedByID(*id)
	}
	return pu
}

// SetModifiedBy sets the "modified_by" edge to the User entity.
func (pu *ProjectUpdate) SetModifiedBy(u *User) *ProjectUpdate {
	return pu.SetModifiedByID(u.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// ClearModifiedBy clears the "modified_by" edge to the User entity.
func (pu *ProjectUpdate) ClearModifiedBy() *ProjectUpdate {
	pu.mutation.ClearModifiedBy()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, ProjectMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProjectUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if project.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized project.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := project.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.RepoID(); ok {
		if err := project.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "Project.repo_id": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Projecturl(); ok {
		if err := project.ProjecturlValidator(v); err != nil {
			return &ValidationError{Name: "projecturl", err: fmt.Errorf(`ent: validator failed for field "Project.projecturl": %w`, err)}
		}
	}
	if _, ok := pu.mutation.TenantID(); pu.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Project.tenant"`)
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.RepoID(); ok {
		_spec.SetField(project.FieldRepoID, field.TypeString, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Projecturl(); ok {
		_spec.SetField(project.FieldProjecturl, field.TypeString, value)
	}
	if pu.mutation.ModifiedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.ModifiedByTable,
			Columns: []string{project.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ModifiedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.ModifiedByTable,
			Columns: []string{project.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProjectUpdateOne) SetUpdatedAt(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetRepoID sets the "repo_id" field.
func (puo *ProjectUpdateOne) SetRepoID(s string) *ProjectUpdateOne {
	puo.mutation.SetRepoID(s)
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetProjecturl sets the "projecturl" field.
func (puo *ProjectUpdateOne) SetProjecturl(s string) *ProjectUpdateOne {
	puo.mutation.SetProjecturl(s)
	return puo
}

// SetModifiedByID sets the "modified_by" edge to the User entity by ID.
func (puo *ProjectUpdateOne) SetModifiedByID(id uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetModifiedByID(id)
	return puo
}

// SetNillableModifiedByID sets the "modified_by" edge to the User entity by ID if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableModifiedByID(id *uuid.UUID) *ProjectUpdateOne {
	if id != nil {
		puo = puo.SetModifiedByID(*id)
	}
	return puo
}

// SetModifiedBy sets the "modified_by" edge to the User entity.
func (puo *ProjectUpdateOne) SetModifiedBy(u *User) *ProjectUpdateOne {
	return puo.SetModifiedByID(u.ID)
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// ClearModifiedBy clears the "modified_by" edge to the User entity.
func (puo *ProjectUpdateOne) ClearModifiedBy() *ProjectUpdateOne {
	puo.mutation.ClearModifiedBy()
	return puo
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Project, ProjectMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProjectUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if project.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized project.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := project.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.RepoID(); ok {
		if err := project.RepoIDValidator(v); err != nil {
			return &ValidationError{Name: "repo_id", err: fmt.Errorf(`ent: validator failed for field "Project.repo_id": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Projecturl(); ok {
		if err := project.ProjecturlValidator(v); err != nil {
			return &ValidationError{Name: "projecturl", err: fmt.Errorf(`ent: validator failed for field "Project.projecturl": %w`, err)}
		}
	}
	if _, ok := puo.mutation.TenantID(); puo.mutation.TenantCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Project.tenant"`)
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(project.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.RepoID(); ok {
		_spec.SetField(project.FieldRepoID, field.TypeString, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Projecturl(); ok {
		_spec.SetField(project.FieldProjecturl, field.TypeString, value)
	}
	if puo.mutation.ModifiedByCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.ModifiedByTable,
			Columns: []string{project.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ModifiedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   project.ModifiedByTable,
			Columns: []string{project.ModifiedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
