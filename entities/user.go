package entities

// User represents a user of a system
type User struct {
	FullName       string `json:"fullName" validate:"required"`
	Identification string `json:"id" validate:"required"`
	BirthDate      string `json:"birthDate"`
}
