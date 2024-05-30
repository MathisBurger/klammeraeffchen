package action

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
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
	vc, err := dc.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = conn.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: types.PlayStatus{
					AudioFile: sound,
					Status:    true,
				},
			})
		}
	}
	dgvoice.PlayAudioFile(vc, "./uploads/"+sound, make(<-chan bool))
	_ = vc.Disconnect()
	vc.Close()
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = conn.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: types.PlayStatus{
					AudioFile: sound,
					Status:    false,
				},
			})
		}
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
