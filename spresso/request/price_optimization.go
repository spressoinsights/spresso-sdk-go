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

type PriceOptimizationData struct {
	DeviceId         string  `json:"deviceId"`
	UserId           string  `json:"userId"`
	ItemId           string  `json:"itemId"`
	Price            float64 `json:"price"`
	IsPriceOptimized bool    `json:"isPriceOptimized"`
	TtlMs            int     `json:"ttlMs"`
}

type PriceOptimizationResponse struct {
	Data PriceOptimizationData `json:"data"`
}

type PriceOptimizationRequest interface {
	GetPriceOptimization(data PriceOptimizationRequestData) (PriceOptimizationData, error)
	PostPriceOptimization(data PriceOptimizationRequestData) (PriceOptimizationData, error)
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

func (r *priceOptimizationRequest) GetPriceOptimization(data PriceOptimizationRequestData) (PriceOptimizationData, error) {
	res, err := r.SetQueryParam("itemId", data.ItemId).
		SetQueryParam("deviceId", data.DeviceId).
		SetQueryParam("userId", data.UserId).
		SetQueryParam("defaultPrice", strconv.FormatFloat(data.DefaultPrice, 'f', 2, 64)).
		SetQueryParam("overrideToDefaultPrice", strconv.FormatBool(data.OverrideToDefaultPrice)).
		SetResult(PriceOptimizationResponse{}).
		Get(r.url + "/v1/priceOptimizations")

	if err != nil {
		return PriceOptimizationData{}, err
	}

	return res.Result().(*PriceOptimizationResponse).Data, nil
}

func (r *priceOptimizationRequest) PostPriceOptimization(data PriceOptimizationRequestData) (PriceOptimizationData, error) {
	res, err := r.SetBody(data).
		Post(r.url + "/v1/priceOptimizations")

	if err != nil {
		return PriceOptimizationData{}, err
	}

	return res.Result().(*PriceOptimizationResponse).Data, nil
}
