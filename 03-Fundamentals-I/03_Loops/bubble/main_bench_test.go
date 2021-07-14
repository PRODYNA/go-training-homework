package main

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

var a10 []int
var a100 []int
var a1k []int
var a10k []int

func setup() {
	rand.Seed(time.Now().Unix())

	// TODO build random arrays of 10, 100, 1k and 10k items
	// SNIP START
	a10 = rand.Perm(10)
	a100 = rand.Perm(100)
	a1k = rand.Perm(1000)
	a10k = rand.Perm(10000)
	// SNIP END
}

func bench(b *testing.B, a []int) {
	for i := 0; i < b.N; i++ {
		BubbleSort(a)
	}
}

func BenchmarkBubbleSort10(b *testing.B) {
	bench(b, a10)
}

func BenchmarkBubbleSort100(b *testing.B) {
	bench(b, a100)
}

func BenchmarkBubbleSort1k(b *testing.B) {
	bench(b, a1k)
}

func BenchmarkBubbleSort10k(b *testing.B) {
	bench(b, a10k)
}
