package core

import (
	"encoding/json"
	"net/http"
)

func RespondOk(res http.ResponseWriter, data any) {
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func RespondCreated(res http.ResponseWriter, data any) {
	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Content-Type", "application/json")
	json.NewEncoder(res).Encode(data)
}

func RespondNoContent(res http.ResponseWriter) {
	res.WriteHeader(http.StatusNoContent)
}
