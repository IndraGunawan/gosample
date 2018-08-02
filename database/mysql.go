package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// MySQL holds the functionality to database related
type MySQL struct {
	db *sql.DB
}

// Option holds the requirement to create database connection
type Option struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Charset  string
}

// New initializes MySQL instance
func New(opt Option) (*MySQL, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", opt.User, opt.Password, opt.Host, opt.Port, opt.Database, opt.Charset)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return &MySQL{}, err
	}

	return &MySQL{db: db}, nil
}
