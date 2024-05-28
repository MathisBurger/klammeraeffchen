package action

import (
	"context"
	"encoding/json"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"io"
	"klammerAeffchen/internal/configuration"
	"net/http"
)

type authRequest struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

type meResponse struct {
	User UserResponseModel `json:"user"`
}

type UserResponseModel struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	GlobalName string `json:"global_name"`
	Avatar     string `json:"avatar"`
}

func AuthorizeWithCode(code string, config configuration.Config) (*oauth2.Token, error) {
	conf := &oauth2.Config{
		RedirectURL:  "http://localhost:5173/authWithCode",
		ClientID:     config.BotClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{discord.ScopeIdentify},
		Endpoint:     discord.Endpoint,
	}
	token, err := conf.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserModel(auth *oauth2.Token) (*UserResponseModel, error) {
	url := "https://discord.com/api/v10/oauth2/@me"
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("Authorization", "Bearer "+auth.AccessToken)
	request.Header.Set("Accept", "application/json")
	res, err := http.DefaultClient.Do(request)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response meResponse
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response.User, nil
}
