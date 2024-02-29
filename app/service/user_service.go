package service

import (
	"net/http"

	"github.com/bagasjs/diy-base/app/model"
	"github.com/bagasjs/diy-base/app/repository"
)

type UserService struct {
    Repository *repository.UserRepository
}

func (service *UserService) List() ([]model.GeneralUserModel, *model.Error) {
    return []model.GeneralUserModel{}, model.NewServerError(http.StatusInternalServerError, "Unimplemented UserService.List()")
}

func (service *UserService) Find(id int) (model.GeneralUserModel, *model.Error) {
    return model.GeneralUserModel{}, model.NewServerError(http.StatusInternalServerError, "Unimplemented UserService.Find()")
}

func (service *UserService) Create(input model.CreateUserModel) *model.Error {
    return model.NewServerError(http.StatusInternalServerError, "Unimplemented UserService.Create()")
}

func (service *UserService) Update(id int, data model.UpdateUserModel) *model.Error {
    return model.NewServerError(http.StatusInternalServerError, "Unimplemented UserService.Update()")
}

func (service *UserService) Destroy(id int) *model.Error {
    return model.NewServerError(http.StatusInternalServerError, "Unimplemented UserService.Destroy()")
}

