package controller

import (
	_ "encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/internal/action"
)

type voiceConnectRequest struct {
	UserId string `json:"userId"`
}

type voiceConnectResponse struct {
	Message string `json:"message"`
}

func ConnectToVoice(ctx *fiber.Ctx) error {
	var req voiceConnectRequest
	if err := ctx.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	discord, _ := ctx.Locals("discord").(*discordgo.Session)
	resp := action.ConnectToChannelWithUserId(discord, req.UserId)
	return ctx.JSON(voiceConnectResponse{
		Message: resp,
	})
}
