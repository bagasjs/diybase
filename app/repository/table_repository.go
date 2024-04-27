package repository

type TableRepository interface {
    ListAll() ([]entity.User, *core.Error)

    // Core Repository Methods
    Query(qb *core.QueryBuilder)([]entity.User, *core.Error)
    Insert(user entity.User) *core.Error 
    Update(id int, user entity.User) *core.Error 
    Destroy(id int) *core.Error 
}
