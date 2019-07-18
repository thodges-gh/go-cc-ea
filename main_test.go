package main

import (
	"testing"

	"github.com/linkpoolio/bridges/bridge"
	"github.com/stretchr/testify/assert"
)

func TestCryptoCompare_Run(t *testing.T) {
	cc := CryptoCompare{}
	data := map[string]interface{}{
		"coin": "ETH",
		"market": "USD",
	}
	query, _ := bridge.ParseInterface(data)
	obj, err := cc.Run(bridge.NewHelper(query))
	assert.Nil(t, err)

	resp, ok := obj.(map[string]interface{})
	assert.True(t, ok)

	_, ok = resp["USD"]
	assert.True(t, ok)
}

func TestCryptoCompare_Opts(t *testing.T) {
	cc := CryptoCompare{}
	opts := cc.Opts()
	assert.Equal(t, opts.Name, "CryptoCompare")
	assert.True(t, opts.Lambda)
}