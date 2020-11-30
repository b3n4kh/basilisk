package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
)

var kitcoBaseURL = "www.kitco.com/charts"

// ETFClient holds the ETFClient parameters.
type ETFClient struct {
	Metal    string `json:"metal"`
	Currency string `json:"currency"`
	Unit     string `json:"unit"`
}

// Quote from ETFClient.
type Quote struct {
	Bid float64 `json:"bid"`
	Ask float64 `json:"ask"`
}

func newDefaultETFClient() *ETFClient {
	return &ETFClient{
		Metal:    Silver,
		Currency: "EUR",
		Unit:     "Kilo",
	}
}

func newETFClient(metal string, currency string, unit string) *ETFClient {
	return &ETFClient{
		Metal:    metal,
		Currency: currency,
		Unit:     unit,
	}
}

func (f *ETFClient) getRates() (Quote, error) {
	var response Quote
	url := f.getURL()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	cookieUnit := http.Cookie{Name: "chart_unit_ag", Value: "kg"}
	cookieCurrency := http.Cookie{Name: "chart_currency_ag", Value: "EUR"}

	req.AddCookie(&cookieUnit)
	req.AddCookie(&cookieCurrency)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	doc, err := htmlquery.Parse(resp.Body)
	if err != nil {
		return response, err
	}

	spanBid := htmlquery.FindOne(doc, `//*[@id="sp-bid"]`)
	spanAsk := htmlquery.FindOne(doc, `//*[@id="sp-ask"]`)
	spanAskText := htmlquery.InnerText(spanAsk)
	spanBidText := htmlquery.InnerText(spanBid)

	response.Ask, err = strconv.ParseFloat(spanAskText, 64)
	response.Bid, err = strconv.ParseFloat(spanBidText, 64)

	return response, err
}

func (f *ETFClient) getURL() string {
	var url bytes.Buffer

	url.WriteString("https://")
	url.WriteString(kitcoBaseURL)
	url.WriteString("/live")
	url.WriteString(f.Metal)
	url.WriteString(".html")

	return url.String()
}
