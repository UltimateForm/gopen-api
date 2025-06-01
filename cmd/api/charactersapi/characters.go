package charactersapi

import (
	"errors"
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
)

func HandleGetCharacters(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	limit, limitErr := core.GetIntQuery(query, "limit", 10)
	skip, skipErr := core.GetIntQuery(query, "skip", 0)
	if queryErr := errors.Join(limitErr, skipErr); queryErr != nil {
		core.RespondError(res, req, core.BadRequest(queryErr.Error()))
		return
	}
	characters, err := charrep.ReadCharacters(req.Context(), charrep.Offset{
		Limit: limit,
		Skip:  skip,
	})
	if err != nil {
		core.RespondError(res, req, err)
		return
	}
	charactersMapped := make([]Character, len(characters))
	// NOTE: eh should i be mapping at all? golang type integrity does mean i dont need to worry so much about leakage
	for i, char := range characters {
		charactersMapped[i] = Character{
			Id:          char.Id,
			Name:        char.Name,
			Description: char.Description,
			Debut:       char.Debut,
		}
	}
	core.RespondOk(
		res, CharacterList{
			Offset: Offset{
				Limit: limit,
				Skip:  skip,
			},
			Characters: charactersMapped,
		})
}
