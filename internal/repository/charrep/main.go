package charrep

import (
	"context"
	"errors"
	"log"

	"github.com/UltimateForm/gopen-api/internal/repository"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func txReadOneCharacter(ctx context.Context, id string) neo4j.ManagedTransactionWorkT[Character] {
	query := repository.NewQuery(map[string]any{"id": id},
		`
MATCH (character:Character {id: $id})
RETURN character
	`)
	return func(tx neo4j.ManagedTransaction) (Character, error) {
		res, err := query.Execute(tx, ctx)
		if err != nil {
			return Character{}, err
		}

		record, singleErr := res.Single(ctx)
		if singleErr != nil {
			return Character{}, singleErr
		}

		rowMap, rowMapOk := record.AsMap()["character"]
		rowCharacter, rowCharacterOk := rowMap.(neo4j.Node)
		if !rowMapOk || !rowCharacterOk {
			return Character{}, errors.New("bad row from db")
		}

		propsMap := rowCharacter.GetProperties()
		name := propsMap["name"]
		id := propsMap["id"]
		description := propsMap["description"]
		debut := propsMap["debut"]
		nameStr := name.(string)
		idStr := id.(string)
		debutNum := int(debut.(float64))
		descrStr := description.(string)

		return Character{
			Name:        nameStr,
			Id:          idStr,
			Debut:       debutNum,
			Description: descrStr,
		}, nil
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
			rowMap, rowMapOk := rec.AsMap()["character"]
			rowCharacter, rowCharacterOk := rowMap.(neo4j.Node)
			if !rowMapOk || !rowCharacterOk {
				log.Printf("Bad row from db %+v\n", rec)
				i++
				continue
			}
			propsMap := rowCharacter.GetProperties()
			name := propsMap["name"]
			id := propsMap["id"]
			description := propsMap["description"]
			debut := propsMap["debut"]
			nameStr := name.(string)
			idStr := id.(string)
			debutNum := int(debut.(float64))
			descrStr := description.(string)
			characters[i] = Character{
				Name:        nameStr,
				Id:          idStr,
				Debut:       debutNum,
				Description: descrStr,
			}
			i++
		}
		return characters, nil
	}
}

func ReadCharacters(ctx context.Context, offset Offset) ([]Character, error) {
	return repository.ExecuteReadWithDriver(ctx, txReadCharacters(ctx, offset))
}

func ReadOneCharacter(ctx context.Context, id string) (Character, error) {
	return repository.ExecuteReadWithDriver(ctx, txReadOneCharacter(ctx, id))
}
