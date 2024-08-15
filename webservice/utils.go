package webservice

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func serveMessageResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	r, _ := json.MarshalIndent(errorResponse{Code: status, Message: message}, "", " ")
	w.Write(r)
}

func serveJsonResponse(w http.ResponseWriter, value interface{}) {
	response, _ := json.MarshalIndent(value, "", " ")
	w.Write(response)
}
