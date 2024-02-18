package repository

import (
	"database/sql"

	"github.com/bagasjs/diy-base/app/entity"
)


type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) error {
    return nil
}

func (repo *userRepositoryImplSQLite3) FindAll() ([]entity.User, error) {
    return []entity.User{}, nil
}

func (repo *userRepositoryImplSQLite3) Update(id int, user entity.User) error {
    return nil
}

func (repo *userRepositoryImplSQLite3) Delete(id int) error {
    return nil
}

