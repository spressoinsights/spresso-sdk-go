package auth

import (
	"context"
	"time"

	"spresso-sdk-go/spresso/http_client"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
)

type AuthClient interface {
	RenewToken() bool
	GetToken() string
	Authenticate() (*resty.Response, error)
}
type authClient struct {
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
	Access_token string `json:"access_token"`
	Scope        string `json:"scope"`
	Expires_in   int    `json:"expires_in"`
	Token_type   string `json:"token_type"`
}

func NewAuthClient(clientId string, clientSecret string) AuthClient {
	return &authClient{

		restyRequest: http_client.NewRestyClient(nil, nil).R(context.TODO(), "Auth", 200),
		clientId:     clientId,
		clientSecret: clientSecret,
		CreatedAt:    time.Now(),
		Audience:     `https://spresso-api`,
		GrantType:    `client_credentials`,
		AuthURL:      `https://dev-369tg5rm.us.auth0.com/oauth/token`,
	}
}

func (c *authClient) RenewToken() bool {
	resp, err := c.Authenticate()
	if err == nil && resp.StatusCode() == 200 {
		return true
	}
	return false
}

func (c *authClient) GetToken() string {
	if int(time.Now().Sub(c.CreatedAt).Seconds()) < c.TTL {
		return "Bearer " + c.Token
	}
	return "Expired Token"
}

func (c *authClient) Authenticate() (*resty.Response, error) {
	body, err := json.Marshal(&AuthBody{c.clientId, c.clientSecret, c.Audience, c.GrantType})
	resp, err := c.restyRequest.SetHeader("Content-type", "application/json").
		SetResult(AuthResponse{}).
		SetBody(body).Post(c.AuthURL)
	if err == nil && resp.StatusCode() == 200 {
		c.Token = resp.Result().(*AuthResponse).Access_token
		c.TTL = resp.Result().(*AuthResponse).Expires_in
		c.CreatedAt = time.Now()
	}
	return resp, err
}
