package types

type WebsocketMessage struct {
	Action  string      `json:"action"`
	Content interface{} `json:"content"`
}

type WebsocketResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
