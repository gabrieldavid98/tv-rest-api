package backend

import (
	"fmt"
	"log"
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
	addr := fmt.Sprintf(":%d", h.port)

	router.Path("/user").
		Methods(http.MethodPost).
		HandlerFunc(userHandler.SaveUser)

	log.Printf("Http server running at %s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

// NewHTTPBackend creates a new http backend
func NewHTTPBackend(connString string, port int) Backend {
	return &httpBackend{
		connString,
		port,
	}
}
