package charactersapi

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
	"github.com/go-chi/chi/v5"
)

func HandleGetCharacter(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		core.RespondError(res, req, core.BadRequest("id in path is required"))
		return
	}
	character, err := charrep.ReadOneCharacter(req.Context(), id)
	if err != nil {
		core.RespondError(res, req, err)
		return
	}
	var char Character = Character{
		Id: character.Id,
		CharacterInfo: CharacterInfo{
			Name:        character.Name,
			Description: character.Description,
			Debut:       character.Debut,
		},
	}
	core.RespondOk(res, char)
}
