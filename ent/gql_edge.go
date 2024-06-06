// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Attestation) AttestationCollection(ctx context.Context) (*AttestationCollection, error) {
	result, err := a.Edges.AttestationCollectionOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryAttestationCollection().Only(ctx)
	}
	return result, err
}

func (a *Attestation) GitAttestation(ctx context.Context) (*GitAttestation, error) {
	result, err := a.Edges.GitAttestationOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryGitAttestation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ac *AttestationCollection) Attestations(ctx context.Context) (result []*Attestation, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ac.NamedAttestations(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ac.Edges.AttestationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ac.QueryAttestations().All(ctx)
	}
	return result, err
}

func (ac *AttestationCollection) Statement(ctx context.Context) (*Statement, error) {
	result, err := ac.Edges.StatementOrErr()
	if IsNotLoaded(err) {
		result, err = ac.QueryStatement().Only(ctx)
	}
	return result, err
}

func (ap *AttestationPolicy) Statement(ctx context.Context) (*Statement, error) {
	result, err := ap.Edges.StatementOrErr()
	if IsNotLoaded(err) {
		result, err = ap.QueryStatement().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (d *Dsse) Statement(ctx context.Context) (*Statement, error) {
	result, err := d.Edges.StatementOrErr()
	if IsNotLoaded(err) {
		result, err = d.QueryStatement().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (d *Dsse) Signatures(ctx context.Context) (result []*Signature, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = d.NamedSignatures(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = d.Edges.SignaturesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = d.QuerySignatures().All(ctx)
	}
	return result, err
}

func (d *Dsse) PayloadDigests(ctx context.Context) (result []*PayloadDigest, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = d.NamedPayloadDigests(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = d.Edges.PayloadDigestsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = d.QueryPayloadDigests().All(ctx)
	}
	return result, err
}

func (ga *GitAttestation) Attestation(ctx context.Context) (*Attestation, error) {
	result, err := ga.Edges.AttestationOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryAttestation().Only(ctx)
	}
	return result, err
}

func (pd *PayloadDigest) Dsse(ctx context.Context) (*Dsse, error) {
	result, err := pd.Edges.DsseOrErr()
	if IsNotLoaded(err) {
		result, err = pd.QueryDsse().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Signature) Dsse(ctx context.Context) (*Dsse, error) {
	result, err := s.Edges.DsseOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryDsse().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Signature) Timestamps(ctx context.Context) (result []*Timestamp, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedTimestamps(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.TimestampsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryTimestamps().All(ctx)
	}
	return result, err
}

func (s *Statement) Subjects(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *SubjectWhereInput,
) (*SubjectConnection, error) {
	opts := []SubjectPaginateOption{
		WithSubjectFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := s.Edges.totalCount[0][alias]
	if nodes, err := s.NamedSubjects(alias); err == nil || hasTotalCount {
		pager, err := newSubjectPager(opts, last != nil)
		if err != nil {
			return nil, err
		}
		conn := &SubjectConnection{Edges: []*SubjectEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return s.QuerySubjects().Paginate(ctx, after, first, before, last, opts...)
}

func (s *Statement) Policy(ctx context.Context) (*AttestationPolicy, error) {
	result, err := s.Edges.PolicyOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryPolicy().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Statement) AttestationCollections(ctx context.Context) (*AttestationCollection, error) {
	result, err := s.Edges.AttestationCollectionsOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryAttestationCollections().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Statement) Dsse(ctx context.Context) (result []*Dsse, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedDsse(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.DsseOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryDsse().All(ctx)
	}
	return result, err
}

func (s *Subject) SubjectDigests(ctx context.Context) (result []*SubjectDigest, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedSubjectDigests(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.SubjectDigestsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QuerySubjectDigests().All(ctx)
	}
	return result, err
}

func (s *Subject) Statement(ctx context.Context) (*Statement, error) {
	result, err := s.Edges.StatementOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryStatement().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (sd *SubjectDigest) Subject(ctx context.Context) (*Subject, error) {
	result, err := sd.Edges.SubjectOrErr()
	if IsNotLoaded(err) {
		result, err = sd.QuerySubject().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (t *Timestamp) Signature(ctx context.Context) (*Signature, error) {
	result, err := t.Edges.SignatureOrErr()
	if IsNotLoaded(err) {
		result, err = t.QuerySignature().Only(ctx)
	}
	return result, MaskNotFound(err)
}
