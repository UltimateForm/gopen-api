package authrep

import (
	"context"

	"github.com/UltimateForm/gopen-api/internal/repository"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func queryComputeUserExistsAtomically(user User) repository.QueryRunner {
	return func(tx neo4j.ManagedTransaction, ctx context.Context) (neo4j.ResultWithContext, error) {
		return tx.Run(
			ctx,
			"MATCH (u:User {email:$email, password:$password}) RETURN count(u) > 0 as exists",
			map[string]any{"email": user.Email, "password": user.Password})
	}
}
