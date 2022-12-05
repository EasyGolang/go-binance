package main

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2/futures"
)

var BinanceKey = struct {
	ApiKey    string `bson:"ApiKey"`
	SecretKey string `bson:"SecretKey"`
}{
	ApiKey:    "e1ch7VN2DHqymTJMDRZLDjoyWHdjAaT2anp06elWJwOOZ51GoXXXQpOAeDenWEml",
	SecretKey: "E9w2QmhtJC6Z326mB9fcCFH6syB2qt7TbvwaFfmxWtrWdAFrQNTIVvT8bSKLtV9e",
}

func main() {
	ApiKey := BinanceKey.ApiKey
	SecretKey := BinanceKey.SecretKey

	BinanceClient := futures.NewClient(ApiKey, SecretKey)

	res, err := BinanceClient.NewGetBalanceService().Do(context.Background())

	fmt.Println(res)
	fmt.Println(err)

	for _, val := range res {
		fmt.Println(val)
	}
}
