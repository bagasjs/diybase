package controllers

import (
	"net/http"

	"github.com/bagasjs/diybase/app/model"
	"github.com/bagasjs/diybase/app/service"
	"github.com/gofiber/fiber/v2"
)

type UserAPIController struct {
    userService *service.UserService
}

func (controller *UserAPIController) allUsers(c *fiber.Ctx) error {
    users, error := controller.userService.List()
    if error != nil {
        return c.JSON(fiber.Map {
            "message" : error.Message,
            "code" : error.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map {
        "message" : "Users fetched",
        "code" : http.StatusOK,
        "data" : users,
    })
}

func (controller *UserAPIController) viewUser(c *fiber.Ctx) error {
    userID, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map {
            "message" : "Parameter ID should type of integer",
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    user, error := controller.userService.Find(userID)
    if error != nil {
        return c.JSON(fiber.Map {
            "message" : error.Message,
            "code" : error.Code,
            "data" : nil,
        })
    }

    return c.JSON(user)
}

func (controller *UserAPIController) storeUser(c *fiber.Ctx) error {
    request := model.CreateUpdateUserRequest{}
    if err := c.BodyParser(&request); err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }
    err := controller.userService.Create(request)
    if err != nil {
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Creating user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *UserAPIController) updateUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    request := model.CreateUpdateUserRequest{}
    if err := c.BodyParser(&request); err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    if err := controller.userService.Update(id, request); err != nil { 
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Updating user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *UserAPIController) destroyUser(c *fiber.Ctx) error {
    id, err := c.ParamsInt("id")
    if err != nil {
        return c.JSON(fiber.Map{
            "message" : err.Error(),
            "code" : http.StatusBadRequest,
            "data" : nil,
        })
    }

    if err := controller.userService.Destroy(id); err != nil { 
        return c.JSON(fiber.Map {
            "message" : err.Message,
            "code" : err.Code,
            "data" : nil,
        })
    }

    return c.JSON(fiber.Map{
        "message" : "Destroying user successful",
        "code" : http.StatusOK,
        "data" : nil,
    })
}

func (controller *UserAPIController) Route(r fiber.Router) {
    r.Get("/", controller.allUsers)
    r.Post("/", controller.storeUser)
    r.Get("/:id", controller.viewUser)
    r.Put("/:id", controller.updateUser)
    r.Delete("/:id", controller.destroyUser)
}

func NewUsersAPIController(userService *service.UserService) *UserAPIController {
    return &UserAPIController {
        userService: userService,
    }
}
