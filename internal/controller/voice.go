package controller

import (
	_ "encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"github.com/mitchellh/mapstructure"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/types"
	"net/http"
)

type voiceConnectRequest struct {
	UserId string `json:"userId"`
}

func ConnectToVoice(c *websocket.Conn, discord *discordgo.Session, content interface{}) {
	var req voiceConnectRequest
	if err := mapstructure.Decode(content, &req); err != nil {
		_ = c.WriteJSON(types.WebsocketResponse{
			Message: "Invalid type for voiceConnectRequest",
			Status:  http.StatusNotAcceptable,
		})
		return
	}
	action.ConnectToChannelWithUserId(discord, req.UserId)

}
