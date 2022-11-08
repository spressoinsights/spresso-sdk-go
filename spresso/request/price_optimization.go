package request

type PriceOptimizationRequestData struct {
	ItemId                 string
	DeviceId               string
	UserId                 string
	DefaultPrice           float64
	OverrideToDefaultPrice bool
}

type PriceOptimizationRequest struct {
	Data *PriceOptimizationRequestData `json:"data"`
}
