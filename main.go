package main

import (
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

	futures.UseTestnet = true
	BinanceClient := futures.NewClient(ApiKey, SecretKey)

	fmt.Println(BinanceClient.APIKey)
}
