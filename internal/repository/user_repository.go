package repository

import (
	"context"
	"time"

	db "github.com/yourusername/go-user-api/db/sqlc"
)

// UserRepository defines the data access contract for the users table.
// All database interaction happens exclusively through SQLC-generated queries.
type UserRepository interface {
	CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error)
	GetUser(ctx context.Context, id int32) (db.User, error)
	ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error)
	UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (db.User, error)
	DeleteUser(ctx context.Context, id int32) error
}

// userRepository is the concrete implementation backed by SQLC Queries.
type userRepository struct {
	queries *db.Queries
}

// NewUserRepository returns a UserRepository that delegates all calls to SQLC.
func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{queries: queries}
}

// CreateUser inserts a new user record and returns the created row.
func (r *userRepository) CreateUser(ctx context.Context, name string, dob time.Time) (db.User, error) {
	return r.queries.CreateUser(ctx, db.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

// GetUser retrieves a single user by primary key.
func (r *userRepository) GetUser(ctx context.Context, id int32) (db.User, error) {
	return r.queries.GetUser(ctx, id)
}

// ListUsers returns a paginated slice of users ordered by id.
func (r *userRepository) ListUsers(ctx context.Context, limit, offset int32) ([]db.User, error) {
	return r.queries.ListUsers(ctx, db.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
}

// UpdateUser modifies name and dob for the given user id and returns the updated row.
func (r *userRepository) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (db.User, error) {
	return r.queries.UpdateUser(ctx, db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

// DeleteUser removes the user record identified by id.
func (r *userRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.queries.DeleteUser(ctx, id)
}
