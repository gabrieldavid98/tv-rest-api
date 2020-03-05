package backend

import (
	"net/http"
)

type genericResponse struct {
	Message string `json:"msg"`
}

func newGenericResponse(msg string) *genericResponse {
	return &genericResponse{Message: msg}
}

func response(statusCode int, w http.ResponseWriter, data interface{}) {
	response, err := JSON.Marshal(data)
	if err != nil {
		internalServerError(w, newGenericErrorResponse(err.Error()))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(response)
}

func badRequest(w http.ResponseWriter, data interface{}) {
	response(http.StatusBadRequest, w, data)
}

func internalServerError(w http.ResponseWriter, data interface{}) {
	response(http.StatusInternalServerError, w, data)
}

func ok(w http.ResponseWriter, data interface{}) {
	response(http.StatusOK, w, data)
}
