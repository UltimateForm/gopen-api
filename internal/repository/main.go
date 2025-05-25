package repository

import (
	"context"
	"time"

	"github.com/UltimateForm/gopen-api/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ExecuteReadWithDriver[T any](ctx context.Context, tx neo4j.ManagedTransactionWorkT[T]) (T, error) {
	readCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*10))
	defer cancel()
	session := db.Driver.NewSession(readCtx, neo4j.SessionConfig{})
	defer session.Close(readCtx)
	return neo4j.ExecuteRead(readCtx, session, tx)
}

func NewQuery(data IQueryParams, query string) QueryRunner {
	return func(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error) {
		return tx.Run(
			ctx,
			query,
			data.ToParams())
	}
}
