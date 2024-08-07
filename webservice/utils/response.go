package utils

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ServeMessageResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	r, _ := json.MarshalIndent(errorResponse{Code: status, Message: message}, "", " ")
	w.Write(r)
}

func ServeJsonResponse(w http.ResponseWriter, value interface{}) {
	response, _ := json.MarshalIndent(value, "", " ")
	w.Write(response)
}

func ValidateMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method == method {
		return true
	}

	w.WriteHeader(http.StatusNotFound)
	res, _ := json.MarshalIndent(errorResponse{Code: http.StatusNotFound, Message: "URL not found"}, "", " ")
	w.Write(res)
	return false
}
