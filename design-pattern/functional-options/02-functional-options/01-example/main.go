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

func Two(two int) func(*Foo) {
	return func(f *Foo) {
		f.Two = two
	}
}

func Three(three bool) func(*Foo) {
	return func(f *Foo) {
		f.Three = three
	}
}

func NewFoo(one string, opts ...func(*Foo)) *Foo {
	ret := &Foo{
		One: one,
		FooOpts: FooOpts{
			Two:   9,
			Three: true,
		},
	}

	for _, opt := range opts {
		opt(ret)
	}

	return ret
}

func main() {

	f1 := NewFoo("gopher", Two(2), Three(false))
	f2 := NewFoo("hello", Three(false))
	f3 := NewFoo("teddy", Two(7))

	fmt.Printf("%+v\n", f1)
	fmt.Printf("%+v\n", f2)
	fmt.Printf("%+v\n", f3)

}
