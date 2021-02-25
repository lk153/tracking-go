package handler

import (
	"encoding/json"
	"net/http"
)

func responseSuccess(resp http.ResponseWriter, req *http.Request, data interface{}) error {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	return json.NewEncoder(resp).Encode(data)
}

func responseError(resp http.ResponseWriter, req *http.Request, errorMessage string) error {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusInternalServerError)
	return json.NewEncoder(resp).Encode(map[string]interface{}{"error": errorMessage})
}
