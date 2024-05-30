package types

import (
	"github.com/gofiber/contrib/websocket"
	"time"
)

const (
	AuthRefreshToken      = "AUTH_REFRESH_TOKEN"
	AuthUserID            = "AUTH_USER_ID"
	ActionGetCommonGuilds = "GET_COMMON_GUILDS"
	ActionGetShortAuth    = "GET_SHORT_AUTH"
	ActionGetAllSounds    = "GET_ALL_SOUNDS"
	PlayStatusUpdated     = "PLAY_STATUS_UPDATED"
)

var WebsocketConnections []*websocket.Conn

type WebsocketMessage struct {
	Action  string      `json:"action"`
	Content interface{} `json:"content"`
}

type WebsocketResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
	Action  string      `json:"action"`
}

type WebsocketAuthModel struct {
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    time.Time `json:"expires_in"`
}

type PlayStatus struct {
	AudioFile string `json:"audio_file"`
	Status    bool   `json:"status"`
}
