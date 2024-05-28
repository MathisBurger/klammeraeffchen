package action

import (
	"bytes"
	"encoding/json"
	"io"
	"klammerAeffchen/internal/configuration"
	"net/http"
)

type AuthResponseInitial struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

type meResponse struct {
	User UserResponseModel `json:"user"`
}

type UserResponseModel struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func AuthorizeWithCode(code string, config *configuration.Config) (*AuthResponseInitial, error) {
	url := "https://discord.com/api/v10/oauth2/token"
	body := `{"grant_type": "authorization_code", "code": "` + code + `"}`
	bodySteam := bytes.NewReader([]byte(body))
	request, _ := http.NewRequest(http.MethodPost, url, bodySteam)
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
	var response AuthResponseInitial
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func GetUserModel(auth *AuthResponseInitial) (*UserResponseModel, error) {
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
