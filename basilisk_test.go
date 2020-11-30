package main

import (
	"testing"
)

func TestURL(t *testing.T) {
	expected := "https://" + exchangeAPIURL + "?base=EUR&symbols=USD"
	client := newCurrencyClient("EUR", []string{"USD"})
	actual := client.getURL()

	if expected != actual {
		t.Fatalf("Expected %s but got %s", expected, actual)
	}
}

func TestRates(t *testing.T) {
	client := newCurrencyClient("EUR", []string{"USD"})
	usd2EuroRate, err := client.getRates()
	if err != nil {
		t.Fatalf("Expected no error but got %v", err)
	}
	if usd2EuroRate == 0.0 {
		t.Fatalf("Expected something different then 0.0 but got %v", usd2EuroRate)
	}
}
