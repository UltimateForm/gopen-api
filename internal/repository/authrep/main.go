package authrep

import (
	"context"
	"errors"

	"github.com/UltimateForm/gopen-api/internal/repository"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func txUsrExistsAtomically(ctx context.Context, user User) neo4j.ManagedTransactionWorkT[bool] {
	query := repository.NewQueryWithParams(
		user,
		`
MATCH (u:User {email:$email, password:$password})
RETURN count(u) > 0 as exists
		`,
	)

	return func(tx neo4j.ManagedTransaction) (bool, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return false, err
		}
		record, singleErr := res.Single(ctx)
		if singleErr != nil {
			return false, singleErr
		}
		exists, found := record.Get("exists")
		if !found {
			return false, nil
		}
		boolValue, isBoolValue := exists.(bool)
		if !isBoolValue {
			return false, errors.New("expected 'exists' bool value from DB row")
		}
		return boolValue, nil
	}

}

func UserExistsAtomically(ctx context.Context, user User) (bool, error) {
	return repository.ExecuteReadWithDriver(ctx, txUsrExistsAtomically(ctx, user))
}
