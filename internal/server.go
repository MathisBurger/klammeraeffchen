package internal

import (
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/internal/controller"
	"strconv"
)

func InitializeWebServer(config configuration.Config) {
	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("configuration", config)
		return ctx.Next()
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/login", controller.GetOAuthURLForBot)
	listenDef := ":" + strconv.Itoa(int(config.ServerPort))
	err := app.Listen(listenDef)
	if err != nil {
		panic(err)
	}

}
