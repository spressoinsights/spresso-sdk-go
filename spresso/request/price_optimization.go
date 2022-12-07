package request

import (
	"spresso-sdk-go/spresso/http_client"
	"strconv"
)

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
	GetPriceOptimization(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error)
	PostPriceOptimization(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error)
}

type priceOptimizationRequest struct {
	http_client.RestyRequest
	url string
}

func PriceOptimization(req http_client.RestyRequest, url string) PriceOptimizationRequest {
	return &priceOptimizationRequest{
		req,
		url,
	}
}

func (r *priceOptimizationRequest) GetPriceOptimization(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error) {
	req, err := r.SetQueryParam("itemId", data.ItemId).
		SetQueryParam("deviceId", data.DeviceId).
		SetQueryParam("userId", data.UserId).
		SetQueryParam("defaultPrice", strconv.FormatFloat(data.DefaultPrice, 'f', 2, 64)).
		SetQueryParam("overrideToDefaultPrice", strconv.FormatBool(data.OverrideToDefaultPrice)).
		Get(r.url)

	if err != nil {
		return nil, err
	}

	return req.Result().(*PriceOptimizationResponse), nil
}

func (r *priceOptimizationRequest) PostPriceOptimization(data PriceOptimizationRequestData) (*PriceOptimizationResponse, error) {
	req, err := r.SetBody(data).
		Post(r.url)

	if err != nil {
		return nil, err
	}

	return req.Result().(*PriceOptimizationResponse), nil
}
