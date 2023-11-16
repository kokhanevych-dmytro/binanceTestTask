package pkg

import (
	"context"
	"github.com/aiviaio/go-binance/v2"
	"log"
	"sync"
)

func GetPrice(symbol string, client *binance.Client, ch chan map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	price, err := client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		log.Printf("Failed to get price for %s: %v\n", symbol, err)
		return
	}

	ch <- map[string]string{symbol: price[0].Price}
}
