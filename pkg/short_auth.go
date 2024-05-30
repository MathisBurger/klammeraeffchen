package pkg

const (
	RequestTypeStore = "STORE"
	RequestTypeAuth  = "AUTH"
)

// Short auth message
type ShortAuthMessage struct {
	Type string
	Data interface{}
}

// Token for short auth
type shortAuthToken struct {
	UserID string
	Token  string
}

// Handles the short auth
func ShortShortAuthHandler(channel chan *ShortAuthMessage) {
	var validTokens []*shortAuthToken

	for {
		req := <-channel
		if req.Type == RequestTypeStore {
			tokenStr := RandStringRunes(64)
			validTokens = append(validTokens, &shortAuthToken{
				UserID: "",
				Token:  tokenStr,
			})
			channel <- &ShortAuthMessage{
				Type: RequestTypeStore,
				Data: tokenStr,
			}
		} else if req.Type == RequestTypeAuth {
			found := false
			for _, token := range validTokens {
				if token.Token == req.Data.(string) {
					channel <- &ShortAuthMessage{
						Type: RequestTypeAuth,
						Data: true,
					}
					var newTokens []*shortAuthToken
					for _, token := range validTokens {
						if token.Token != req.Data.(string) {
							newTokens = append(newTokens, token)
						}
					}
					validTokens = newTokens
					found = true
					break
				}
			}
			if !found {
				channel <- &ShortAuthMessage{
					Type: RequestTypeAuth,
					Data: false,
				}
			}
		}
	}
}
