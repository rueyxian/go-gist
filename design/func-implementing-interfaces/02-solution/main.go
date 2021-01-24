package main

import (
	"bytes"
	"fmt"
	"sandbox/go-jottings/design/func-implementing-interfaces/02-solution/brewer"
	"time"
)

// ========================================
type grain string

func (g grain) Brew(d time.Duration, add ...brewer.AddIn) {
	brewOperation(string(g), d, add)
}

// ========================================
type tea struct {
	name   string
	brand  string
	weight float64
}

func (t tea) Brew(d time.Duration, add ...brewer.AddIn) {
	brewOperation(string(t.name), d, add)
}

// ========================================
func newBrewerFunc(name string) brewer.BrewerFunc {
	return func(d time.Duration, add ...brewer.AddIn) {
		brewOperation(name, d, add)
	}
}

// ========================================
func brewOperation(drink string, d time.Duration, add []brewer.AddIn) {
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
	d3 := brewer.BrewerFunc(func(d time.Duration, add ...brewer.AddIn) {
		brewOperation("arabica", d, add)
	})

	barista.Assign(d1, time.Second*5, brewer.AddInsFunc("BrownSugar", "Lemon"))
	barista.Assign(d2, time.Second*2, brewer.AddInsFunc("FreshMilk", "Ginger", "BrownSugar", "WhippedCream", "Matcha"))
	barista.Assign(d3, time.Second*4, brewer.AddInsFunc("FreshMilk", "Butter", "Cocoa", "BrownSugar"))
	barista.Assign(newBrewerFunc("green tea"), time.Second*3, brewer.AddInsFunc("Honey", "Lemon", "PapermintOil"))

	barista.Brew()

	barista.Wait()
	fmt.Println("FINISH!!!")

}
