package auth

import (
	"log"
	"spresso-sdk-go/spresso/http_client"
	"time"
)

type Auth struct {
	ClientId     string
	ClientSecret string
	Audience     string "https://spresso-api"
	GrantType    string "client_credentials"
	Token        string "Bearer"
	TTL          int
	CreatedAt    time.Time "Time"
}

var (
	auth = newAuth()
)

func NewAuth(clientId string, clientSecret string, logger log.Logger, restyClient http_client.RestyClient) {

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
