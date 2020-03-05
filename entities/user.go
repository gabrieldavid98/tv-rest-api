package entities

// User represents a user of a system
type User struct {
	FullName       string `json:"fullName" validate:"required,min=10,max=50"`
	Identification string `json:"identification" validate:"required,min=10,max=15"`
	BirthDate      string `json:"birthDate"`
}
