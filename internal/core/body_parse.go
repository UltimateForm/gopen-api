package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

func GetQuery(query url.Values, key string, def string) (string, bool) {
	val := query.Get(key)
	if val == "" {
		return "", false
	}
	return val, true
}

func GetIntQuery(query url.Values, key string, def int) (int, error) {
	stringVal, exists := GetQuery(query, key, "")
	if !exists {
		return def, nil
	}
	numericVal, convErr := strconv.Atoi(stringVal)
	if convErr != nil {
		return 0, BadRequest(fmt.Sprintf("expected '%v' numeric query parameter but instead found string '%v'", key, stringVal))
	}
	return numericVal, nil
}
