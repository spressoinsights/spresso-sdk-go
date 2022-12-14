package main

import (
	"fmt"
	"spresso-sdk-go/spresso/client"
	"spresso-sdk-go/spresso/request"
)

func main() {
	fmt.Println("Hello, world.")
	config := client.Config{
		Environment:  "local",
		OrgId:        "test",
		Service:      "PriceOptimization",
		ClientId:     "BKW7vdWHkSplXj6VshA7iEB8iiH6lNSI",
		ClientSecret: "Pv1om60PDE0RV7M8WmhsGPyeeLaXkwF_1bFXQSFZT12Hgnr0bb8vXp_Lm3tHViWY",
		Url:          "https://public-catalog-api.us-east4.staging.spresso.com"}

	poc := request.PriceOptimizationRequestData{
		ItemId:                 "BOXED-123",
		DeviceId:               "f77c77ad-7a8f-43f1-be14-b92a614788d5",
		UserId:                 "e37c76ad-9a8f-83f1-be13-192a614788d3",
		DefaultPrice:           23.99,
		OverrideToDefaultPrice: true,
	}

	client, err := client.NewClient(&config)

	client.GetAuth().Authenticate()

	// fmt.Println(client.GetAuth().GetToken())

	fmt.Println(err)

	price, err2 := client.PriceOptimization().GetPriceOptimization(poc)

	fmt.Println(price)
	// fmt.Println(poc)
	fmt.Println(err2)

}
