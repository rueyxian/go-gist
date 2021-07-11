package main

import "fmt"

type dumber interface {
	dumb(int)
}

// ============================================================

type stupid string

func (s stupid) dumb(n int) {
	fmt.Printf("%s: dumb-o-meter: %d\n", s, n)
}

// ============================================================

type wrapstupid struct {
	stupid
}

// func (s wrapstupid) dumb(n int) {
//   s.stupid.dumb(n)
// fmt.Printf("%s (wrapped): dumb-o-meter: %d\n", s.stupid, n)
// }

// ============================================================

func rundumber(d dumber, n int) {
	d.dumb(n)
}

// ============================================================

func main() {

	s1 := stupid("jellytea")
	s2 := wrapstupid{stupid("milktea")}

	rundumber(s1, 99)
	rundumber(s2, 77)

}
