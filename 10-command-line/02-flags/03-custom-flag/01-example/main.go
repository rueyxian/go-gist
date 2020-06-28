package main

import (
	"flag"
	"fmt"
	"strings"
)

type mascots []string

var mascotsFlag mascots

func (m *mascots) String() string {
	return fmt.Sprint(*m)
}

func (m *mascots) Set(value string) error {
	*m = strings.Split(value, ",")
	return nil
}

func init() {
	mascotsFlag = mascots([]string{"gopher", "tux", "gnu", "ferris"})

	/*
	   func Var(value Value, name string, usage string)

	   type Value interface {
	       String() string
	       Set(string) error
	   }
	*/
	flag.Var(&mascotsFlag, "mascots", "computing mascots <3")
	flag.Var(&mascotsFlag, "m", "computing mascots <3")
	flag.Parse()
}

func main() {

	fmt.Printf("mascots: %v", mascotsFlag)

}
