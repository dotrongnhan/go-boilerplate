package usecases

import (
	"context"

	"git.savvycom.vn/go-services/go-boilerplate/internal/app/models"
	"git.savvycom.vn/go-services/go-boilerplate/internal/domain/users/entities"
	"git.savvycom.vn/go-services/go-boilerplate/internal/domain/users/repositories"
	"github.com/cockroachdb/errors"
)

// UserUseCase is the interface for the use case that managing users.
type UserUseCase interface {
	CreateUser(ctx context.Context, input *models.CreateUserInput) (*models.CreateUserOutput, error)
}

// UserUsecases represents an usecase for managing users.
type userUsecase struct {
	userRepo repositories.UserRepository
}

// NewUserUsecase creates a new instance of the UserUsecase.
func NewUserUsecase(repo repositories.UserRepository) UserUseCase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (c *userUsecase) CreateUser(ctx context.Context, input *models.CreateUserInput) (*models.CreateUserOutput, error) {
	if input.Username == "" {
		return nil, errors.New("username is required")
	}
	if input.Password == "" {
		return nil, errors.New("password is required")
	}
	if input.Email == "" {
		return nil, errors.New("email is required")
	}
	// Hash the password.
	//passwordHash, err := s.cryptoRepo.HashPassword(input.Password)
	//if err != nil {
	//	return nil, err
	//}
	passwordHash := "hashedpassword"

	// Create the user.
	user := &entities.User{
		Username: input.Username,
		Password: passwordHash,
		Email:    input.Email,
	}
	userId, err := c.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &models.CreateUserOutput{
		UserID: userId,
	}, nil
}
