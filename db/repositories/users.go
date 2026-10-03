package db

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u *UserRepositoryImpl) Create() error { // now UserRepositoryImpl is implementing the UserRepository interface
	fmt.Println("create user called in repository")
	return nil
}