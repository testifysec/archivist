// Code generated by ent, DO NOT EDIT.

package project

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/testifysec/judge/judge-api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// RepoID applies equality check predicate on the "repo_id" field. It's identical to RepoIDEQ.
func RepoID(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRepoID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// Projecturl applies equality check predicate on the "projecturl" field. It's identical to ProjecturlEQ.
func Projecturl(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProjecturl, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldUpdatedAt, v))
}

// RepoIDEQ applies the EQ predicate on the "repo_id" field.
func RepoIDEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldRepoID, v))
}

// RepoIDNEQ applies the NEQ predicate on the "repo_id" field.
func RepoIDNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldRepoID, v))
}

// RepoIDIn applies the In predicate on the "repo_id" field.
func RepoIDIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldRepoID, vs...))
}

// RepoIDNotIn applies the NotIn predicate on the "repo_id" field.
func RepoIDNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldRepoID, vs...))
}

// RepoIDGT applies the GT predicate on the "repo_id" field.
func RepoIDGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldRepoID, v))
}

// RepoIDGTE applies the GTE predicate on the "repo_id" field.
func RepoIDGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldRepoID, v))
}

// RepoIDLT applies the LT predicate on the "repo_id" field.
func RepoIDLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldRepoID, v))
}

// RepoIDLTE applies the LTE predicate on the "repo_id" field.
func RepoIDLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldRepoID, v))
}

// RepoIDContains applies the Contains predicate on the "repo_id" field.
func RepoIDContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldRepoID, v))
}

// RepoIDHasPrefix applies the HasPrefix predicate on the "repo_id" field.
func RepoIDHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldRepoID, v))
}

// RepoIDHasSuffix applies the HasSuffix predicate on the "repo_id" field.
func RepoIDHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldRepoID, v))
}

// RepoIDEqualFold applies the EqualFold predicate on the "repo_id" field.
func RepoIDEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldRepoID, v))
}

// RepoIDContainsFold applies the ContainsFold predicate on the "repo_id" field.
func RepoIDContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldRepoID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldName, v))
}

// ProjecturlEQ applies the EQ predicate on the "projecturl" field.
func ProjecturlEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldEQ(FieldProjecturl, v))
}

// ProjecturlNEQ applies the NEQ predicate on the "projecturl" field.
func ProjecturlNEQ(v string) predicate.Project {
	return predicate.Project(sql.FieldNEQ(FieldProjecturl, v))
}

// ProjecturlIn applies the In predicate on the "projecturl" field.
func ProjecturlIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldIn(FieldProjecturl, vs...))
}

// ProjecturlNotIn applies the NotIn predicate on the "projecturl" field.
func ProjecturlNotIn(vs ...string) predicate.Project {
	return predicate.Project(sql.FieldNotIn(FieldProjecturl, vs...))
}

// ProjecturlGT applies the GT predicate on the "projecturl" field.
func ProjecturlGT(v string) predicate.Project {
	return predicate.Project(sql.FieldGT(FieldProjecturl, v))
}

// ProjecturlGTE applies the GTE predicate on the "projecturl" field.
func ProjecturlGTE(v string) predicate.Project {
	return predicate.Project(sql.FieldGTE(FieldProjecturl, v))
}

// ProjecturlLT applies the LT predicate on the "projecturl" field.
func ProjecturlLT(v string) predicate.Project {
	return predicate.Project(sql.FieldLT(FieldProjecturl, v))
}

// ProjecturlLTE applies the LTE predicate on the "projecturl" field.
func ProjecturlLTE(v string) predicate.Project {
	return predicate.Project(sql.FieldLTE(FieldProjecturl, v))
}

// ProjecturlContains applies the Contains predicate on the "projecturl" field.
func ProjecturlContains(v string) predicate.Project {
	return predicate.Project(sql.FieldContains(FieldProjecturl, v))
}

// ProjecturlHasPrefix applies the HasPrefix predicate on the "projecturl" field.
func ProjecturlHasPrefix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasPrefix(FieldProjecturl, v))
}

// ProjecturlHasSuffix applies the HasSuffix predicate on the "projecturl" field.
func ProjecturlHasSuffix(v string) predicate.Project {
	return predicate.Project(sql.FieldHasSuffix(FieldProjecturl, v))
}

// ProjecturlEqualFold applies the EqualFold predicate on the "projecturl" field.
func ProjecturlEqualFold(v string) predicate.Project {
	return predicate.Project(sql.FieldEqualFold(FieldProjecturl, v))
}

// ProjecturlContainsFold applies the ContainsFold predicate on the "projecturl" field.
func ProjecturlContainsFold(v string) predicate.Project {
	return predicate.Project(sql.FieldContainsFold(FieldProjecturl, v))
}

// HasTenant applies the HasEdge predicate on the "tenant" edge.
func HasTenant() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTenantWith applies the HasEdge predicate on the "tenant" edge with a given conditions (other predicates).
func HasTenantWith(preds ...predicate.Tenant) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TenantInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TenantTable, TenantColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreatedBy applies the HasEdge predicate on the "created_by" edge.
func HasCreatedBy() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatedByWith applies the HasEdge predicate on the "created_by" edge with a given conditions (other predicates).
func HasCreatedByWith(preds ...predicate.User) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatedByInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CreatedByTable, CreatedByColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasModifiedBy applies the HasEdge predicate on the "modified_by" edge.
func HasModifiedBy() predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ModifiedByTable, ModifiedByColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModifiedByWith applies the HasEdge predicate on the "modified_by" edge with a given conditions (other predicates).
func HasModifiedByWith(preds ...predicate.User) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ModifiedByInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ModifiedByTable, ModifiedByColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Project) predicate.Project {
	return predicate.Project(func(s *sql.Selector) {
		p(s.Not())
	})
}
