package service

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/yourusername/go-user-api/internal/models"
	"github.com/yourusername/go-user-api/internal/repository"
)

// UserService defines the business-logic contract for user operations.
type UserService interface {
	CreateUser(ctx context.Context, req models.CreateUserRequest) (models.CreateUserResponse, error)
	GetUser(ctx context.Context, id int32) (models.UserResponse, error)
	ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error)
	UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.CreateUserResponse, error)
	DeleteUser(ctx context.Context, id int32) error
}

// userService is the concrete implementation of UserService.
type userService struct {
	repo   repository.UserRepository
	logger *zap.Logger
}

// NewUserService creates a UserService with the given repository and logger.
func NewUserService(repo repository.UserRepository, logger *zap.Logger) UserService {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}

// CalculateAge calculates the age in whole years from a date of birth to today.
// It correctly handles the case where the birthday has not yet occurred this year.
func CalculateAge(dob time.Time) int {
	now := time.Now()
	years := now.Year() - dob.Year()

	// If the birthday hasn't happened yet this calendar year, subtract one year
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}
	return years
}

// CreateUser validates, creates a user in the DB, and logs the event.
func (s *userService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.CreateUserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		s.logger.Error("invalid dob format", zap.String("dob", req.DOB), zap.Error(err))
		return models.CreateUserResponse{}, err
	}

	user, err := s.repo.CreateUser(ctx, req.Name, dob)
	if err != nil {
		s.logger.Error("database error on CreateUser", zap.Error(err))
		return models.CreateUserResponse{}, err
	}

	s.logger.Info("User created",
		zap.Int32("id", user.ID),
		zap.String("name", user.Name),
	)

	return models.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

// GetUser fetches a user by ID and attaches the dynamically calculated age.
func (s *userService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		s.logger.Error("database error on GetUser",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
		Age:  CalculateAge(user.Dob),
	}, nil
}

// ListUsers returns a page of users, each with dynamically calculated age.
func (s *userService) ListUsers(ctx context.Context, page, limit int32) ([]models.UserResponse, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	users, err := s.repo.ListUsers(ctx, limit, offset)
	if err != nil {
		s.logger.Error("database error on ListUsers", zap.Error(err))
		return nil, err
	}

	responses := make([]models.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.Dob.Format("2006-01-02"),
			Age:  CalculateAge(u.Dob),
		})
	}
	return responses, nil
}

// UpdateUser modifies an existing user's name and dob, then logs the event.
func (s *userService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.CreateUserResponse, error) {
	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		s.logger.Error("invalid dob format on UpdateUser",
			zap.Int32("id", id),
			zap.String("dob", req.DOB),
			zap.Error(err),
		)
		return models.CreateUserResponse{}, err
	}

	user, err := s.repo.UpdateUser(ctx, id, req.Name, dob)
	if err != nil {
		s.logger.Error("database error on UpdateUser",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return models.CreateUserResponse{}, err
	}

	s.logger.Info("User updated",
		zap.Int32("id", user.ID),
		zap.String("name", user.Name),
	)

	return models.CreateUserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Format("2006-01-02"),
	}, nil
}

// DeleteUser removes a user by ID and logs the event.
func (s *userService) DeleteUser(ctx context.Context, id int32) error {
	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		s.logger.Error("database error on DeleteUser",
			zap.Int32("id", id),
			zap.Error(err),
		)
		return err
	}

	s.logger.Info("User deleted", zap.Int32("id", id))
	return nil
}
