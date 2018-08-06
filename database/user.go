package database

import (
	"context"
	"errors"

	"github.com/IndraGunawan/goq"

	"github.com/IndraGunawan/gosample"
)

// UserRepository holds any user from/to database
type UserRepository struct {
	mysql *MySQL
}

// NewUserRepository initializes UserRepository instance
func NewUserRepository(database *MySQL) *UserRepository {
	return &UserRepository{database}
}

// Create is method to insert new user to database
func (u *UserRepository) Create(ctx context.Context, user gosample.User) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, errors.New("Timeout")
	default:
	}

	result, err := u.mysql.Db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	var lastInsertID int64
	lastInsertID, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

// FindAll fetchs all user record from database
func (u *UserRepository) FindAll(ctx context.Context) ([]gosample.User, error) {
	var users []gosample.User

	select {
	case <-ctx.Done():
		return users, errors.New("Timeout")
	default:
	}

	builder := goq.Select("id", "name", "email", "password").
		From("users")

	rows, err := u.mysql.Db.Query(builder.ToSQL())
	if err != nil {
		return users, nil
	}

	defer rows.Close()
	for rows.Next() {
		var user gosample.User

		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID fetchs single user record from database by id
func (u *UserRepository) FindByID(ctx context.Context, id int64) (gosample.User, error) {
	var user gosample.User

	select {
	case <-ctx.Done():
		return user, errors.New("Timeout")
	default:
	}

	builder := goq.Select("id", "name", "email", "password").
		From("users").
		Where("id = ?", id)

	err := u.mysql.Db.QueryRow(builder.ToSQL(), builder.GetBindingParameters()...).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	return user, err
}
