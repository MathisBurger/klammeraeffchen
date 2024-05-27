package internal

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func InitializeWebServer(config Config) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	listenDef := ":" + strconv.Itoa(int(config.ServerPort))
	err := app.Listen(listenDef)
	if err != nil {
		panic(err)
	}

}
