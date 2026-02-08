package meta

import (
	"net/http"

	"github.com/UltimateForm/gopen-api/internal/core"
)

func HandleGetHealth(res http.ResponseWriter, req *http.Request) {
	core.RespondOk(res, map[string]string{"status": "healthy"})
}
