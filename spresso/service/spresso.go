package spresso

import (
	"context"
	"spresso-sdk-go/spresso/auth"
	"spresso-sdk-go/spresso/http_client"
	"spresso-sdk-go/spresso/request"
)

// Spresso API.
const API = "https://spresso-api"

// Client Credential for API.
const Client_Credential = "client_credentials"

type Client struct {
	config       Config
	restyRequest http_client.RestyRequest
	auth         *auth.AuthClient
}

// Config represents the configuration used to initialize an Client.
type Config struct {
	Environment  string
	OrgId        string
	Service      string
	clientId     string
	clientSecret string
	url          string
}

func NewClient(config *Config) (*Client, error) {
	if config == nil {
		var err error
		if config, err = getConfigDefaults(); err != nil {
			return nil, err
		}
	}
	ctx := context.TODO()
	resty := http_client.NewRestyClient(nil, nil).R(ctx, "SpressoClient", 200)

	return &Client{
		config:       *config,
		auth:         auth.NewAuthClient(config.clientId, config.clientSecret),
		restyRequest: resty,
	}, nil
}

func getConfigDefaults() (*Config, error) {

	return &Config{
		Environment:  "Staging",
		OrgId:        "BOXED",
		Service:      "POC",
		clientId:     "",
		clientSecret: "",
		url:          "",
	}, nil
}

func (c *Client) prepareRequest(config Config) http_client.RestyRequest {
	return c.restyRequest.SetHeader("Authorization", c.auth.Token).
		SetHeader("Accept", "application/json")
}

func (c *Client) PriceOptimization() request.PriceOptimizationRequest {
	return request.PriceOptimization(c.prepareRequest(c.config), c.config.url)
}
