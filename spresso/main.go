package main

import (
	"fmt"
	"spresso-sdk-go/spresso/auth"
)

func main() {
	fmt.Println("Hello, world.")
	authClient := auth.NewAuthClient("foKGFuInp9llIfVIXWoa5M6fJvFZmM4E", "7ugRF2iE7wDpJ5-IZkybHXZ2E5XRuket91HhBc-94F2MuXF6rUsL8Sl09WOdZF5I")
	// client := http_client.NewRestyClient(nil, nil)
	resp, err := authClient.Authenticate()
	// resp, err := client.R(context.TODO(), "SDK", 200).
	// 	SetHeader("Content-Type", "application/json").
	// 	SetBody(`{"client_id": "foKGFuInp9llIfVIXWoa5M6fJvFZmM4E",
	// 	"client_secret": "7ugRF2iE7wDpJ5-IZkybHXZ2E5XRuket91HhBc-94F2MuXF6rUsL8Sl09WOdZF5I",
	// 	"audience": "https://spresso-api",
	// 	"grant_type": "client_credentials"
	// }`).
	// 	Post("https://dev-369tg5rm.us.auth0.com/oauth/token")
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)

}
