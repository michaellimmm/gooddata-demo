package adapter

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Errors string `json:"errors"`
}

func (e *ErrorResponse) ToJson() []byte {
	result, _ := json.Marshal(e)
	return result
}

func (e *ErrorResponse) ToJsonString() string {
	result, _ := json.Marshal(e)
	return string(result)
}

func (e *ErrorResponse) WriteResponse(w http.ResponseWriter, statusCode int) {
	http.Error(w, e.ToJsonString(), statusCode)
}
