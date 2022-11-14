package spresso

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
)

// Spresso API.
const API = "https://spresso-api"

// Client Credential for API.
const Client_Credential = "client_credentials"

type Client struct {
	config Config
	id     string
	secret string
}

// Config represents the configuration used to initialize an Client.
type Config struct {
	Environment string
	OrgId       string
	Service     string
}

// Auth returns an instance of auth.Client.
func (c *Client) Auth(ctx context.Context) (*auth.Client, error) {
	conf := &config.AuthConfig{
		ClientId:     c.id,
		ClientSecret: c.secret,
		Audience:     API,
		GrantType:    Client_Credential,
	}
	return auth.NewClient(ctx, conf)
}

func NewClient(ctx context.Context, config *Config) (*Client, error) {
	if config == nil {
		var err error
		if config, err = getConfigDefaults(); err != nil {
			return nil, err
		}
	}

	pid := getProjectID(ctx, config, o...)
	ao := defaultAuthOverrides
	if config.AuthOverride != nil {
		ao = *config.AuthOverride
	}

	return &Client{
		authOverride:     ao,
		dbURL:            config.DatabaseURL,
		projectID:        pid,
		serviceAccountID: config.ServiceAccountID,
		storageBucket:    config.StorageBucket,
		opts:             o,
	}, nil
}

// getConfigDefaults reads the default config file, defined by the FIREBASE_CONFIG
// env variable, used only when options are nil.
func getConfigDefaults() (*Config, error) {
	fbc := &Config{}
	confFileName := os.Getenv(spressoEnvName)
	if confFileName == "" {
		return fbc, nil
	}
	var dat []byte
	if confFileName[0] == byte('{') {
		dat = []byte(confFileName)
	} else {
		var err error
		if dat, err = ioutil.ReadFile(confFileName); err != nil {
			return nil, err
		}
	}
	if err := json.Unmarshal(dat, fbc); err != nil {
		return nil, err
	}

	// Some special handling necessary for db auth overrides
	var m map[string]interface{}
	if err := json.Unmarshal(dat, &m); err != nil {
		return nil, err
	}
	if ao, ok := m["databaseAuthVariableOverride"]; ok && ao == nil {
		// Auth overrides are explicitly set to null
		var nullMap map[string]interface{}
		fbc.AuthOverride = &nullMap
	}
	return fbc, nil
}
