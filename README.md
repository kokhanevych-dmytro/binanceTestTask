# README.md for Assessment Task Golang

## Overview

This project is an assessment task for Golang, focusing on API interaction, concurrency, and basic data handling. It
connects to the Binance API to retrieve and display the latest prices for a set of trading pairs.

## Getting Started

### Prerequisites

- Ensure you have the latest version of [Golang](https://golang.org/dl/) installed on your system.

### Usage

To run the application, execute the following command in the terminal:

```bash
go run main.go
```

This will start the application, which performs the following tasks:

1. **Connects to the Binance API**: The program uses the `go-binance` library for API interaction. No API keys are
   required as it only accesses public endpoints.

2. **Retrieves Trading Pairs**: The application fetches the first five trading pairs (symbols) from the
   Binance `api/v3/exchangeInfo` endpoint.

3. **Fetches Latest Prices**: For each trading pair, a separate goroutine is created to fetch the latest price from
   the `/fapi/v1/ticker/price` endpoint. The data is communicated back via a channel.

4. **Displays Results**: The main goroutine receives data from worker goroutines and displays the latest prices in the
   format `SYMBOL PRICE`, e.g., `BTCUSDT 22000`.

### Output

The final output on the console will display the latest prices for the first five trading pairs, similar to:

```
BTCUSDT 22000
ETHUSDT 1500
...
```

### Notes

- Ensure your internet connection is active as the application requires access to the Binance API.
- The goroutines are managed and terminated appropriately to prevent any resource leaks.

## Contribution

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what
you would like to change.
