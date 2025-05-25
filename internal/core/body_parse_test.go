package core

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

type testBody struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
	Other  string `json:"other"`
}
type testBodyWithValidation struct {
	Text   string `json:"text"`
	Number int    `json:"number"`
	Other  string `json:"other"`
}

func (src *testBodyWithValidation) Validate() error {
	if src.Text == "" {
		return BadRequest("text is required")
	}
	return nil
}
func TestParseBodyOk(t *testing.T) {
	bodyVals := testBody{
		Text:   "lorem",
		Number: 34,
		Other:  "ipsum",
	}
	jsonStr, _ := json.Marshal(bodyVals)
	testRequest := httptest.NewRequest("POST", "http://localhost:3000", bytes.NewBuffer(jsonStr))
	var parsedBody testBody
	parseErr := ParseBody(testRequest, &parsedBody)
	if parseErr != nil {
		t.Fatalf("Expected parseErr to be nil but instead it is %v", parseErr)
	}
	if !reflect.DeepEqual(bodyVals, parsedBody) {
		t.Fatalf("Expected parsedBody and bodVals to be deeply equal but they're not, bodyVals: %+v ; parsedBody: %+v", bodyVals, parsedBody)
	}
}

func TestParseBodyWithValidationFail(t *testing.T) {
	bodyVals := testBodyWithValidation{
		Number: 34,
		Other:  "ipsum",
	}
	jsonStr, _ := json.Marshal(bodyVals)
	testRequest := httptest.NewRequest("POST", "http://localhost:3000", bytes.NewBuffer(jsonStr))
	var parsedBody testBodyWithValidation
	parseErr := ParseBody(testRequest, &parsedBody)
	if parseErr == nil {
		t.Fatalf("Expected parseErr to not be nil but instead it is %v", parseErr)
	}
	badRequestErr, isErrorResponse := parseErr.(*ErrorResponse)
	if !isErrorResponse {
		t.Fatalf("Expected error (%v) to be of type error response but it is not", parseErr)
	}
	errorRes := &ErrorResponse{
		Code:    400,
		Message: "text is required",
	}
	if !reflect.DeepEqual(errorRes, badRequestErr) {
		t.Fatalf("Expected errorRes and badRequestErr to be deeply equal but they're not, badRequestErr: %+v ; errorRes: %+v", badRequestErr, errorRes)
	}
}

func TestParseBodyWithValidationOk(t *testing.T) {
	bodyVals := testBodyWithValidation{
		Text:   "lorem",
		Number: 34,
		Other:  "ipsum",
	}
	jsonStr, _ := json.Marshal(bodyVals)
	testRequest := httptest.NewRequest("POST", "http://localhost:3000", bytes.NewBuffer(jsonStr))
	var parsedBody testBodyWithValidation
	parseErr := ParseBody(testRequest, &parsedBody)
	if parseErr != nil {
		t.Fatalf("Expected parseErr to be nil but instead it is %v", parseErr)
	}
	if !reflect.DeepEqual(bodyVals, parsedBody) {
		t.Fatalf("Expected parsedBody and bodVals to be deeply equal but they're not, bodyVals: %+v ; parsedBody: %+v", bodyVals, parsedBody)
	}
}
