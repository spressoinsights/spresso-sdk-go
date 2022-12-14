package client

import (
	"context"

	"github.com/spressoinsights/spresso-sdk-go/spresso/auth"
	"github.com/spressoinsights/spresso-sdk-go/spresso/http_client"
	"github.com/spressoinsights/spresso-sdk-go/spresso/request"
)

// Spresso API.
const API = "https://spresso-api"

// Client Credential for API.
const Client_Credential = "client_credentials"

type Client interface {
	PriceOptimization() request.PriceOptimizationRequest
	PrepareRequest(config Config) http_client.RestyRequest
	GetAuth() auth.AuthClient
}

type client struct {
	config       Config
	restyRequest http_client.RestyRequest
	auth         auth.AuthClient
}

// Config represents the configuration used to initialize an Client.
type Config struct {
	Environment  string
	OrgId        string
	Service      string
	ClientId     string
	ClientSecret string
	Url          string
}

func NewClient(config *Config) (Client, error) {
	if config == nil {
		var err error
		if config, err = getConfigDefaults(); err != nil {
			return nil, err
		}
	}
	ctx := context.TODO()
	resty := http_client.NewRestyClient(nil, nil).R(ctx, "SpressoClient", 200)

	return &client{
		config:       *config,
		auth:         auth.NewAuthClient(config.ClientId, config.ClientSecret),
		restyRequest: resty,
	}, nil
}

func getConfigDefaults() (*Config, error) {

	return &Config{
		Environment:  "Staging",
		OrgId:        "BOXED",
		Service:      "POC",
		ClientId:     "",
		ClientSecret: "",
		Url:          "",
	}, nil
}

func (c *client) PrepareRequest(config Config) http_client.RestyRequest {
	return c.restyRequest.SetHeader("Authorization", c.auth.GetToken()).
		SetHeader("Accept", "application/json").
		SetHeader("Content-type", "application/json")
}

func (c *client) GetAuth() auth.AuthClient {
	return c.auth
}

func (c *client) PriceOptimization() request.PriceOptimizationRequest {
	return request.PriceOptimization(c.PrepareRequest(c.config), c.config.Url)
}
