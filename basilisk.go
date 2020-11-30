package main

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func main() {
	q, err := quote.Get("SLV")
	if err != nil {
		fmt.Println(err)
	}
	financeQuote := Quote{Bid: q.Bid, Ask: q.Ask}

	currencyClient := newCurrencyClient("EUR", []string{"USD"})
	usd2EuroRate, err := currencyClient.getRates()

	printStonks(financeQuote, usd2EuroRate)

	etfClient := newDefaultETFClient()
	etfQuotes, err := etfClient.getRates()

	printStonks(etfQuotes, usd2EuroRate)

}

func printStonks(q Quote, rate float64) {

	euroAskOz := q.Ask / rate
	euroBidOz := q.Bid / rate
	euroAskKg := q.Ask / 0.0283495
	euroBidKg := q.Bid / 0.0283495

	fmt.Printf("ASK:\nEUR/oz %.2f --> EUR/kg %.2f\n\n", euroAskOz, euroAskKg)
	fmt.Printf("BID:\nEUR/oz %.2f --> EUR/kg %.2f\n\n", euroBidOz, euroBidKg)
}

func getMetalRate(metal string) string {
	return ""
}
