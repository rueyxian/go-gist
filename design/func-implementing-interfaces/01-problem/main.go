package main

import (
	"bytes"
	"fmt"
	"sandbox/go-jottings/design-pattern/func-implementing-interfaces/02-solution/brewer"
	"sandbox/go-jottings/design-pattern/func-implementing-interfaces/drink"
	"time"
)

// ========================================
// interface definition
type Brewer interface {
	Brew(d time.Duration, a ...drink.AddIn)
}

// ========================================
// string type
// The go way of a type implementing a interface is
// by creating a method that matches that interface's signature
type grain string

func (drink grain) Brew(d time.Duration, add ...drink.AddIn) {
	brewOperation(string(drink), d, add)
	// fmt.Printf("Name: %+v    Duration: %+v    Add-In: %+v\n", drink, d, add)
}

// ========================================
// struct type
// If you want to have more complex data structure,
// struct is the way to go,
// and the implementation of interface still very straight forward
type tea struct {
	name   string
	brand  string
	weight float64
}

func (t tea) Brew(d time.Duration, add ...drink.AddIn) {
	brewOperation(string(t.name), d, add)
	// fmt.Printf("Name: %+v    Duration: %+v    Add-In: %+v\n", t.name, d, add)
}

// ========================================
// func type
// what if we want make a func type to implement an inteface?

// it does have to be func() string, it can be func that return even a struct type
type coffeeFunc func() string

func (f coffeeFunc) Brew(d time.Duration, add ...drink.AddIn) {
	// the func type return required value that needed to work with the subsequent operation
	name := f()
	// do some operation
	brewOperation(name, d, add)
}

// a factory function that return closure func
func newCoffeeFunc(name string) coffeeFunc {
	return func() string {
		return name
	}
}

// this is completely superfluous as "type coffee string" will do the job in this case
// this is not why a func type is implements an interface

// ----------------------------------------
// first, it should be define on package level
type BrewerFunc func(d time.Duration, add ...drink.AddIn)

// since it is on package level, the definition has to be generic to be useful
// thus, the logic shouldn't be defined here
// the func act as an adapter to allow an odinary func to be used as an interface
// as long as if the func's signature is matched
func (f BrewerFunc) Brew(d time.Duration, add ...drink.AddIn) {
	f(d, add...)
}

// ========================================
func brewOperation(drink string, d time.Duration, add []drink.AddIn) {
	fmt.Printf("%v: start brewing...\n", drink)
	time.Sleep(d)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%v: add ", drink))
	for i, a := range add {
		buf.WriteString(fmt.Sprintf("%v", a))
		if i < len(add)-1 {
			buf.WriteString(fmt.Sprint(", "))
		} else {
			buf.WriteString(fmt.Sprint("\n"))
		}
	}
	fmt.Print(buf.String())
	time.Sleep(500 * time.Millisecond * time.Duration(len(add)))
	fmt.Printf("%v: ready to be served!\n", drink)
}

// func brewMachine(wg *sync.WaitGroup, b Brewer, d time.Duration, add ...drink.AddIn) {
//   wg.Add(1)
//   go func() {
//     b.Brew(d, add...)
//     wg.Done()
//   }()
// }

// func brew(b brewer, d time.Duration, add ...drink.AddIn) {
//   if b != nil {
//     panic("nil brewer")
//   }
//   b.Brew
// }

// ========================================

func newBrewerFunc(name string) brewer.BrewerFunc {
	return func(d time.Duration, add ...drink.AddIn) {

	}
}

func main() {

	// wg := new(sync.WaitGroup)

	// var d1 grain = "barley"
	// brewMachine(wg, d1, time.Second*5, drink.BrownSugar, drink.Lemon)

	// d2 := tea{name: "oolong", brand: "Kirin", weight: 2.50}
	// brewMachine(wg, d2, time.Second*2, drink.FreshMilk, drink.Ginger, drink.BrownSugar, drink.WhippedCream, drink.Matcha)

	// //convert the function into BrewerFunc so that it can be used as Brewer
	// d3 := BrewerFunc(func(d time.Duration, add ...drink.AddIn) {
	//   fmt.Printf("Name: %+v    Duration: %+v    Add-In: %+v\n", "arabica", d, add)
	// })
	// brewMachine(wg, d3, time.Second*4, drink.FreshMilk, drink.Butter, drink.Cocoa, drink.BrownSugar)

	// wg.Wait()

	barista := brewer.NewBarista()

	// var d1 grain = "barley"
	// d2 := tea{name: "oolong", brand: "Kirin", weight: 2.50}

	// barista.Assign(d1, time.Second*5, drink.BrownSugar, drink.Lemon)
	// barista.Assign(d2, time.Second*2, drink.FreshMilk, drink.Ginger, drink.BrownSugar, drink.WhippedCream, drink.Matcha)

	d3 := brewer.BrewerFunc(func(d time.Duration, add ...drink.AddIn) {
		brewOperation("arabica", d, add)
	})

	barista.Assign(d3, time.Second*4, drink.FreshMilk, drink.Butter, drink.Cocoa, drink.BrownSugar)

	barista.Brew()

	barista.Wait()

}
