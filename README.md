# Chainlink CryptoCompare Go Bridges Example

This example builds upon LinkPool's [CryptoCompare example](https://github.com/linkpoolio/bridges/tree/master/examples/cryptocompare).

## Build

```
go build -o go-cc
```

## Lambda

Zip the binary

```
zip go-cc.zip go-cc
```

Then upload to AWS Lambda, use `go-cc` as the handler.