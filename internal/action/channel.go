package action

import (
	"fmt"
	player "github.com/MathisBurger/discord-dca-player"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
)

func PlaySound(dc *discordgo.Session, userId string, sound string, ws *websocket.Conn) {
	vs := getChannelWithUserId(dc, userId)
	if vs == nil {
		_ = ws.WriteJSON(types.WebsocketResponse{
			Message: "You are not in channel",
			Status:  200,
			Content: nil,
			Action:  types.ActionPlay,
		})
		return
	}

	vc, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
	if err != nil {
		_ = ws.WriteJSON(types.WebsocketResponse{
			Message: "Cannot connect to channel " + err.Error(),
			Status:  200,
			Content: nil,
		})
		return
	}
	_ = ws.WriteJSON(types.WebsocketResponse{
		Message: "Start playing sound",
		Status:  200,
		Content: nil,
	})
	err = player.Play(sound, vc, false)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getChannelWithUserId(dc *discordgo.Session, userId string) *discordgo.VoiceState {
	for _, guild := range dc.State.Guilds {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == userId {
				return vs
			}
		}
	}
	return nil
}

/*func Disconnect(c *websocket.Conn, dc *discordgo.Session, userId string) {
	vs := getChannelWithUserId(dc, userId)
	if vs != nil && dc.VoiceConnections[vs.GuildID] != nil && dc.VoiceConnections[vs.GuildID].ChannelID == vs.ChannelID {
		vc, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
		if err != nil {
			fmt.Println(err.Error())
		}
		err = vc.Disconnect()
		if err != nil {
			fmt.Println(err.Error())
		}
		_ = c.WriteJSON(types.WebsocketResponse{
			Message: "Successfully disconnected",
			Status:  200,
			Action:  types.ActionDisconnect,
			Content: nil,
		})
		return
	}
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "You are not in a channel",
		Status:  200,
		Action:  types.ActionDisconnect,
		Content: nil,
	})
}

func ConnectToChannelWithUserId(c *websocket.Conn, dc *discordgo.Session, userId string) {
	vs := getChannelWithUserId(dc, userId)
	if vs == nil {
		_ = c.WriteJSON(types.WebsocketResponse{
			Message: "You are not in a channel",
			Status:  200,
			Action:  types.ActionConnect,
			Content: nil,
		})
		return
	}
	vc, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "Successfully joined voice channel",
		Status:  200,
		Action:  types.ActionConnect,
		Content: nil,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		_ = <-vc.OpusRecv
	}
}*/
