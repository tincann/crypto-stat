package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Currency struct {
	name   string
	prices map[string]float64
}

func StartPolling(currency string, interval int, cur chan *Currency) {
	url := fmt.Sprintf("https://min-api.cryptocompare.com/data/price?fsym=%s&tsyms=BTC,USD,EUR", currency)
	for {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Could not retrieve %s: %s", currency, err)
			return
		}

		jsonMap := make(map[string]float64)
		err = json.NewDecoder(resp.Body).Decode(&jsonMap)
		if err != nil {
			fmt.Println(err)
		}
		cur <- &Currency{name: currency, prices: jsonMap}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
