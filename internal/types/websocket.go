package types

type WebsocketMessage struct {
	Action  string      `json:"action"`
	Content interface{} `json:"content"`
}

type WebsocketResponse struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Content interface{} `json:"content"`
}

type WebsocketAuthModel struct {
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}
