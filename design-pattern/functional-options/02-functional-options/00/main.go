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

type fooOpt func(*Foo)

func TwoOpt(two int) fooOpt {
	return func(f *Foo) {
		f.Two = two
	}
}

func ThreeOpt(three bool) fooOpt {
	return func(f *Foo) {
		f.Three = three
	}
}

func NewFoo(one string, opts ...fooOpt) Foo {
	foo := Foo{
		One: one,
		FooOpts: FooOpts{
			Two:   9,
			Three: true,
		},
	}

	for _, opt := range opts {
		opt(&foo)
	}
	return foo
}

func main() {

	f1 := NewFoo("f1", TwoOpt(4), ThreeOpt(false))
	f2 := NewFoo("f2", ThreeOpt(false))
	f3 := NewFoo("f3")
	f4 := NewFoo("f4", TwoOpt(7))

	fmt.Printf("%+v \n", f1)
	fmt.Printf("%+v \n", f2)
	fmt.Printf("%+v \n", f3)
	fmt.Printf("%+v \n", f4)

}
