package request

type PriceOptimizationRequestData struct {
	ItemId                 string
	DeviceId               string
	UserId                 string
	DefaultPrice           float64
	OverrideToDefaultPrice bool
}

type PriceOptimizationResponse struct {
	ItemId                 *string
	DeviceId               *string
	UserId                 *string
	DefaultPrice           *float64
	OverrideToDefaultPrice *bool
}

type PriceOptimizationRequest interface {
	PriceOptimization() PriceOptimizationRequest
	Get(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error)
	Post(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error)
}
