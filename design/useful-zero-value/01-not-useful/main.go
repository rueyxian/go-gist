package main

import "fmt"

type Payload struct {
	value string
}

func (p *Payload) Value() string {
	return p.value
}

func main() {

	p1 := &Payload{"blueberry"}
	var p2 *Payload

	fmt.Println(p1.Value())
	fmt.Println(p2.Value()) //this will fail

}
