package core

import (
	"encoding/json"
	"net/http"
)

type RequestBody interface {
	Validate() error
}

func ParseBody(req *http.Request, data any) error {
	readErr := json.NewDecoder(req.Body).Decode(data)
	if readErr != nil {
		return readErr
	}

	if dataAsReqBody := data.(RequestBody); dataAsReqBody != nil {
		return dataAsReqBody.Validate()
	}
	return nil
}
