package controllers

import (
	"net/http"

	"github.com/bagasjs/diybase/application/service"
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

    return c.JSON(fiber.Map {
        "message" : "Users fetched",
        "code" : http.StatusOK,
        "data" : user,
    })
}

func (controller *UserAPIController) storeUser(c *fiber.Ctx) error {
    return c.SendString("New user has been stored")
}

func (controller *UserAPIController) updateUser(c *fiber.Ctx) error {
    return c.SendString("New user has been created")
}

func (controller *UserAPIController) destroyUser(c *fiber.Ctx) error {
    return c.SendString("You are deleting an user")
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
