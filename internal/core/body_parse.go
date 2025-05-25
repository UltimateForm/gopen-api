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

	if dataAsReqBody, isValidate := data.(RequestBody); isValidate {
		return dataAsReqBody.Validate()
	}
	return nil
}
