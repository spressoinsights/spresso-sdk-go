package auth

type Auth struct {
	ClientId     string
	ClientSecret string
	Audience     string "https://spresso-api"
	GrantType    string "client_credentials"
	Token        string "Bearer"
}

type PriceOptimizationRequest struct {
	Data *PriceOptimizationRequestData `json:"data"`
}
