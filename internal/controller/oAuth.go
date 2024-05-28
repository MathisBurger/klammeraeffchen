package controller

import (
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/configuration"
)

func GetOAuthURLForBot(ctx *fiber.Ctx) error {

	/*apiUrl := "https://discord.com"
	resource := ""
	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://localhost:5173/authWithCode")
	data.Set("code", code)

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)*/
	configRaw := ctx.Locals("configuration")
	config, _ := configRaw.(configuration.Config)
	return ctx.Redirect(config.OAuthUrl)
}
