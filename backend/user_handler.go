package backend

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
	"tv-rest-api/entities"

	_ "github.com/lib/pq"

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
	db, err := sql.Open("postgres", u.connString)
	if err != nil {
		internalServerError(w, newGenericErrorResponse("Something went wrong :("))
		return
	}

	defer db.Close()

	var user entities.User
	if err := JSON.NewDecoder(r.Body).Decode(&user); err != nil {
		badRequest(w, newGenericErrorResponse("The data could not be processed"))
		return
	}

	var birthDate interface{}
	if user.BirthDate != "" {
		birthDate, err = time.Parse("02-01-2006", user.BirthDate)
		if err != nil {
			badRequest(w, newGenericErrorResponse("The field birthDate has incorrect format, it should be dd-mm-yyyy"))
			return
		}
	}

	validate := validator.New()
	err = validate.Struct(user)

	if err != nil {
		if err, ok := err.(*validator.InvalidValidationError); ok {
			internalServerError(w, newGenericErrorResponse("Something went wrong :("))
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

	var userid int
	err = db.QueryRow(`
		INSERT INTO Users (full_name, identification, birth_date)
		VALUES ($1, $2, $3) RETURNING id_user
	`, user.FullName, user.Identification, birthDate).Scan(&userid)

	if err != nil {
		internalServerError(w, newGenericErrorResponse("Something went wrong :("))
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
