package action

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"
	"io"
	"klammerAeffchen/internal/configuration"
	"net/http"
	"net/url"
	"strings"
)

// Request for authorization
type authRequest struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`
}

// Response with user model
type meResponse struct {
	User UserResponseModel `json:"user"`
}

// The response user model
type UserResponseModel struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	GlobalName string `json:"global_name"`
	Avatar     string `json:"avatar"`
}

// AuthorizeWithCode authorizes the user with a code
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

// AuthorizeWithToken authorizes the user with a token
func AuthorizeWithToken(token string, config configuration.Config) (*oauth2.Token, error) {
	authUrl := "https://discord.com/api/v10/oauth2/token"
	data := url.Values{
		"client_id":     {config.BotClientID},
		"client_secret": {config.ClientSecret},
		"grant_type":    {"refresh_token"},
		"refresh_token": {token},
	}
	request, _ := http.NewRequest(http.MethodPost, authUrl, strings.NewReader(data.Encode()))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.SetBasicAuth(config.BotClientID, config.ClientSecret)
	res, err := http.DefaultClient.Do(request)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var response oauth2.Token
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	if response.AccessToken == "" {
		return nil, errors.New("access_token is empty")
	}
	return &response, nil
}

// GetUserModel gets the user model of the currently authorized user
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
