package pkg

const (
	RequestTypeStore = "STORE"
	RequestTypeAuth  = "AUTH"
)

type ShortAuthMessage struct {
	Type string
	Data interface{}
}

type shortAuthToken struct {
	UserID string
	Token  string
}

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
