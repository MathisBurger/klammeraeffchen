package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/configuration"
	"net/url"
)

// Gets the oauth for the bot
func GetOAuthURLForBot(ctx *fiber.Ctx) error {

	configRaw := ctx.Locals("configuration")
	config, _ := configRaw.(configuration.Config)
	authUrl := "https://discord.com"
	resource := "/oauth2/authorize"
	data := url.Values{
		"client_id":     {config.BotClientID},
		"response_type": {"code"},
		"redirect_uri":  {config.OAuthRedirect},
		"scope":         {"identify"},
	}
	u, _ := url.ParseRequestURI(authUrl)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)
	return ctx.Redirect(urlStr)
}
