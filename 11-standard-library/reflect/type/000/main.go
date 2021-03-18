package main

import (
	"fmt"
)

type counter int

func (c *counter) count() {
	(*c)++
	fmt.Println("count:", *c)
}

// ============================================================
type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending email to %v <%v>\n", u.name, u.email)
}

func sendNotification(n notifier) {
	n.notify()
}

// ============================================================

type A = int
type B int

func main() {

	// var y A
	// ta := reflect.TypeOf(y)
	// fmt.Println(ta)

	// var x B
	// tb := reflect.TypeOf(x)

	// ========================================

	// x := 9
	// y := 42

	// var p *int
	// p = &y

	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(*p)

	// ========================================

	u := user{"nikola tesla", "nikotes@gmail.com"}
	sendNotification(u)
	// sendNotification(&u)

}
