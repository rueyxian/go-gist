package main

import "fmt"

// user defines a user in the system.
type user struct {
	name  string
	email string
}

// String implements the fmt.Stringer interface.
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

// func (u user) String() string {
//   return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
// }

func main() {

	// Create a value of type user.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Display the values.
	fmt.Println(u)  //{Bill bill@ardanlabs.com}
	fmt.Println(&u) //My name is "Bill" and my email is "bill@ardanlabs.com"
}
