package models

// CreateUserInput is the input for the CreateUser use case.
type CreateUserInput struct {
	Username string
	Password string
	Email    string
}

// CreateUserOutput is the output for the CreateUser use case.
type CreateUserOutput struct {
	UserID int64
}
