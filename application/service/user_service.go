package service

import (
	"net/http"

	"github.com/bagasjs/diy-base/application/model"
	"github.com/bagasjs/diy-base/application/repository"
)

type UserService struct {
    userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
    return &UserService{
        userRepository: repository,
    }
}

func (service *UserService) List() ([]model.GeneralUserModel, *model.Error) {
    users, err:= service.userRepository.ListAll()
    if err != nil {
        return []model.GeneralUserModel{}, err
    }

    usersModels := []model.GeneralUserModel{}
    for _, user := range users {
        userModel := model.GeneralUserModel{
            ID: user.ID,
            Name: user.Name,
            Email: user.Email,
        }
        usersModels = append(usersModels, userModel)
    }
    return usersModels, nil
}

func (service *UserService) Find(id int) (model.GeneralUserModel, *model.Error) {
    return model.GeneralUserModel{}, model.NewError(http.StatusNotImplemented, "Unimplemented UserService.Find()")
}

func (service *UserService) Create(input model.CreateUserModel) *model.Error {
    return model.NewError(http.StatusNotImplemented, "Unimplemented UserService.Create()")
}

func (service *UserService) Update(id int, data model.UpdateUserModel) *model.Error {
    return model.NewError(http.StatusNotImplemented, "Unimplemented UserService.Update()")
}

func (service *UserService) Destroy(id int) *model.Error {
    return model.NewError(http.StatusNotImplemented, "Unimplemented UserService.Destroy()")
}

