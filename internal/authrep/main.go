package authrep

import (
	"context"

	"github.com/UltimateForm/gopen-api/internal/db"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func UserExistsAtomically(ctx context.Context, user User) (bool, error) {
	driver, driverError := db.GetDriver()
	if driverError != nil {
		return false, driverError
	}
	usr, queryErr := neo4j.ExecuteQuery(
		ctx,
		driver,
		"MATCH (u:User {email:$email, password:$password}) RETURN count(u) > 0 as exists",
		map[string]any{"email": user.Email, "password": user.Password},
		neo4j.EagerResultTransformer)
	if queryErr != nil {
		return false, queryErr
	}
	exists, found := usr.Records[0].Get("exists")
	if !found {
		// TODO: add handling here
		return false, nil
	}
	return exists.(bool), nil
}
