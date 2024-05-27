package action

import (
	"fmt"
	player "github.com/MathisBurger/discord-dca-player"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
)

func ConnectToChannelWithUserId(dc *discordgo.Session, userId string) string {
	vs := getChannelWithUserId(dc, userId)
	if vs == nil {
		return "User not in channel"
	}
	_, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
	if err != nil {
		return err.Error()
	}
	return "Successfully joined channel"
}

func PlaySound(dc *discordgo.Session, userId string, sound string, ws *websocket.Conn) {
	vs := getChannelWithUserId(dc, userId)
	if vs == nil {
		_ = ws.WriteJSON(types.WebsocketResponse{
			Message: "You are not in channel",
			Status:  200,
		})
		return
	}

	vc, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, false)
	if err != nil {
		_ = ws.WriteJSON(types.WebsocketResponse{
			Message: "Cannot connect to channel " + err.Error(),
			Status:  200,
		})
		return
	}
	_ = ws.WriteJSON(types.WebsocketResponse{
		Message: "Start playing sound",
		Status:  200,
	})
	err = player.Play(sound, vc)
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
