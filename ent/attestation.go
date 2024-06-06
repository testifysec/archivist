// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/attestation"
	"github.com/in-toto/archivista/ent/attestationcollection"
	"github.com/in-toto/archivista/ent/gitattestation"
)

// Attestation is the model entity for the Attestation schema.
type Attestation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttestationQuery when eager-loading is set.
	Edges                               AttestationEdges `json:"edges"`
	attestation_collection_attestations *uuid.UUID
	selectValues                        sql.SelectValues
}

// AttestationEdges holds the relations/edges for other nodes in the graph.
type AttestationEdges struct {
	// AttestationCollection holds the value of the attestation_collection edge.
	AttestationCollection *AttestationCollection `json:"attestation_collection,omitempty"`
	// GitAttestation holds the value of the git_attestation edge.
	GitAttestation *GitAttestation `json:"git_attestation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// AttestationCollectionOrErr returns the AttestationCollection value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttestationEdges) AttestationCollectionOrErr() (*AttestationCollection, error) {
	if e.AttestationCollection != nil {
		return e.AttestationCollection, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: attestationcollection.Label}
	}
	return nil, &NotLoadedError{edge: "attestation_collection"}
}

// GitAttestationOrErr returns the GitAttestation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttestationEdges) GitAttestationOrErr() (*GitAttestation, error) {
	if e.GitAttestation != nil {
		return e.GitAttestation, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: gitattestation.Label}
	}
	return nil, &NotLoadedError{edge: "git_attestation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attestation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attestation.FieldType:
			values[i] = new(sql.NullString)
		case attestation.FieldID:
			values[i] = new(uuid.UUID)
		case attestation.ForeignKeys[0]: // attestation_collection_attestations
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attestation fields.
func (a *Attestation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attestation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case attestation.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				a.Type = value.String
			}
		case attestation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field attestation_collection_attestations", values[i])
			} else if value.Valid {
				a.attestation_collection_attestations = new(uuid.UUID)
				*a.attestation_collection_attestations = *value.S.(*uuid.UUID)
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attestation.
// This includes values selected through modifiers, order, etc.
func (a *Attestation) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryAttestationCollection queries the "attestation_collection" edge of the Attestation entity.
func (a *Attestation) QueryAttestationCollection() *AttestationCollectionQuery {
	return NewAttestationClient(a.config).QueryAttestationCollection(a)
}

// QueryGitAttestation queries the "git_attestation" edge of the Attestation entity.
func (a *Attestation) QueryGitAttestation() *GitAttestationQuery {
	return NewAttestationClient(a.config).QueryGitAttestation(a)
}

// Update returns a builder for updating this Attestation.
// Note that you need to call Attestation.Unwrap() before calling this method if this Attestation
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attestation) Update() *AttestationUpdateOne {
	return NewAttestationClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attestation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attestation) Unwrap() *Attestation {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attestation is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attestation) String() string {
	var builder strings.Builder
	builder.WriteString("Attestation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("type=")
	builder.WriteString(a.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Attestations is a parsable slice of Attestation.
type Attestations []*Attestation
