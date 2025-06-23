package charactersapi

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
	"github.com/UltimateForm/gopen-api/internal/repository/charrep"
	"github.com/go-chi/chi/v5"
)

func HandleDeleteCharacter(res http.ResponseWriter, req *http.Request) {
	id := chi.URLParam(req, "id")
	if id == "" {
		core.RespondKnownError(res, core.BadRequest("id required"))
		return
	}
	deleted, deleteError := charrep.DeleteOneCharacter(req.Context(), id)
	if deleteError != nil {
		core.RespondError(res, req, deleteError)
		return
	}
	if !deleted {
		core.RespondKnownError(res, core.NotFound())
		return
	}
	core.RespondNoContent(res)
}
