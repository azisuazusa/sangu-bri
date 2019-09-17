package main

import (
	"fmt"

	bri "github.com/abhisuta/sangu-bri"
)

func main() {
	briClient := bri.NewClient()
	briClient.BaseUrl = "https://sandbox.partner.api.bri.co.id"
	briClient.ClientId = "p6FGDaCZGoaL8F26dvjCdBfhl8VA0wjf"
	briClient.ClientSecret = "5L4QGueGYTdzin30"
	briClient.LogLevel = 3

	coreGateway := bri.CoreGateway{
		Client: briClient,
	}

	res, _ := coreGateway.GetToken()
	fmt.Println(res)
}
