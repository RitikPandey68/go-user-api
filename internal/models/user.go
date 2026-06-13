package models

// CreateUserRequest is the expected request body for creating a new user.
// Validation rules:
//   - Name: required
//   - DOB:  required, must be in YYYY-MM-DD format
type CreateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required,dob_format"`
}

// UpdateUserRequest is the expected request body for updating an existing user.
// Validation rules:
//   - Name: required
//   - DOB:  required, must be in YYYY-MM-DD format
type UpdateUserRequest struct {
	Name string `json:"name" validate:"required"`
	DOB  string `json:"dob"  validate:"required,dob_format"`
}

// UserResponse is the shape returned to the client for GET and List operations.
// Age is calculated dynamically at request time and is never stored in the DB.
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}

// CreateUserResponse is the shape returned after a successful POST /users.
// Age is intentionally omitted from create response (as per spec).
type CreateUserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
}

// ErrorResponse is the standard error envelope returned on 4xx/5xx responses.
type ErrorResponse struct {
	Error string `json:"error"`
}
