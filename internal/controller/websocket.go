package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
)

const (
	ACTION_PLAY    = "PLAY"
	ACTION_CONNECT = "CONNECT"
)

func ApplicationWebsocket(c *websocket.Conn) {
	discord, _ := c.Locals("discord").(*discordgo.Session)
	var msg types.WebsocketMessage
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			_ = c.Close()
			return
		}
		switch msg.Action {
		case ACTION_PLAY:
			PlaySound(c, discord, msg.Content)
			break
		case ACTION_CONNECT:
			ConnectToVoice(c, discord, msg.Content)
		default:
			_ = c.WriteJSON(types.WebsocketResponse{
				Message: "Cannot handle message",
				Status:  400,
			})
			break
		}
	}
}
