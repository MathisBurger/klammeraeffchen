package controller

import (
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
	"os"
	"strings"
)

type allSoundsData struct {
	Sounds []string `json:"sounds"`
}

func GetAllSounds(c *websocket.Conn) {
	entries, _ := os.ReadDir("./uploads")
	var data []string
	for _, entry := range entries {
		info, _ := entry.Info()
		if strings.Contains(info.Name(), ".mp3") {
			data = append(data, info.Name())
		}
	}
	_ = c.WriteJSON(types.WebsocketResponse{
		Action: types.ActionGetAllSounds,
		Content: allSoundsData{
			Sounds: data,
		},
	})
}
