package main

import (
	"bytes"
	"fmt"
	"sandbox/go-jottings/design-pattern/func-implementing-interfaces/02-solution/brewer"
	"sandbox/go-jottings/design-pattern/func-implementing-interfaces/drink"
	"time"
)

// ========================================
type grain string

func (drink grain) Brew(d time.Duration, add ...drink.AddIn) {
	brewOperation(string(drink), d, add)
}

// ========================================
type tea struct {
	name   string
	brand  string
	weight float64
}

func (t tea) Brew(d time.Duration, add ...drink.AddIn) {
	brewOperation(string(t.name), d, add)
}

// ========================================
func newBrewerFunc(name string) brewer.BrewerFunc {
	return func(d time.Duration, add ...drink.AddIn) {
		brewOperation(name, d, add)
	}
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

// ========================================

func main() {

	barista := brewer.NewBarista()

	var d1 grain = "barley"
	d2 := tea{name: "oolong", brand: "Kirin", weight: 2.50}
	d3 := brewer.BrewerFunc(func(d time.Duration, add ...drink.AddIn) {
		brewOperation("arabica", d, add)
	})

	barista.Assign(d1, time.Second*5, drink.BrownSugar, drink.Lemon)
	barista.Assign(d2, time.Second*2, drink.FreshMilk, drink.Ginger, drink.BrownSugar, drink.WhippedCream, drink.Matcha)
	barista.Assign(d3, time.Second*4, drink.FreshMilk, drink.Butter, drink.Cocoa, drink.BrownSugar)
	barista.Assign(newBrewerFunc("green tea"), time.Second*3, drink.Honey, drink.Lemon, drink.PapermintOil)

	barista.Brew()

	barista.Wait()
	fmt.Println("FINISH!!!")

}
