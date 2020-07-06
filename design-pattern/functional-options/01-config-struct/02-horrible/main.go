package main

import "fmt"

type Foo struct {
	One string
	FooOpts
}

//what if we changed to pointer type?
type FooOpts struct {
	Two   *int
	Three *bool
}

// super ugly code!!!!!
func NewFoo(one string, fooOpts FooOpts) *Foo {
	defaultTwo := 9
	defaultThree := true
	ret := &Foo{
		One: one,
		FooOpts: FooOpts{
			Two:   &defaultTwo,
			Three: &defaultThree,
		},
	}

	if fooOpts.Two != nil {
		ret.Two = fooOpts.Two
	}

	if fooOpts.Three != nil {
		ret.Three = fooOpts.Three
	}

	return ret
}

func (f Foo) String() string {
	return fmt.Sprintf("One: %v Two:%v Three:%v", f.One, *(f.Two), *(f.Three))
}

func main() {

	//and super hard to use
	two, three := 2, false
	f1 := NewFoo("gopher", FooOpts{&two, &three})
	f2 := NewFoo("hello", FooOpts{Three: &thee})
	two = 7
	f3 := NewFoo("teddy", FooOpts{Two: &two})

	fmt.Printf("%+v\n", f1)
	fmt.Printf("%+v\n", f2)
	fmt.Printf("%+v\n", f3)

}
