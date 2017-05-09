package main

import (
	"fmt"
	"time"
)

var (
	currencies = []string{"BTC", "ETH", "XRP"}
)

func main() {

	updates := make(chan *Currency)
	for _, c := range currencies {
		go StartPolling(c, 5, updates)
	}

	values := make(map[string]string)
	for {
		currency := <-updates
		values[currency.name] = fmt.Sprint(currency.prices["USD"])

		drawText(values)
	}
}

func drawText(values map[string]string) {
	fmt.Print("\r")
	for _, c := range currencies {
		fmt.Printf("[%s: %s USD] ", c, values[c])
	}
	fmt.Print(time.Now().Format("2006-01-02 15:04:05"), "         ")
}
