package auth

import (
	"context"
	"time"

	"spresso-sdk-go/spresso/http_client"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

type AuthClient struct {
	restyRequest http_client.RestyRequest
	clientId     string
	clientSecret string
	Audience     string "https://spresso-api"
	GrantType    string "client_credentials"
	AuthURL      string "https://dev-369tg5rm.us.auth0.com/oauth/token"
	Token        string "Bearer"
	TTL          int
	CreatedAt    time.Time "Time"
}

type AuthBody struct {
	Client_id     string `json:"client_id"`
	Client_secret string `json:"client_secret"`
	Audience      string `json:"audience"`
	GrantType     string `json:"grant_type"`
}

type AuthResponse struct {
	access_token string
	scope        string
	expires_in   int
	token_type   string
}

func NewAuthClient(clientId string, clientSecret string) *AuthClient {
	return &AuthClient{

		restyRequest: http_client.NewRestyClient(nil, nil).R(context.TODO(), "Auth", 200),
		clientId:     clientId,
		clientSecret: clientSecret,
		CreatedAt:    time.Now(),
		Audience:     `https://spresso-api`,
		GrantType:    `client_credentials`,
		AuthURL:      `https://dev-369tg5rm.us.auth0.com/oauth/token`,
	}
}

func (c *AuthClient) RenewToken() bool {
	resp, err := c.Authenticate()
	if err == nil && resp.StatusCode() == 200 {
		return true
	}
	return false
}

func (c *AuthClient) getToken() string {
	if int(time.Now().Sub(c.CreatedAt).Seconds()) < c.TTL {
		return c.Token
	}
	return ""
}

func (c *AuthClient) Authenticate() (*resty.Response, error) {
	body, err := json.Marshal(&AuthBody{c.clientId, c.clientSecret, c.Audience, c.GrantType})

	resp, err := c.restyRequest.SetHeader("Content-type", "application/json").
		SetResult(AuthResponse{}).
		SetBody(body).Post(c.AuthURL)
	if err != nil && resp.StatusCode() == 200 {
		c.Token = resp.Result().(*AuthResponse).access_token
	}
	return resp, err
}
