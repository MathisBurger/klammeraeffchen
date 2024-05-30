package controller

import (
	_ "encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/action"
)

func ConnectToVoice(c *websocket.Conn, discord *discordgo.Session, userId string) {

	c.SetCloseHandler(func(code int, text string) error {
		return c.Close()
	})
	go action.ConnectToChannelWithUserId(c, discord, userId)
}
