package main

import "fmt"

type Foo struct {
	One string
	FooOpts
}

type FooOpts struct {
	Two   int
	Three bool
}

// downside of this pattern is that, you can't set default value
// you can't determine whether the passed arguments are non-zero or zero value.
func NewFoo(one string, fooOpts FooOpts) *Foo {
	ret := &Foo{
		One:     one,
		FooOpts: fooOpts,
	}

	return ret
}

func main() {

	f1 := NewFoo("gopher", FooOpts{2, false})
	f2 := NewFoo("hello", FooOpts{Three: false})
	f3 := NewFoo("teddy", FooOpts{Two: 7})

	fmt.Printf("%+v\n", f1)
	fmt.Printf("%+v\n", f2)
	fmt.Printf("%+v\n", f3)

}
