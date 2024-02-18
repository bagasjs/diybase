package controllers

import "github.com/gofiber/fiber/v2"

type UserAPIController struct {
}

func (controller *UserAPIController) allUsers(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{})
}

func (controller *UserAPIController) viewUser(c *fiber.Ctx) error {
    return c.SendString("You are viewing an user")
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
