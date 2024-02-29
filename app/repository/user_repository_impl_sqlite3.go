package repository

import (
	"database/sql"
	"net/http"

	"github.com/bagasjs/diy-base/app/entity"
	"github.com/bagasjs/diy-base/app/model"
)


type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) *model.Error {
    return model.NewError(http.StatusInternalServerError, 
        "Unimplemented userRepositoryImplSQLite3.Insert()")
}

func (repo *userRepositoryImplSQLite3) List() ([]entity.User, *model.Error) {
    return []entity.User{}, model.NewError(http.StatusInternalServerError, 
        "Unimplemented userRepositoryImplSQLite3.List()")
}

func (repo *userRepositoryImplSQLite3) Find(id int) (entity.User, *model.Error) {
    return entity.User{}, model.NewError(http.StatusInternalServerError, 
        "Unimplemented userRepositoryImplSQLite3.Find()")
}

func (repo *userRepositoryImplSQLite3) Update(id int, user entity.User) *model.Error {
    return model.NewError(http.StatusInternalServerError, 
        "Unimplemented userRepositoryImplSQLite3.Update()")
}

func (repo *userRepositoryImplSQLite3) Destroy(id int) *model.Error {
    return model.NewError(http.StatusInternalServerError, 
        "Unimplemented userRepositoryImplSQLite3.Destroy()")
}

