package main

import (
	"testing"
)

const input = "Date,Open,High,Low,Close,Adj Close,Volume\n1998-07-22,63.148163,65.053703,63.148163,64.596451,34.570709,56336\n1998-07-23,65.308174,65.816116,61.750568,61.750568,33.047661,95571"

func TestFromBytes(t *testing.T) {

	sp := fromBytes([]byte(input))
	if len(sp) != 2 {
		t.Fatalf("incorrect number of stock prices: %d", len(sp))
	}

	if sp[1].open != 65.308174 {
		t.Fatalf("incorrect opening price: %f", sp[1].open)
	}

	if sp[1].close != 33.047661 {
		t.Fatalf("incorrect closing price: %f", sp[1].close)
	}
}
