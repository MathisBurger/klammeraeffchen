package controller

import (
	_ "encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/types"
)

func ConnectToVoice(c *websocket.Conn, discord *discordgo.Session, userId string) {

	c.SetCloseHandler(func(code int, text string) error {
		return c.Close()
	})
	go action.ConnectToChannelWithUserId(discord, userId)
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "Successfully joined voice channel",
		Status:  200,
		Action:  types.ActionConnect,
		Content: nil,
	})
}
