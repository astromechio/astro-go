package servlib

import (
	"encoding/json"
	"net/http"
)

type ARequest struct {
	ID          string        `json:"id"`
	UserRequest *http.Request `json:"userRequest`
}

func ARequestFromJSON(raw []byte) (*ARequest, error) {
	var newReq ARequest
	err := json.Unmarshal(raw, newReq)

	return &newReq, err
}
