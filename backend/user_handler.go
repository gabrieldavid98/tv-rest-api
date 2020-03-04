package backend

import (
	"net/http"
	"tv-rest-api/entities"

	validator "github.com/go-playground/validator/v10"
)

// UserHandler describes how user controller will behave
type UserHandler interface {
	SaveUser(w http.ResponseWriter, r *http.Request)
}

// userHandler is a concrete implementation of UserHandler interface
type userHandler struct {
	connString string
}

// SaveUser handles the request for the end point [POST] /user
func (u *userHandler) SaveUser(w http.ResponseWriter, r *http.Request) {
	// TODO: db connection
	var user entities.User
	if err := JSON.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("The data sent is not valid"))
		return
	}

	validate := validator.New()
	validate.Struct(user)
}

// NewUserHandler creates a new user handler instance
func NewUserHandler(connString string) UserHandler {
	return &userHandler{
		connString,
	}
}
