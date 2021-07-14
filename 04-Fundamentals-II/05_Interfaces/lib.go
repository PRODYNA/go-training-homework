package main

// Can't touch this!
// Hammer time!

type FooLib struct {
}

func (f *FooLib) fooLib() int {
	panic("This should never be called.")
}

func (f FooLib) fooLibWithArgs(_ []int) int {
	panic("This should never be called.")
}

func (f FooLib) fooLibWithVarArgs(...int) {
	panic("This should never be called.")
}
