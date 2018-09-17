package mock

import (
	"context"

	"github.com/IndraGunawan/gosample"
)

// UserRepository holds any user from/to database
type UserRepository struct{}

// Create is method to insert new user to database
func (u *UserRepository) Create(ctx context.Context, user gosample.User) (int64, error) {
	return 0, nil
}

// FindAll fetchs all user record from database
func (u *UserRepository) FindAll(ctx context.Context) ([]gosample.User, error) {
	users := []gosample.User{
		gosample.User{
			ID:       1,
			Name:     "test1",
			Email:    "test1@example.com",
			Password: "test1password",
		},
		gosample.User{
			ID:       2,
			Name:     "test2",
			Email:    "test2@example.com",
			Password: "test2password",
		},
	}

	return users, nil
}

// FindByID fetchs single user record from database by id
func (u *UserRepository) FindByID(ctx context.Context, id int64) (gosample.User, error) {
	user := gosample.User{
		ID:       1,
		Name:     "test1",
		Email:    "test1@example.com",
		Password: "test1password",
	}

	return user, nil
}
