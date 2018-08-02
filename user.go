package gosample

import (
	"context"
)

// User maps data for user model
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UserRepository is a contract about user action
type UserRepository interface {
	Create(context.Context, User) error
	FindAll(context.Context) ([]User, error)
	FindByID(context.Context, int) (User, error)
}

// UserService holds any user service related
type UserService struct {
	userRepository UserRepository
}

// NewUserService initializes UserService instance
func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{userRepository}
}

// Create is a method to create new user
func (us *UserService) Create(ctx context.Context, user User) error {
	return us.userRepository.Create(ctx, user)
}

// FindAll is a method to fetch all user record
func (us *UserService) FindAll(ctx context.Context) ([]User, error) {
	return us.userRepository.FindAll(ctx)
}

// FindByID is a method to fetch single user record by ID
func (us *UserService) FindByID(ctx context.Context, id int) (User, error) {
	return us.userRepository.FindByID(ctx, id)
}
