package repository

import (
	"github.com/bagasjs/diybase/application/entity"
	"github.com/bagasjs/diybase/application/model"
)

type UserRepository interface {
    Insert(user entity.User) *model.Error 
    ListAll() ([]entity.User, *model.Error)
    Find(id int) (entity.User, *model.Error)
    Update(id int, user entity.User) *model.Error 
    Destroy(id int) *model.Error 
}
