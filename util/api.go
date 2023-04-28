package util

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
	StatusCode       int    `json:"-"`
	MoreInfo         string `json:"moreInfo"`
}

func ReplyJSON(w http.ResponseWriter, statusCode int, payload interface{}) string {
	w.Header().Set("content-type", "application/json")

	if payload != nil {
		if _, ok := payload.(*APIError); ok {
			w.WriteHeader(payload.(*APIError).StatusCode)
		} else {
			w.WriteHeader(statusCode)
		}
		responseBytes, _ := json.Marshal(payload)
		w.Write(responseBytes)
		return string(responseBytes)
	}

	w.WriteHeader(statusCode)
	return ""
}
