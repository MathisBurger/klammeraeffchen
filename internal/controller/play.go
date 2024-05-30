package controller

import (
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/types"
	"log"
)

type PlaySoundRequest struct {
	FileName string `json:"fileName"`
	UserID   string `json:"userId"`
}

func PlaySoundController(ctx *fiber.Ctx) error {
	var request PlaySoundRequest
	discord, _ := ctx.Locals("discord").(*discordgo.Session)
	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	vs := getChannelWithUserId(discord, request.UserID)
	if vs == nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}
	vc, err := discord.ChannelVoiceJoin(vs.GuildID, vs.ChannelID, false, true)
	if err != nil {
		log.Println(err)
	}
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = conn.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: types.PlayStatus{
					AudioFile: request.FileName,
					Status:    true,
				},
			})
		}
	}
	dgvoice.PlayAudioFile(vc, "./uploads/"+request.FileName, make(<-chan bool))
	err = vc.Disconnect()
	if err != nil {
		log.Println(err)
	}
	defer vc.Close()
	for _, conn := range types.WebsocketConnections {
		if conn != nil {
			_ = conn.WriteJSON(types.WebsocketResponse{
				Message: "Playing sound",
				Status:  200,
				Action:  types.PlayStatusUpdated,
				Content: types.PlayStatus{
					AudioFile: request.FileName,
					Status:    false,
				},
			})
		}
	}
	return ctx.SendStatus(fiber.StatusOK)
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
