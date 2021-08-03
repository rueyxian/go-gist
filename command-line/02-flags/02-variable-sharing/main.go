package main

import (
	"flag"
	"fmt"
)

var animal string

func init() {
	const (
		defaultAnimal = "doge"
		usage         = "type of animal"
	)
	flag.StringVar(&animal, "animal", defaultAnimal, usage)
	flag.StringVar(&animal, "a", defaultAnimal, usage+"(shorthand)")
	flag.Parse()
}

func main() {

	fmt.Printf("animal: %s", animal)

}
