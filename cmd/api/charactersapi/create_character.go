package charactersapi

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
	"github.com/UltimateForm/gopen-api/internal/util"
)

func HandleCreateCharacter(res http.ResponseWriter, req *http.Request) {
	var bodyChar CharacterInfo
	bodyErr := core.ParseBody(req, &bodyChar)
	if bodyErr != nil {
		core.RespondError(res, req, bodyErr)
		return
	}
	id, idError := util.RandomHex(11)
	if idError != nil {
		core.RespondError(res, req, idError)
		return
	}
	createChar, createErr := charrep.WriteOneCharacter(req.Context(), charrep.Character{
		Id:          id,
		Name:        bodyChar.Name,
		Description: bodyChar.Description,
		Debut:       bodyChar.Debut,
	})
	if createErr != nil {
		core.RespondError(res, req, createErr)
		return
	}
	core.RespondCreated(res, Character{
		Id: createChar.Id,
		CharacterInfo: CharacterInfo{
			Name:        createChar.Name,
			Description: createChar.Description,
			Debut:       createChar.Debut,
		},
	})
}
