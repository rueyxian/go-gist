package main

import (
	"fmt"
	"go-gist/design/error-handling/05-behavior-as-context/config"
)

type cfg struct {
	User     string
	Password string
	Secure   bool
}

// ================================================================================

// Even though the temporary interface in config package is defined as unexported,
// we can define our own unexported temporary interface, so that we bring temporary
// behavior into the case statement.

type temporary interface {
	Temporary() bool
}

// ================================================================================

func main() {

	var c cfg
	_ = c
	// s := []string{"User=nova", "Password=4896", "Secure=true"}
	s := []string{"User=nova", "Drowssap=4896", "Secure=true"}

	// ==============================
	// if err := config.Parse(&c, s); err != nil {
	// if err := config.Parse(c, s); err != nil {
	// if err := config.Parse((*cfg)(nil), s); err != nil {
	if err := config.Parse((*cfg)(nil), s); err != nil {
		switch e := err.(type) {
		case temporary:
			// fmt.Printf("%s: %s\n", e.Name(), e)
			if !e.Temporary() {
				fmt.Println(e)
				return
			}
		default:
			fmt.Println(err)
		}
		return
	}

}
