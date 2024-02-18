package repository

import (
	"github.com/bagasjs/diy-base/app/entity"
)

type UserRepository interface {
    Insert(user entity.User) error
    FindAll() ([]entity.User, error)
    Update(id int, user entity.User) error
    Delete(id int) error
}
