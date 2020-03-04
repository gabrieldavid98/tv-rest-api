package backend

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type httpBackend struct {
	connString string
	port       int
}

func (h *httpBackend) Start() {
	router := mux.NewRouter()
	userHandler := NewUserHandler(h.connString)

	router.Path("/user").
		Methods(http.MethodPost).
		HandlerFunc(userHandler.SaveUser)

	http.ListenAndServe(fmt.Sprintf(":%d", h.port), router)
}

// NewHTTPBackend creates a new http backend
func NewHTTPBackend(connString string, port int) Backend {
	return &httpBackend{
		connString,
		port,
	}
}
