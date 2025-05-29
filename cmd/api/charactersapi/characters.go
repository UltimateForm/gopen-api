package charactersapi

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
)

type Offset struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

func (src Offset) Validate() error {
	if src.Limit <= 0 {
		return core.BadRequest("limit must be a positive higher than 0 integer")
	}
	if src.Skip < 0 {
		return core.BadRequest("skip be a positive integer")
	}
	return nil
}

func HandleGetCharacters(res http.ResponseWriter, req *http.Request) {
	var offset Offset
	// TODO: use query parameters silly
	err := core.ParseBody(req, &offset)
	if err != nil {
		core.RespondError(res, req, err)
		return
	}
	characters, err := charrep.ReadCharacters(req.Context(), charrep.Offset{
		Limit: offset.Limit,
		Skip:  offset.Skip,
	})
	if err != nil {
		core.RespondError(res, req, err)
		return
	}
	charactersMapped := make([]Character, len(characters))
	// NOTE: eh should i be mapping at all, golang type integrity does mean i dont need to worry so much about leakage
	for i, char := range characters {
		charactersMapped[i] = Character{
			Id:          char.Id,
			Name:        char.Name,
			Description: char.Description,
			Debut:       char.Debut,
		}
	}
	core.RespondOk(res, CharacterList{Offset: offset, Characters: charactersMapped})
}
