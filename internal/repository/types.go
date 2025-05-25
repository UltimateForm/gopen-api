package repository

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type QueryRunner func(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error)

func (src QueryRunner) Execute(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error) {
	return src(tx, ctx)
}
