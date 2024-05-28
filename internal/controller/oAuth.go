package controller

import (
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/configuration"
)

func GetOAuthURLForBot(ctx *fiber.Ctx) error {
	configRaw := ctx.Locals("configuration")
	config, _ := configRaw.(configuration.Config)
	url := "https://discord.com/oauth2/authorize?client_id=" + config.BotClientID + "&response_type=code&redirect_uri=http%3A%2F%2F" + config.WebHost + "%2FauthWithCode&scope=connections"
	return ctx.Redirect(url)
}
