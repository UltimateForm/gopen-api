package repository

import (
	"context"
	"time"

	"github.com/UltimateForm/gopen-api/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func mapNeo4JToRepo(err *neo4j.Neo4jError) RepoError {
	switch err.Code {
	case "Neo.ClientError.Statement.EntityNotFound":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Schema.ConstraintValidationFailed":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Statement.ArgumentError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Statement.InvalidArgument":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.General.InvalidConfiguration":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Statement.SyntaxError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Transaction.InvalidOperation":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Statement.ParameterError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Security.Unauthorized":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Statement.TypeError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Statement.ValueError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.General.NotSupported":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Statement.ArithmeticError":
		return RepoError{Code: ValidationError}
	case "Neo.ClientError.Security.AuthenticationFailed":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Security.AuthorizationExpired":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Security.Forbidden":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Transaction.LockClientStopped":
		return RepoError{Code: UnknownRepoError}
	case "Neo.TransientError.General.DatabaseUnavailable":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Transaction.InvalidState":
		return RepoError{Code: UnknownRepoError}
	case "Neo.TransientError.Transaction.Terminated":
		return RepoError{Code: UnknownRepoError}
	case "Neo.ClientError.Statement.NoSuchEntity":
		return RepoError{Code: EmptyCollectionRepoError}
	default:
		return RepoError{Code: UnknownRepoError}
	}
}

func ExecuteReadWithDriver[T any](ctx context.Context, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := neo4j.ExecuteRead(ctx, session, tx, func(config *neo4j.TransactionConfig) {
		config.Timeout = time.Second * 10
	})
	neo4jError, isNeo4jError := err.(*neo4j.Neo4jError)
	if isNeo4jError {
		return result, mapNeo4JToRepo(neo4jError)
	}
	return result, err
}

func ExecuteWriteWithDriver[T any](ctx context.Context, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	result, err := neo4j.ExecuteWrite(ctx, session, tx, func(config *neo4j.TransactionConfig) {
		config.Timeout = time.Second * 10
	})
	neo4jError, isNeo4jError := err.(*neo4j.Neo4jError)
	if isNeo4jError {
		return result, mapNeo4JToRepo(neo4jError)
	}
	return result, err
}

func NewQuery(data map[string]any, query string) QueryRunner {
	return func(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error) {
		return tx.Run(
			ctx,
			query,
			data)
	}
}

func NewQueryWithParams(data IQueryParams, query string) QueryRunner {
	return NewQuery(data.ToParams(), query)
}
