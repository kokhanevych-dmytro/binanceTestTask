package main

import (
	"binanceTestTask/pkg"
	"fmt"
	"github.com/aiviaio/go-binance/v2"
	"log"
	"sync"
)

func main() {
	client := binance.NewClient("", "")
	symbols, err := pkg.GetTradingPairs(client)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan map[string]string)
	var wg sync.WaitGroup

	for _, symbol := range symbols {
		wg.Add(1)
		go pkg.GetPrice(symbol, client, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		for symbol, price := range msg {
			fmt.Printf("%s %s\n", symbol, price)
		}
	}
}
