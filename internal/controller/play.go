package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"github.com/mitchellh/mapstructure"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/types"
	"net/http"
)

type playSoundRequest struct {
	FileName string `json:"fileName"`
}

func PlaySound(c *websocket.Conn, discord *discordgo.Session, content interface{}, userId string) {
	var contentRequest playSoundRequest
	if err := mapstructure.Decode(content, &contentRequest); err != nil {
		_ = c.WriteJSON(types.WebsocketResponse{
			Message: "Invalid type for playSoundRequest",
			Status:  http.StatusNotAcceptable,
		})
		return
	}
	// Sends websocket response automatically
	action.PlaySound(discord, userId, contentRequest.FileName, c)
}
