package main

import (
	"net/http"
	"os"

	"github.com/linkpoolio/bridges/bridge"
)

// CryptoCompare is the most basic Bridge implementation, as it only calls the api:
// https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD,JPY,EUR
type CryptoCompare struct{}

var (
	ApiKey = os.Getenv("API_KEY")
	BaseUrl = "https://min-api.cryptocompare.com/data/price"
)

// Run is the bridge.Bridge Run implementation that returns the price response
func (cc *CryptoCompare) Run(h *bridge.Helper) (interface{}, error) {
	r := make(map[string]interface{})
	err := h.HTTPCallWithOpts(
		http.MethodGet,
		BaseUrl,
		&r,
		bridge.CallOpts{
			Auth: bridge.NewAuth(bridge.AuthParam, "apikey", ApiKey),
			Query: map[string]interface{}{
				"fsym": h.GetParam("coin"),
				"tsyms": h.GetParam("market"),
			},
		},
	)
	return r, err
}

// Opts is the bridge.Bridge implementation
func (cc *CryptoCompare) Opts() *bridge.Opts {
	return &bridge.Opts{
		Name:   "CryptoCompare",
		Lambda: true,
	}
}

func main() {
	bridge.NewServer(&CryptoCompare{}).Start(8080)
}