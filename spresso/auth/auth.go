package auth

import (
	"encoding/json"
	"net/http"
	"spresso-sdk-go/spresso/http_client"
	"time"

	"gopkg.in/resty.v1"
)

type Context struct {
	ClientId     string
	ClientSecret string
	Audience     string "https://spresso-api"
	GrantType    string "client_credentials"
	AuthURL string "https://dev-369tg5rm.us.auth0.com/oauth/token"
	Token        string "Bearer"
	TTL          int
	CreatedAt    time.Time "Time"
}

type AuthBody struct {
	client_id     string
	client_secret string
	audience     string "https://spresso-api"
	grantType    string "client_credentials"
}
type Client struct {
	restyClient http.Client
	context     Context
}

func NewContext(clientId string, clientSecret string) *Context {
	return &Context{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		CreatedAt: time.Now,
	}
}

func NewClient(clientId string, clientSecret string) Client {

	timeout := 10 * time.Second
	if defaultTimeout != nil {
		timeout = *defaultTimeout
	}

	retryCount := 0
	if defaultRetryCount != nil {
		retryCount = *defaultRetryCount
	}

	return &Client{
		restyClient = http_client.NewRestyClient(1, 2),
		context = Context{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			CreatedAt: time.Now,
		}
	}
}

func (c *Client) RenewToken() bool {
	return true
}

func (client *Client) getToken() string {
	if(time.now().Sub(client.createdAt) == context.TTL)
		return Auth.Token
	return nil
}

func (c *Client) Authenticate() {
	resp, err = c.restyClient.setbody(AuthBody{clinet_id = c.Context.ClientId, client_secret= c.Context.ClientSecret})ForceContentType("application/json").Post(c.Context.AuthURL)
	c.Token = resp.Body.get('access_token')
}
