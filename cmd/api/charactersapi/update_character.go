package charactersapi

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
	"github.com/go-chi/chi/v5"
)

func HandleUpdateCharacter(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		core.RespondError(res, req, core.BadRequest("id required"))
		return
	}
	var bodyChar CharacterInfo
	bodyErr := core.ParseBody(req, &bodyChar)
	if bodyErr != nil {
		core.RespondError(res, req, bodyErr)
		return
	}
	updatedChar, updateErr := charrep.UpdateOneCharacter(req.Context(), charrep.Character{
		Id:          id,
		Name:        bodyChar.Name,
		Description: bodyChar.Description,
		Debut:       bodyChar.Debut,
	})
	if updateErr != nil {
		core.RespondError(res, req, updateErr)
		return
	}
	core.RespondOk(res, Character{
		Id: updatedChar.Id,
		CharacterInfo: CharacterInfo{
			Name:        updatedChar.Name,
			Description: updatedChar.Description,
			Debut:       updatedChar.Debut,
		},
	})
}
