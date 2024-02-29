package repository

import (
	"database/sql"
	"net/http"

	"github.com/bagasjs/diy-base/application/entity"
	"github.com/bagasjs/diy-base/application/model"
)


type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) *model.Error {
    stmt, err := repo.db.Prepare("INSERT INTO users (`name`, `email`, `password`) VALUES (?, ?, ?)")
    if err != nil {
        return model.NewError(http.StatusNotImplemented, err.Error())
    }
    defer stmt.Close()
    _, err = stmt.Exec(user.Name, user.Email, user.Password)
    if err != nil {
        return model.NewError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) ListAll() ([]entity.User, *model.Error) {
    rows, err := repo.db.Query("SELECT * FROM users")
    if err != nil {
        return nil, model.NewError(http.StatusNotImplemented, err.Error())
    }
    defer rows.Close()
    users := []entity.User{}
    for rows.Next() {
        user := entity.User{}
        err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
        if err != nil {
            return nil, model.NewError(http.StatusNotImplemented, err.Error())
        }
        users = append(users, user)
    }
    return users, nil
}

func (repo *userRepositoryImplSQLite3) Find(id int) (entity.User, *model.Error) {
    return entity.User{}, model.NewError(http.StatusNotImplemented, 
        "Unimplemented userRepositoryImplSQLite3.Find()")
}

func (repo *userRepositoryImplSQLite3) Update(id int, user entity.User) *model.Error {
    return model.NewError(http.StatusNotImplemented, 
        "Unimplemented userRepositoryImplSQLite3.Update()")
}

func (repo *userRepositoryImplSQLite3) Destroy(id int) *model.Error {
    return model.NewError(http.StatusNotImplemented, 
        "Unimplemented userRepositoryImplSQLite3.Destroy()")
}

func NewUserRepository(conn *sql.DB) UserRepository {
    return &userRepositoryImplSQLite3 {
        db: conn,
    };
}

