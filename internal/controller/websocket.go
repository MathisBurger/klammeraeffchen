package controller

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/internal/types"
)

const (
	ActionPlay    = "PLAY"
	ActionConnect = "CONNECT"
)

func ApplicationWebsocket(c *websocket.Conn) {
	code := c.Query("code", "")
	config, _ := c.Locals("configuration").(configuration.Config)
	if code == "" {
		_ = c.Close()
		return
	}
	auth, err := action.AuthorizeWithCode(code, config)
	if err != nil {
		fmt.Println(err.Error())
		_ = c.Close()
		return
	}
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "Successfully authorized",
		Status:  200,
		Content: types.WebsocketAuthModel{
			RefreshToken: auth.RefreshToken,
			ExpiresIn:    auth.Expiry,
		},
	})
	me, err := action.GetUserModel(auth)
	if err != nil {
		fmt.Println(err.Error())
		_ = c.Close()
		return
	}
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "User fetched successfully",
		Status:  200,
		Content: me,
	})
	discord, _ := c.Locals("discord").(*discordgo.Session)
	var msg types.WebsocketMessage
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			_ = c.Close()
			return
		}
		switch msg.Action {
		case ActionPlay:
			PlaySound(c, discord, msg.Content, me.Id)
			break
		case ActionConnect:
			ConnectToVoice(c, discord, me.Id)
		default:
			_ = c.WriteJSON(types.WebsocketResponse{
				Message: "Cannot handle message",
				Status:  400,
				Content: nil,
			})
			break
		}
	}
}
