// Code generated by ent, DO NOT EDIT.

package attestation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the attestation type in the database.
	Label = "attestation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeOmnitrail holds the string denoting the omnitrail edge name in mutations.
	EdgeOmnitrail = "omnitrail"
	// EdgeAttestationCollection holds the string denoting the attestation_collection edge name in mutations.
	EdgeAttestationCollection = "attestation_collection"
	// Table holds the table name of the attestation in the database.
	Table = "attestations"
	// OmnitrailTable is the table that holds the omnitrail relation/edge.
	OmnitrailTable = "omnitrails"
	// OmnitrailInverseTable is the table name for the Omnitrail entity.
	// It exists in this package in order to avoid circular dependency with the "omnitrail" package.
	OmnitrailInverseTable = "omnitrails"
	// OmnitrailColumn is the table column denoting the omnitrail relation/edge.
	OmnitrailColumn = "attestation_omnitrail"
	// AttestationCollectionTable is the table that holds the attestation_collection relation/edge.
	AttestationCollectionTable = "attestations"
	// AttestationCollectionInverseTable is the table name for the AttestationCollection entity.
	// It exists in this package in order to avoid circular dependency with the "attestationcollection" package.
	AttestationCollectionInverseTable = "attestation_collections"
	// AttestationCollectionColumn is the table column denoting the attestation_collection relation/edge.
	AttestationCollectionColumn = "attestation_collection_attestations"
)

// Columns holds all SQL columns for attestation fields.
var Columns = []string{
	FieldID,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "attestations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"attestation_collection_attestations",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Attestation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByOmnitrailField orders the results by omnitrail field.
func ByOmnitrailField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOmnitrailStep(), sql.OrderByField(field, opts...))
	}
}

// ByAttestationCollectionField orders the results by attestation_collection field.
func ByAttestationCollectionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttestationCollectionStep(), sql.OrderByField(field, opts...))
	}
}
func newOmnitrailStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OmnitrailInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, OmnitrailTable, OmnitrailColumn),
	)
}
func newAttestationCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttestationCollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AttestationCollectionTable, AttestationCollectionColumn),
	)
}
