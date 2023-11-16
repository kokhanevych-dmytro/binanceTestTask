package pkg

import "github.com/aiviaio/go-binance/v2"

type binanceClient interface {
	NewExchangeInfoService() *binance.ExchangeInfoService
	NewListPricesService() *binance.ListPricesService
}
