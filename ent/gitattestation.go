// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/in-toto/archivista/ent/attestation"
	"github.com/in-toto/archivista/ent/gitattestation"
)

// GitAttestation is the model entity for the GitAttestation schema.
type GitAttestation struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CommitHash holds the value of the "commit_hash" field.
	CommitHash string `json:"commit_hash,omitempty"`
	// Author holds the value of the "author" field.
	Author string `json:"author,omitempty"`
	// AuthorEmail holds the value of the "author_email" field.
	AuthorEmail string `json:"author_email,omitempty"`
	// CommitterName holds the value of the "committer_name" field.
	CommitterName string `json:"committer_name,omitempty"`
	// CommitterEmail holds the value of the "committer_email" field.
	CommitterEmail string `json:"committer_email,omitempty"`
	// CommitDate holds the value of the "commit_date" field.
	CommitDate string `json:"commit_date,omitempty"`
	// CommitMessage holds the value of the "commit_message" field.
	CommitMessage string `json:"commit_message,omitempty"`
	// Status holds the value of the "status" field.
	Status []string `json:"status,omitempty"`
	// CommitType holds the value of the "commit_type" field.
	CommitType string `json:"commit_type,omitempty"`
	// CommitDigest holds the value of the "commit_digest" field.
	CommitDigest string `json:"commit_digest,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature string `json:"signature,omitempty"`
	// ParentHashes holds the value of the "parent_hashes" field.
	ParentHashes []string `json:"parent_hashes,omitempty"`
	// TreeHash holds the value of the "tree_hash" field.
	TreeHash string `json:"tree_hash,omitempty"`
	// Refs holds the value of the "refs" field.
	Refs []string `json:"refs,omitempty"`
	// Remotes holds the value of the "remotes" field.
	Remotes []string `json:"remotes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GitAttestationQuery when eager-loading is set.
	Edges                       GitAttestationEdges `json:"edges"`
	attestation_git_attestation *uuid.UUID
	selectValues                sql.SelectValues
}

