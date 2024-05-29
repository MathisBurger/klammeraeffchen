package controller

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/types"
)

func CommonGuilds(ws *websocket.Conn, discord *discordgo.Session, user *action.UserResponseModel) {
	_ = ws.WriteJSON(types.WebsocketResponse{
		Message: "Successfully fetched common guilds",
		Status:  200,
		Action:  types.ActionGetCommonGuilds,
		Content: action.GetAllCommonGuilds(discord, user.Id),
	})
}
