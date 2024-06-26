package controller

import (
	"github.com/gofiber/contrib/websocket"
	"klammerAeffchen/internal/types"
	"klammerAeffchen/pkg"
)

// Gets the request for a short living auth and sends the response
// through the websocket
func RequestShortAuth(c *websocket.Conn) {
	authChannel := c.Locals("auth").(chan *pkg.ShortAuthMessage)
	authChannel <- &pkg.ShortAuthMessage{
		Type: pkg.RequestTypeStore,
		Data: nil,
	}
	response := <-authChannel
	_ = c.WriteJSON(types.WebsocketResponse{
		Message: "Successfully authorized",
		Status:  200,
		Action:  types.ActionGetShortAuth,
		Content: response.Data,
	})
}
