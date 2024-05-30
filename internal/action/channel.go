package action

import (
	"fmt"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
)

type playStatus struct {
	AudioFile string `json:"audio_file"`
	Status    bool   `json:"status"`
}

func PlaySound(dc *discordgo.Session, userId string, sound string, ws *websocket.Conn, audioEndChan chan bool) {
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
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = conn.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: playStatus{
					AudioFile: sound,
					Status:    true,
				},
			})
		}
	}
	dgvoice.PlayAudioFile(vc, "./uploads/"+sound, audioEndChan)
	err = vc.Disconnect()
	if err != nil {
		fmt.Println(err.Error())
	}
	vc.Close()
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = ws.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: playStatus{
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
