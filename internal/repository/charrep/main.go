package charrep

import (
	"context"
	"errors"
	"log"

	"github.com/UltimateForm/gopen-api/internal/repository"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func extractCharacterFromRecord(record *neo4j.Record) (Character, error) {
	rowMap, rowMapOk := record.AsMap()["character"]
	rowCharacter, rowCharacterOk := rowMap.(neo4j.Node)
	if !rowMapOk || !rowCharacterOk {
		return Character{}, errors.New("bad row from db")
	}

	propsMap := rowCharacter.GetProperties()
	name := propsMap["name"].(string)
	id := propsMap["id"].(string)
	description := propsMap["description"].(string)
	debutNum := propsMap["debut"]
	var debut int
	// why you ask? mock data i'm using was created badly
	if debutAsInt, isInt64 := debutNum.(int64); isInt64 {
		debut = int(debutAsInt)
	}
	if debutAsFloat, isFloat64 := debutNum.(float64); isFloat64 {
		debut = int(debutAsFloat)
	}
	return Character{
		Name:        name,
		Id:          id,
		Debut:       debut,
		Description: description,
	}, nil
}

func txWriteOneCharacter(ctx context.Context, character Character) neo4j.ManagedTransactionWorkT[Character] {
	query := repository.NewQueryWithParams(character,
		`
CREATE (character:Character {id: $id, name: $name, description: $description, debut: $debut}) 
RETURN character
LIMIT 1
			`)
	return func(tx neo4j.ManagedTransaction) (Character, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return Character{}, err
		}
		record, recordErr := res.Single(ctx)
		if recordErr != nil {
			return Character{}, recordErr
		}
		return extractCharacterFromRecord(record)
	}
}

func txReadOneCharacter(ctx context.Context, id string) neo4j.ManagedTransactionWorkT[Character] {
	query := repository.NewQuery(map[string]any{"id": id},
		`
MATCH (character:Character {id: $id})
RETURN character
LIMIT 1
	`)
	return func(tx neo4j.ManagedTransaction) (Character, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return Character{}, err
		}
		var record *neo4j.Record
		hasNext := res.NextRecord(ctx, &record)
		if !hasNext {
			return Character{}, repository.RepoError{Code: repository.EmptyCollectionRepoError}
		}

		return extractCharacterFromRecord(record)
	}
}

func txReadCharacters(ctx context.Context, offset Offset) neo4j.ManagedTransactionWorkT[[]Character] {
	query := repository.NewQueryWithParams(offset,
		`
MATCH (character:Character)
RETURN character
ORDER BY character.debut
SKIP $skip
LIMIT $limit
	`)
	return func(tx neo4j.ManagedTransaction) ([]Character, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return nil, err
		}
		characters := make([]Character, offset.Limit)
		i := 0
		for rec, yieldErr := range res.Records(ctx) {
			if yieldErr != nil {
				log.Printf("Error while iterating DB records, %v", yieldErr)
				i++
				continue
			}
			character, extractErr := extractCharacterFromRecord(rec)
			if extractErr != nil {
				log.Printf("Bad row from db %+v\n", extractErr)
				i++
				continue
			}
			characters[i] = character
			i++
		}
		return characters, nil
	}
}

func txUpdateOneCharacter(ctx context.Context, character Character) neo4j.ManagedTransactionWorkT[Character] {
	query := repository.NewQueryWithParams(character,
		`
MATCH (character:Character {id: $id})
SET character.name = $name, character.description = $description, character.debut = $debut
RETURN character
	`)
	return func(tx neo4j.ManagedTransaction) (Character, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return Character{}, err
		}
		record, recordErr := res.Single(ctx)
		if recordErr != nil {
			return Character{}, recordErr
		}
		return extractCharacterFromRecord(record)
	}
}

func txDeleteOneCharacter(ctx context.Context, id string) neo4j.ManagedTransactionWorkT[bool] {
	query := repository.NewQuery(map[string]any{"id": id},
		`
MATCH (character:Character {id: $id})
DELETE character
RETURN count(character) > 0 as deleted
	`)
	return func(tx neo4j.ManagedTransaction) (bool, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return false, err
		}
		record, recordErr := res.Single(ctx)
		if recordErr != nil {
			return false, recordErr
		}

		deleted, found := record.Get("deleted")
		if !found {
			return false, errors.New("expected 'deleted' boolean value from DB row")
		}
		boolValue, isBoolValue := deleted.(bool)
		if !isBoolValue {
			return false, errors.New("expected 'deleted' bool value from DB row")
		}
		return boolValue, nil
	}
}

func ReadCharacters(ctx context.Context, offset Offset) ([]Character, error) {
	return repository.ExecuteReadWithDriver(ctx, txReadCharacters(ctx, offset))
}

func ReadOneCharacter(ctx context.Context, id string) (Character, error) {
	return repository.ExecuteReadWithDriver(ctx, txReadOneCharacter(ctx, id))
}

func WriteOneCharacter(ctx context.Context, character Character) (Character, error) {
	return repository.ExecuteWriteWithDriver(ctx, txWriteOneCharacter(ctx, character))
}

func UpdateOneCharacter(ctx context.Context, character Character) (Character, error) {
	return repository.ExecuteWriteWithDriver(ctx, txUpdateOneCharacter(ctx, character))
}

func DeleteOneCharacter(ctx context.Context, id string) (bool, error) {
	return repository.ExecuteWriteWithDriver(ctx, txDeleteOneCharacter(ctx, id))
}
