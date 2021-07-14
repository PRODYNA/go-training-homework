package main

// TODO something is missing here

// SNIP START
type IFooLib interface {
	fooLib() int
}

type IFooLibWithArgs interface {
	fooLibWithArgs([]int) int
}

// SNIP END

type TestMe struct {
}

// TODO create/modify the function under test
//func (t TestMe) TestMe(l FooLib) int {
//	return l.fooLib()
//}
// SNIP START

func (t TestMe) TestMe(l IFooLib) int {
	return l.fooLib()
}

func (t TestMe) TestMeWithArgs(l IFooLibWithArgs) int {
	return l.fooLibWithArgs([]int{})
}

// SNIP END

// Bonus: Create TestMeWithArgs and uncomment the call in main.
// func (t TestMe) TestMeWithArgs(l FooLib) int {
//	return l.fooLibWithArgs([]int{})
//}

func main() {
	l := FooLib{}
	t := TestMe{}

	// this will panic
	t.TestMe(&l)
	// for the bonus assignment
	// t.TestMeWithArgs(l)
}
