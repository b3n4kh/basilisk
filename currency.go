package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

// CurrencyClient holds the CurrencyClient parameters.
type CurrencyClient struct {
	base    string
	symbols []string
}

// Response is the JSON response object from exchangeratesapi
type Response struct {
	Base  string `json:"base"`
	Date  string `json:"date"`
	Rates rates  `json:"rates"`
}
type rates map[string]float32

var exchangeAPIURL = "api.exchangeratesapi.io/latest"

func newCurrencyClient(base string, symbols []string) *CurrencyClient {
	return &CurrencyClient{
		base:    base,
		symbols: symbols,
	}
}

// Sets base currency.
func (f *CurrencyClient) setBase(currency string) {
	f.base = currency
}

// Sets base currency.
func (f *CurrencyClient) setSymbols(currencies []string) {
	f.symbols = currencies
}

func (f *CurrencyClient) getRates() (float64, error) {
	var response Response
	url := f.getURL()

	body, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer body.Body.Close()

	err = json.NewDecoder(body.Body).Decode(&response)
	if err != nil {
		return 0, err
	}
	rate := response.Rates["USD"]

	return float64(rate), nil
}

func (f *CurrencyClient) getURL() string {
	var url bytes.Buffer

	url.WriteString("https://")
	url.WriteString(exchangeAPIURL)
	url.WriteString("?base=")
	url.WriteString(f.base)
	url.WriteString("&symbols=")
	url.WriteString(strings.Join(f.symbols[:], ","))

	return url.String()
}
