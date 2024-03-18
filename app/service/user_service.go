package service

import (
	"net/http"
	"strings"

	"github.com/bagasjs/diybase/app/core"
	"github.com/bagasjs/diybase/app/entity"
	"github.com/bagasjs/diybase/app/model"
	"github.com/bagasjs/diybase/app/repository"
)

type UserService struct {
    userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
    return &UserService{
        userRepository: repository,
    }
}

func (service *UserService) List() ([]model.GeneralUserResponse, *core.Error) {
    users, err:= service.userRepository.ListAll()
    if err != nil {
        return []model.GeneralUserResponse{}, err
    }

    usersModels := []model.GeneralUserResponse{}
    for _, user := range users {
        userModel := model.GeneralUserResponse {
            ID: user.ID,
            Name: user.Name,
            Email: user.Email,
            Created: user.Created,
            Updated: user.Updated,
        }
        usersModels = append(usersModels, userModel)
    }
    return usersModels, nil
}

func (service *UserService) Find(id int) (model.GeneralUserResponse, *core.Error) {
    response := model.GeneralUserResponse{}
    query := core.NewQueryBuilder()
    users, err:=  service.userRepository.Query(query.Where("id", "=", id).Limit(1))
    if err != nil {
        return response, err
    }

    if len(users) == 0 {
        return response, core.NewError(http.StatusNotFound, "Failed to find user")
    }

    response.Name = users[0].Name
    response.Email = users[0].Email
    response.ID = users[0].ID
    return response, nil
}

func (service *UserService) Create(request model.CreateUpdateUserRequest) *core.Error {
    if strings.Compare(request.Password, request.PasswordConfirmation) != 0 {
        return core.NewError(http.StatusForbidden, "Password and it's confirmation should be equal")
    }

    entt := entity.User {
        Name: request.Name,
        Password: request.Password,
        Email: request.Email,
    }

    if err := service.userRepository.Insert(entt); err != nil {
        return err
    }

    return nil
}

func (service *UserService) Update(id int, request model.CreateUpdateUserRequest) *core.Error {
    if strings.Compare(request.Password, request.PasswordConfirmation) != 0 {
        return core.NewError(http.StatusForbidden, "Password and it's confirmation should be equal")
    }

    entt := entity.User {
        Name: request.Name,
        Email: request.Email,
        Password: request.Password,
    }
    if err := service.userRepository.Update(id, entt); err != nil {
        return err
    }

    return nil
}

func (service *UserService) Destroy(id int) *core.Error {
    if err := service.userRepository.Destroy(id); err != nil {
        return err
    }
    return nil
}

