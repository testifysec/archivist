// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttestationsColumns holds the columns for the "attestations" table.
	AttestationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "attestation_collection_attestations", Type: field.TypeUUID},
	}
	// AttestationsTable holds the schema information for the "attestations" table.
	AttestationsTable = &schema.Table{
		Name:       "attestations",
		Columns:    AttestationsColumns,
		PrimaryKey: []*schema.Column{AttestationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attestations_attestation_collections_attestations",
				Columns:    []*schema.Column{AttestationsColumns[2]},
				RefColumns: []*schema.Column{AttestationCollectionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attestation_type",
				Unique:  false,
				Columns: []*schema.Column{AttestationsColumns[1]},
			},
		},
	}
	// AttestationCollectionsColumns holds the columns for the "attestation_collections" table.
	AttestationCollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "statement_attestation_collections", Type: field.TypeUUID, Unique: true},
	}
	// AttestationCollectionsTable holds the schema information for the "attestation_collections" table.
	AttestationCollectionsTable = &schema.Table{
		Name:       "attestation_collections",
		Columns:    AttestationCollectionsColumns,
		PrimaryKey: []*schema.Column{AttestationCollectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attestation_collections_statements_attestation_collections",
				Columns:    []*schema.Column{AttestationCollectionsColumns[2]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attestationcollection_name",
				Unique:  false,
				Columns: []*schema.Column{AttestationCollectionsColumns[1]},
			},
		},
	}
	// AttestationPoliciesColumns holds the columns for the "attestation_policies" table.
	AttestationPoliciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "statement_policy", Type: field.TypeUUID, Unique: true, Nullable: true},
	}
	// AttestationPoliciesTable holds the schema information for the "attestation_policies" table.
	AttestationPoliciesTable = &schema.Table{
		Name:       "attestation_policies",
		Columns:    AttestationPoliciesColumns,
		PrimaryKey: []*schema.Column{AttestationPoliciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attestation_policies_statements_policy",
				Columns:    []*schema.Column{AttestationPoliciesColumns[2]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "attestationpolicy_name",
				Unique:  false,
				Columns: []*schema.Column{AttestationPoliciesColumns[1]},
			},
		},
	}
	// DssesColumns holds the columns for the "dsses" table.
	DssesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "gitoid_sha256", Type: field.TypeString, Unique: true},
		{Name: "payload_type", Type: field.TypeString},
		{Name: "dsse_statement", Type: field.TypeUUID, Nullable: true},
	}
	// DssesTable holds the schema information for the "dsses" table.
	DssesTable = &schema.Table{
		Name:       "dsses",
		Columns:    DssesColumns,
		PrimaryKey: []*schema.Column{DssesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "dsses_statements_statement",
				Columns:    []*schema.Column{DssesColumns[3]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GitAttestationsColumns holds the columns for the "git_attestations" table.
	GitAttestationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "commit_hash", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "author_email", Type: field.TypeString},
		{Name: "committer_name", Type: field.TypeString},
		{Name: "committer_email", Type: field.TypeString},
		{Name: "commit_date", Type: field.TypeString},
		{Name: "commit_message", Type: field.TypeString},
		{Name: "status", Type: field.TypeJSON},
		{Name: "commit_type", Type: field.TypeString},
		{Name: "commit_digest", Type: field.TypeString},
		{Name: "signature", Type: field.TypeString},
		{Name: "parent_hashes", Type: field.TypeJSON},
		{Name: "tree_hash", Type: field.TypeString},
		{Name: "refs", Type: field.TypeJSON},
		{Name: "remotes", Type: field.TypeJSON},
		{Name: "attestation_git_attestation", Type: field.TypeUUID, Unique: true},
	}
	// GitAttestationsTable holds the schema information for the "git_attestations" table.
	GitAttestationsTable = &schema.Table{
		Name:       "git_attestations",
		Columns:    GitAttestationsColumns,
		PrimaryKey: []*schema.Column{GitAttestationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "git_attestations_attestations_git_attestation",
				Columns:    []*schema.Column{GitAttestationsColumns[16]},
				RefColumns: []*schema.Column{AttestationsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PayloadDigestsColumns holds the columns for the "payload_digests" table.
	PayloadDigestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "dsse_payload_digests", Type: field.TypeUUID, Nullable: true},
	}
	// PayloadDigestsTable holds the schema information for the "payload_digests" table.
	PayloadDigestsTable = &schema.Table{
		Name:       "payload_digests",
		Columns:    PayloadDigestsColumns,
		PrimaryKey: []*schema.Column{PayloadDigestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "payload_digests_dsses_payload_digests",
				Columns:    []*schema.Column{PayloadDigestsColumns[3]},
				RefColumns: []*schema.Column{DssesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "payloaddigest_value",
				Unique:  false,
				Columns: []*schema.Column{PayloadDigestsColumns[2]},
			},
		},
	}
	// SignaturesColumns holds the columns for the "signatures" table.
	SignaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "key_id", Type: field.TypeString},
		{Name: "signature", Type: field.TypeString, SchemaType: map[string]string{"mysql": "text"}},
		{Name: "dsse_signatures", Type: field.TypeUUID, Nullable: true},
	}
	// SignaturesTable holds the schema information for the "signatures" table.
	SignaturesTable = &schema.Table{
		Name:       "signatures",
		Columns:    SignaturesColumns,
		PrimaryKey: []*schema.Column{SignaturesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "signatures_dsses_signatures",
				Columns:    []*schema.Column{SignaturesColumns[3]},
				RefColumns: []*schema.Column{DssesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "signature_key_id",
				Unique:  false,
				Columns: []*schema.Column{SignaturesColumns[1]},
			},
		},
	}
	// StatementsColumns holds the columns for the "statements" table.
	StatementsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "predicate", Type: field.TypeString},
	}
	// StatementsTable holds the schema information for the "statements" table.
	StatementsTable = &schema.Table{
		Name:       "statements",
		Columns:    StatementsColumns,
		PrimaryKey: []*schema.Column{StatementsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "statement_predicate",
				Unique:  false,
				Columns: []*schema.Column{StatementsColumns[1]},
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "statement_subjects", Type: field.TypeUUID, Nullable: true},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subjects_statements_subjects",
				Columns:    []*schema.Column{SubjectsColumns[2]},
				RefColumns: []*schema.Column{StatementsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subject_name",
				Unique:  false,
				Columns: []*schema.Column{SubjectsColumns[1]},
			},
		},
	}
	// SubjectDigestsColumns holds the columns for the "subject_digests" table.
	SubjectDigestsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "algorithm", Type: field.TypeString},
		{Name: "value", Type: field.TypeString},
		{Name: "subject_subject_digests", Type: field.TypeUUID, Nullable: true},
	}
	// SubjectDigestsTable holds the schema information for the "subject_digests" table.
	SubjectDigestsTable = &schema.Table{
		Name:       "subject_digests",
		Columns:    SubjectDigestsColumns,
		PrimaryKey: []*schema.Column{SubjectDigestsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subject_digests_subjects_subject_digests",
				Columns:    []*schema.Column{SubjectDigestsColumns[3]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "subjectdigest_value",
				Unique:  false,
				Columns: []*schema.Column{SubjectDigestsColumns[2]},
			},
		},
	}
	// TimestampsColumns holds the columns for the "timestamps" table.
	TimestampsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeString},
		{Name: "timestamp", Type: field.TypeTime},
		{Name: "signature_timestamps", Type: field.TypeUUID, Nullable: true},
	}
	// TimestampsTable holds the schema information for the "timestamps" table.
	TimestampsTable = &schema.Table{
		Name:       "timestamps",
		Columns:    TimestampsColumns,
		PrimaryKey: []*schema.Column{TimestampsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timestamps_signatures_timestamps",
				Columns:    []*schema.Column{TimestampsColumns[3]},
				RefColumns: []*schema.Column{SignaturesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttestationsTable,
		AttestationCollectionsTable,
		AttestationPoliciesTable,
		DssesTable,
		GitAttestationsTable,
		PayloadDigestsTable,
		SignaturesTable,
		StatementsTable,
		SubjectsTable,
		SubjectDigestsTable,
		TimestampsTable,
	}
)

func init() {
	AttestationsTable.ForeignKeys[0].RefTable = AttestationCollectionsTable
	AttestationCollectionsTable.ForeignKeys[0].RefTable = StatementsTable
	AttestationPoliciesTable.ForeignKeys[0].RefTable = StatementsTable
	DssesTable.ForeignKeys[0].RefTable = StatementsTable
	GitAttestationsTable.ForeignKeys[0].RefTable = AttestationsTable
	PayloadDigestsTable.ForeignKeys[0].RefTable = DssesTable
	SignaturesTable.ForeignKeys[0].RefTable = DssesTable
	SubjectsTable.ForeignKeys[0].RefTable = StatementsTable
	SubjectDigestsTable.ForeignKeys[0].RefTable = SubjectsTable
	TimestampsTable.ForeignKeys[0].RefTable = SignaturesTable
}
