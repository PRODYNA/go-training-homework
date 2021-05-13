package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

type StockPrice struct {
	date  time.Time
	open  float64
	close float64
}

func main() {

	b, err := ioutil.ReadFile("stock.csv")
	if err != nil {
		log.Fatal(fmt.Errorf("could not read file: %v", err))
	}

	fromBytes(b)
}

func fromBytes(b []byte) []StockPrice {
	// TODO convert bytes to lines
	// lines := ...
	// SNIP START
	l := string(b)
	lines := strings.Split(l, "\n")
	// SNIP END

	// TODO remove first line from csv as this contains header
	// SNIP START
	lines = lines[1:]
	// SNIP END

	// TODO declare a slice of type StockPrice with name prices
	// SNIP START
	var prices []StockPrice
	// SNIP END
	for _, line := range lines {
		f := strings.Split(line, ",")
		d := f[0] // date of stock price
		o := f[1] // opening price
		c := f[5] // adjusted close price

		// TODO parse d (date), o (open price) and c (adjusted close price) to t, op and cl
		// SNIP START
		// Mon Jan 2 15:04:05 MST 2006
		t, err := time.Parse("2006-01-02", d)
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse date %s: %v", d, err))
		}
		op, err := strconv.ParseFloat(o, 64)
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse open %s: %v", o, err))
		}
		cl, err := strconv.ParseFloat(c, 64)
		if err != nil {
			log.Fatal(fmt.Errorf("could not parse adjusted close %s: %v", o, err))
		}
		// SNIP END

		// TODO create a StockPrice struct and add it to the prices slice
		// SNIP START
		sp := StockPrice{
			date:  t,
			open:  op,
			close: cl,
		}

		log.Printf("parsed stock price: %v", sp)
		prices = append(prices, sp)
		// SNIP END
	}

	return prices
}
