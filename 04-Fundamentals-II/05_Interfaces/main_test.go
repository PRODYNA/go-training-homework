package main

import (
	"testing"
)

// TODO create a test, that does not panic
//
//  something is missing here
//
//func TestMeNoPanic(t *testing.T) {
//	u := TestMe{}
//	l := .... and here
//	r := u.TestMe(&l)
//	if r != 1 {
//		t.Errorf("Test failed! Result %v", r)
//	}
//}

// SNIP START
type FooLibMock struct {
}

func (f *FooLibMock) fooLib() int {
	return 1
}

func (f FooLibMock) fooLibWithArgs([]int) int {
	return 1
}

func TestMeNoPanic(t *testing.T) {
	u := TestMe{}
	l := FooLibMock{}
	r := u.TestMe(&l)
	if r != 1 {
		t.Errorf("Test failed! Result %v", r)
	}
}

func TestMeNoPanicWithArgs(t *testing.T) {
	u := TestMe{}
	l := FooLibMock{}
	r := u.TestMeWithArgs(l)
	if r != 1 {
		t.Errorf("Test failed! Result %v", r)
	}
}

// SNIP END

func TestMePanic(t *testing.T) {
	// TODO test for panic
	// SNIP START
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	// SNIP END

	u := TestMe{}
	l := FooLib{}
	_ = u.TestMe(&l)
	t.Errorf("This code should never be reached.")
}
