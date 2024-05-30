package controller

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"klammerAeffchen/pkg"
	"strings"
)

type uploadAudioResponse struct {
	Message string `json:"message"`
}

func UploadAudio(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("audiofile")
	authCode := ctx.Query("authCode", "")
	authChannel, _ := ctx.Locals("auth").(chan *pkg.ShortAuthMessage)
	authChannel <- &pkg.ShortAuthMessage{
		Type: pkg.RequestTypeAuth,
		Data: authCode,
	}
	response := <-authChannel
	if response.Data.(bool) == false {
		return fiber.NewError(fiber.StatusUnauthorized)
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if !strings.HasSuffix(file.Filename, ".mp3") {
		return fiber.NewError(fiber.StatusBadRequest, "audio file is not mp3")
	}
	destination := fmt.Sprintf("./uploads/%s", file.Filename)
	if err := ctx.SaveFile(file, destination); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	return ctx.JSON(&uploadAudioResponse{
		Message: fmt.Sprintf("%s uploaded successfully", file.Filename),
	})
}
