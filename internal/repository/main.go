package repository

import (
	"context"
	"time"

	"github.com/UltimateForm/gopen-api/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ExecuteReadWithDriver[T any](ctx context.Context, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	session := db.Driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	return neo4j.ExecuteRead(ctx, session, tx, func(config *neo4j.TransactionConfig) {
		config.Timeout = time.Second * 10
	})
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
