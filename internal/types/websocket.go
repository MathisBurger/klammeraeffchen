package types

import (
	"github.com/gofiber/contrib/websocket"
	"time"
)

const (
	AuthRefreshToken   = "AUTH_REFRESH_TOKEN"
	AuthUserID         = "AUTH_USER_ID"
	ActionGetShortAuth = "GET_SHORT_AUTH"
	ActionGetAllSounds = "GET_ALL_SOUNDS"
)

// All websocket connections
var WebsocketConnections []*websocket.Conn

// A message sent over the websocket
type WebsocketMessage struct {
	Action  string      `json:"action"`
	Content interface{} `json:"content"`
}

// The response sent over the websocket
type WebsocketResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
	Action  string      `json:"action"`
}

// Auth model for the websocket
type WebsocketAuthModel struct {
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
}

// Audio play status
type PlayStatus struct {
	AudioFile string `json:"audio_file"`
	Status    bool   `json:"status"`
}
