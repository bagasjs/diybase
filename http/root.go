package http

import (
	"log"

	"github.com/bagasjs/diy-base/http/controllers"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
)

func Serve() {
    engine := html.New("./res/views/", ".html")
    app := fiber.New(fiber.Config {
        Views: engine,
    })
    app.Static("/public", "./res/public/")

    app.Get("/admin", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to admin page")
    })

    userAdminController := controllers.UserAdminController{}
    app.Route("/admin/users", userAdminController.Route)

    userAPIController := controllers.UserAPIController{}
    app.Route("/api/users", userAPIController.Route)

    log.Fatal(app.Listen(":8080"))
}
