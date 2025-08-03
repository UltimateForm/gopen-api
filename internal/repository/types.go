package repository

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type QueryRunner func(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error)

func (src QueryRunner) Execute(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error) {
	return src(tx, ctx)
}

type IQueryParams interface {
	ToParams() map[string]any
}
type RepoErrorCode int

const (
	UnknownRepoError RepoErrorCode = iota
	EmptyCollectionRepoError
	ValidationError
)

type RepoError struct {
	Code RepoErrorCode
}

func (src RepoError) Error() string {
	switch src.Code {
	case EmptyCollectionRepoError:
		return "Empty collection"
	case ValidationError:
		return "Conflicting records"
	case UnknownRepoError:
		fallthrough
	default:
		return "Unknown repo error"
	}
}
