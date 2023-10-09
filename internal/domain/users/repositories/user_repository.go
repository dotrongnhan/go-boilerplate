package repositories

import (
	"context"
	"git.savvycom.vn/go-services/go-boilerplate/internal/domain/users/entities"
)

// UserRepository is the repository for users.
type UserRepository interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *entities.User) (int64, error)
	// UpdateUser updates an existing user.
	UpdateUser(ctx context.Context, user *entities.User) error
	// DeleteUser deletes an existing user.
	DeleteUser(ctx context.Context, userID int64) error
	// GetUser gets an existing user by ID.
	GetUser(ctx context.Context, userID int64) (*entities.User, error)
}
