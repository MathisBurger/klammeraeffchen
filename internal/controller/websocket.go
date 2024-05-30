package controller

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"golang.org/x/oauth2"
	"klammerAeffchen/internal/action"
	"klammerAeffchen/internal/configuration"
	"klammerAeffchen/internal/types"
)

func ApplicationWebsocket(c *websocket.Conn) {
	code := c.Query("code", "")
	refreshToken := c.Query("refreshToken", "")
	config, _ := c.Locals("configuration").(configuration.Config)
	if code == "" && refreshToken == "" {
		_ = c.Close()
		return
	}
	var auth *oauth2.Token
	var err error
	if code != "" {
		auth, err = action.AuthorizeWithCode(code, config)
		if err != nil {
			fmt.Println(err.Error())
			_ = c.Close()
			return
		}
	} else {
		auth, err = action.AuthorizeWithToken(refreshToken, config)
		if err != nil {
			fmt.Println(err.Error())
			_ = c.Close()
			return
		}
	}
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "Successfully authorized",
		Status:  200,
		Content: types.WebsocketAuthModel{
			RefreshToken: auth.RefreshToken,
			ExpiresIn:    auth.Expiry,
		},
		Action: types.AuthRefreshToken,
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
		Action:  types.AuthUserID,
	})
	discord, _ := c.Locals("discord").(*discordgo.Session)
	types.WebsocketConnections = append(types.WebsocketConnections, c)
	c.SetCloseHandler(closeHandler(c))
	var msg types.WebsocketMessage
	for {
		err := c.ReadJSON(&msg)
		if err != nil {
			_ = c.Close()
			return
		}
		switch msg.Action {
		case types.ActionPlay:
			PlaySound(c, discord, msg.Content, me.Id)
			break
		case types.ActionGetCommonGuilds:
			CommonGuilds(c, discord, me)
			break
		case types.ActionGetShortAuth:
			RequestShortAuth(c)
			break
		case types.ActionGetAllSounds:
			GetAllSounds(c)
			break
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

func closeHandler(c *websocket.Conn) func(code int, text string) error {
	return func(code int, text string) error {
		var newConns []*websocket.Conn
		for _, conn := range types.WebsocketConnections {
			if conn != c {
				newConns = append(newConns, conn)
			}
		}
		types.WebsocketConnections = newConns
		err := c.Close()
		if err != nil {
			return err
		}
		return nil
	}
}
