package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	BubbleSort(a)

	e := []int{2, 14, 33, 37, 45, 45, 45, 91, 99, 102, 102, 104, 106, 106, 109, 212, 501, 3001, 7800, 9932}
	if !reflect.DeepEqual(a, e) {
		t.Errorf("BubbleSort() = %v, want %v", a, e)
	}
}
