package backend

import (
	"fmt"
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
		badRequest(w, newGenericErrorResponse("The data could not be processed"))
		return
	}

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			internalServerError(w, newGenericErrorResponse(err.Error()))
			return
		}

		var erros []string

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				erros = append(
					erros,
					fmt.Sprintf("The field %s is required", err.Field()),
				)
				break
			case "min":
				erros = append(
					erros,
					fmt.Sprintf(
						"The field %s is too short, needs at least %s elements/characters",
						err.Field(),
						err.Param(),
					),
				)
				break
			case "max":
				erros = append(
					erros,
					fmt.Sprintf(
						"The field %s is too long, maximun %s elements/characters",
						err.Field(),
						err.Param(),
					),
				)
				break
			default:
				break
			}
		}

		badRequest(w, newGenericErrorResponse(erros...))
		return
	}

	ok(w, newGenericResponse("resource created"))
}

// NewUserHandler creates a new user handler instance
func NewUserHandler(connString string) UserHandler {
	return &userHandler{
		connString,
	}
}
