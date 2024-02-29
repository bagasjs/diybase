package repository

import (
	"github.com/bagasjs/diy-base/app/entity"
	"github.com/bagasjs/diy-base/app/model"
)

type UserRepository interface {
    Insert(user entity.User) *model.Error 
    Update(id int, user entity.User) *model.Error 
    Destroy(id int) *model.Error 
    List() ([]entity.User, *model.Error)
    Find(id int) (entity.User, *model.Error)
}
