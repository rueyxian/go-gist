/*
https://github.com/go-yaml/yaml
*/

package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to correctly populate the data
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func main() {
	t := T{}
	m := make(map[interface{}]interface{})

	//========================================
	// yaml -> struct
	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- t:\n%+v \n\n", t)

	//========================================
	// struct -> yaml
	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- t dump:\n%s...\n\n", d)

	//========================================
	// yaml -> map
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- m:\n%+v \n\n", m)

	//========================================
	// map -> yaml
	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("--- m dump:\n%s...\n\n", d)

}
