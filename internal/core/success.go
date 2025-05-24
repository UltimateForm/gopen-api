package core

import (
	"encoding/json"
	"net/http"
)

func RespondOk(res http.ResponseWriter, data any) {
	json.NewEncoder(res).Encode(data)
}

func RespondNoContent(res http.ResponseWriter) {
	res.WriteHeader(http.StatusNoContent)
}
