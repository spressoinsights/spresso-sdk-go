package auth

import (
	"encoding/json"
	"spresso-sdk-go/spresso/http_client"
	"time"

	"gopkg.in/resty.v1"
)

type Auth struct {
	ClientId     string
	HTTPClient   http_client.RestyClient
	ClientSecret string
	Audience     string "https://spresso-api"
	GrantType    string "client_credentials"
	Token        string "Bearer"
	TTL          int
	CreatedAt    time.Time "Time"
}

func NewAuth(clientId string, clientSecret string) {
	return &Auth{
		ClientId:          clientId,
		ClientSecret:      clientSecret,
		defaultRetryCount: retryCount,
	}
}

func NewRestyClient(defaultTimeout *time.Duration, defaultRetryCount *int) RestyClient {
	client := resty.New()

	// use go-json for faster marshal/unmarshal
	client.JSONMarshal = json.Marshal
	client.JSONUnmarshal = json.Unmarshal

	timeout := 10 * time.Second
	if defaultTimeout != nil {
		timeout = *defaultTimeout
	}

	retryCount := 0
	if defaultRetryCount != nil {
		retryCount = *defaultRetryCount
	}

	return &restyClient{
		underlyingClient:  client,
		defaultTimeout:    timeout,
		defaultRetryCount: retryCount,
	}
}

func RenewToken() bool {
	return true
}

func getToken() string {
	return Auth.Token

	return nil
}

func Authenticate() {

}
