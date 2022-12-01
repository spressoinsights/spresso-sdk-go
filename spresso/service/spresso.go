package spresso

import (
	"spresso-sdk-go/spresso/auth"
	"spresso-sdk-go/spresso/http_client"
)

// Spresso API.
const API = "https://spresso-api"

// Client Credential for API.
const Client_Credential = "client_credentials"

type Client struct {
	config      Config
	restyClient http_client.RestyClient
	auth        *auth.AuthClient
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

	return &Client{
		config:      *config,
		auth:        auth.NewAuthClient(config.clientId, config.clientSecret),
		restyClient: http_client.NewRestyClient(nil, nil),
	}, nil
}

// getConfigDefaults reads the default config file, defined by the config file
// env variable, used only when options are nil.
func getConfigDefaults() (*Config, error) {
	// fbc := &Config{}
	// confFileName := os.Getenv("")
	// if confFileName == "" {
	// 	return fbc, nil
	// }
	// var dat []byte
	// if confFileName[0] == byte('{') {
	// 	dat = []byte(confFileName)
	// } else {
	// 	var err error
	// 	if dat, err = ioutil.ReadFile(confFileName); err != nil {
	// 		return nil, err
	// 	}
	// }
	// if err := json.Unmarshal(dat, fbc); err != nil {
	// 	return nil, err
	// }

	// // Some special handling necessary for db auth overrides
	// var m map[string]interface{}
	// if err := json.Unmarshal(dat, &m); err != nil {
	// 	return nil, err
	// }
	// if ao, ok := m["databaseAuthVariableOverride"]; ok && ao == nil {
	// 	// Auth overrides are explicitly set to null
	// 	var nullMap map[string]interface{}
	// 	fbc.AuthOverride = &nullMap
	// }
	// return fbc, nil

	return &Config{
		Environment:  "Staging",
		OrgId:        "BOXED",
		Service:      "POC",
		clientId:     "",
		clientSecret: "",
		url:          "",
	}, nil
}
