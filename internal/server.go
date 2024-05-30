package internal

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/internal/controller"
	"klammerAeffchen/pkg"
	"strconv"
)

// Initializes the webserver that runs the whole web application
func InitializeWebServer(config configuration.Config, authChannel chan *pkg.ShortAuthMessage) {
	app := fiber.New()
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Locals("configuration", config)
		ctx.Locals("auth", authChannel)
		return ctx.Next()
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(controller.ApplicationWebsocket))
	app.Get("/login", controller.GetOAuthURLForBot)
	app.Post("/api/uploadAudio", controller.UploadAudio)

	listenDef := ":" + strconv.Itoa(int(config.ServerPort))
	err := app.Listen(listenDef)
	if err != nil {
		panic(err)
	}

}
