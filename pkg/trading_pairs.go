package pkg

import (
	"context"
)

func GetTradingPairs(client binanceClient) ([]string, error) {
	exchangeInfo, err := client.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	var symbols []string
	for _, s := range exchangeInfo.Symbols {
		if len(symbols) >= 5 {
			break
		}
		symbols = append(symbols, s.Symbol)
	}
	return symbols, nil
}
