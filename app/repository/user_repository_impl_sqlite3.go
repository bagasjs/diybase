package repository

import (
	"database/sql"
	"net/http"

	"github.com/bagasjs/diybase/app/core"
	"github.com/bagasjs/diybase/app/entity"
)


type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Query(query *core.QueryBuilder) ([]entity.User, *core.Error) {
    query.Table("internal__users")
    rows, err := repo.db.Query(query.ToString(), query.Values()...)
    if err != nil {
        return []entity.User{}, core.NewError(http.StatusInternalServerError, err.Error())
    }
    defer rows.Close()
    users := []entity.User{}
    for rows.Next() {
        user := entity.User{}
        err := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
        if err != nil {
            return []entity.User{}, core.NewError(http.StatusNotImplemented, err.Error())
        }
        users = append(users, user)
    }

    return users, nil
}

func (repo *userRepositoryImplSQLite3) ListAll() ([]entity.User, *core.Error) {
    return repo.Query(core.NewQueryBuilder())
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) *core.Error {
    stmt, err := repo.db.Prepare("INSERT INTO `internal__users` (`name`, `email`, `password`) VALUES (?, ?, ?)")
    if err != nil {
        return core.NewError(http.StatusNotImplemented, err.Error())
    }
    defer stmt.Close()
    _, err = stmt.Exec(user.Name, user.Email, user.Password)
    if err != nil {
        return core.NewError(http.StatusBadRequest, err.Error())
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) Update(id int, user entity.User) *core.Error {
    return core.NewError(http.StatusNotImplemented, 
        "Unimplemented userRepositoryImplSQLite3.Update()")
}

func (repo *userRepositoryImplSQLite3) Destroy(id int) *core.Error {
    return core.NewError(http.StatusNotImplemented, 
        "Unimplemented userRepositoryImplSQLite3.Destroy()")
}

func NewUserRepository(conn *sql.DB) UserRepository {
    return &userRepositoryImplSQLite3 {
        db: conn,
    };
}

