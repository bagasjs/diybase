package repository

import (
	"github.com/bagasjs/diybase/app/core"
	"github.com/bagasjs/diybase/app/entity"
)

type UserRepository interface {
    ListAll() ([]entity.User, *core.Error)

    // Core Repository Methods
    Query(qb *core.QueryBuilder)([]entity.User, *core.Error)
    Insert(user entity.User) *core.Error 
    Update(id int, user entity.User) *core.Error 
    Destroy(id int) *core.Error 
}
