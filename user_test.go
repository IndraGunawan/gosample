package gosample_test

import (
	"context"
	"testing"

	"github.com/IndraGunawan/gosample"

	"github.com/IndraGunawan/gosample/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	userRepository := &mock.UserRepository{}
	userService := gosample.NewUserService(userRepository)

	assert.NotNil(t, userService)
}

func TestCreate(t *testing.T) {
	userRepository := &mock.UserRepository{}
	userService := gosample.NewUserService(userRepository)
	user := gosample.User{
		Name:     "name1",
		Password: "pass1",
		Email:    "email1",
	}

	lastInsertID, err := userService.Create(context.Background(), user)
	assert.Equal(t, int64(0), lastInsertID)
	assert.Nil(t, err)
}

func TestFindByID(t *testing.T) {
	userRepository := &mock.UserRepository{}
	userService := gosample.NewUserService(userRepository)

	expedtedValue := []gosample.User{
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

	users, err := userService.FindAll(context.Background())
	assert.Equal(t, expedtedValue, users)
	assert.Nil(t, err)
}

func TestFindAll(t *testing.T) {
	userRepository := &mock.UserRepository{}
	userService := gosample.NewUserService(userRepository)

	expedtedValue := gosample.User{
		ID:       1,
		Name:     "test1",
		Email:    "test1@example.com",
		Password: "test1password",
	}

	user, err := userService.FindByID(context.Background(), 1)
	assert.Equal(t, expedtedValue, user)
	assert.Nil(t, err)
}