// GitAttestationEdges holds the relations/edges for other nodes in the graph.
type GitAttestationEdges struct {
	// Attestation holds the value of the attestation edge.
	Attestation *Attestation `json:"attestation,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// AttestationOrErr returns the Attestation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GitAttestationEdges) AttestationOrErr() (*Attestation, error) {
	if e.Attestation != nil {
		return e.Attestation, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: attestation.Label}
	}
	return nil, &NotLoadedError{edge: "attestation"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GitAttestation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case gitattestation.FieldStatus, gitattestation.FieldParentHashes, gitattestation.FieldRefs, gitattestation.FieldRemotes:
			values[i] = new([]byte)
		case gitattestation.FieldCommitHash, gitattestation.FieldAuthor, gitattestation.FieldAuthorEmail, gitattestation.FieldCommitterName, gitattestation.FieldCommitterEmail, gitattestation.FieldCommitDate, gitattestation.FieldCommitMessage, gitattestation.FieldCommitType, gitattestation.FieldCommitDigest, gitattestation.FieldSignature, gitattestation.FieldTreeHash:
			values[i] = new(sql.NullString)
		case gitattestation.FieldID:
			values[i] = new(uuid.UUID)
		case gitattestation.ForeignKeys[0]: // attestation_git_attestation
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GitAttestation fields.
func (ga *GitAttestation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case gitattestation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ga.ID = *value
			}
		case gitattestation.FieldCommitHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_hash", values[i])
			} else if value.Valid {
				ga.CommitHash = value.String
			}
		case gitattestation.FieldAuthor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author", values[i])
			} else if value.Valid {
				ga.Author = value.String
			}
		case gitattestation.FieldAuthorEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field author_email", values[i])
			} else if value.Valid {
				ga.AuthorEmail = value.String
			}
		case gitattestation.FieldCommitterName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field committer_name", values[i])
			} else if value.Valid {
				ga.CommitterName = value.String
			}
		case gitattestation.FieldCommitterEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field committer_email", values[i])
			} else if value.Valid {
				ga.CommitterEmail = value.String
			}
		case gitattestation.FieldCommitDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_date", values[i])
			} else if value.Valid {
				ga.CommitDate = value.String
			}
		case gitattestation.FieldCommitMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_message", values[i])
			} else if value.Valid {
				ga.CommitMessage = value.String
			}
		case gitattestation.FieldStatus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ga.Status); err != nil {
					return fmt.Errorf("unmarshal field status: %w", err)
				}
			}
		case gitattestation.FieldCommitType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_type", values[i])
			} else if value.Valid {
				ga.CommitType = value.String
			}
		case gitattestation.FieldCommitDigest:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field commit_digest", values[i])
			} else if value.Valid {
				ga.CommitDigest = value.String
			}
		case gitattestation.FieldSignature:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value.Valid {
				ga.Signature = value.String
			}
		case gitattestation.FieldParentHashes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field parent_hashes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ga.ParentHashes); err != nil {
					return fmt.Errorf("unmarshal field parent_hashes: %w", err)
				}
			}
		case gitattestation.FieldTreeHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tree_hash", values[i])
			} else if value.Valid {
				ga.TreeHash = value.String
			}
		case gitattestation.FieldRefs:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field refs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ga.Refs); err != nil {
					return fmt.Errorf("unmarshal field refs: %w", err)
				}
			}
		case gitattestation.FieldRemotes:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field remotes", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ga.Remotes); err != nil {
					return fmt.Errorf("unmarshal field remotes: %w", err)
				}
			}
		case gitattestation.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field attestation_git_attestation", values[i])
			} else if value.Valid {
				ga.attestation_git_attestation = new(uuid.UUID)
				*ga.attestation_git_attestation = *value.S.(*uuid.UUID)
			}
		default:
			ga.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GitAttestation.
// This includes values selected through modifiers, order, etc.
func (ga *GitAttestation) Value(name string) (ent.Value, error) {
	return ga.selectValues.Get(name)
}

// QueryAttestation queries the "attestation" edge of the GitAttestation entity.
func (ga *GitAttestation) QueryAttestation() *AttestationQuery {
	return NewGitAttestationClient(ga.config).QueryAttestation(ga)
}

// Update returns a builder for updating this GitAttestation.
// Note that you need to call GitAttestation.Unwrap() before calling this method if this GitAttestation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *GitAttestation) Update() *GitAttestationUpdateOne {
	return NewGitAttestationClient(ga.config).UpdateOne(ga)
}

// Unwrap unwraps the GitAttestation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *GitAttestation) Unwrap() *GitAttestation {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: GitAttestation is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *GitAttestation) String() string {
	var builder strings.Builder
	builder.WriteString("GitAttestation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("commit_hash=")
	builder.WriteString(ga.CommitHash)
	builder.WriteString(", ")
	builder.WriteString("author=")
	builder.WriteString(ga.Author)
	builder.WriteString(", ")
	builder.WriteString("author_email=")
	builder.WriteString(ga.AuthorEmail)
	builder.WriteString(", ")
	builder.WriteString("committer_name=")
	builder.WriteString(ga.CommitterName)
	builder.WriteString(", ")
	builder.WriteString("committer_email=")
	builder.WriteString(ga.CommitterEmail)
	builder.WriteString(", ")
	builder.WriteString("commit_date=")
	builder.WriteString(ga.CommitDate)
	builder.WriteString(", ")
	builder.WriteString("commit_message=")
	builder.WriteString(ga.CommitMessage)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", ga.Status))
	builder.WriteString(", ")
	builder.WriteString("commit_type=")
	builder.WriteString(ga.CommitType)
	builder.WriteString(", ")
	builder.WriteString("commit_digest=")
	builder.WriteString(ga.CommitDigest)
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(ga.Signature)
	builder.WriteString(", ")
	builder.WriteString("parent_hashes=")
	builder.WriteString(fmt.Sprintf("%v", ga.ParentHashes))
	builder.WriteString(", ")
	builder.WriteString("tree_hash=")
	builder.WriteString(ga.TreeHash)
	builder.WriteString(", ")
	builder.WriteString("refs=")
	builder.WriteString(fmt.Sprintf("%v", ga.Refs))
	builder.WriteString(", ")
	builder.WriteString("remotes=")
	builder.WriteString(fmt.Sprintf("%v", ga.Remotes))
	builder.WriteByte(')')
	return builder.String()
}

// GitAttestations is a parsable slice of GitAttestation.
type GitAttestations []*GitAttestation
