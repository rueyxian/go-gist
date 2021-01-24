package main

import "fmt"

type player struct {
	id     int
	name   string
	active bool
}

type option func(*player)

func Name(name string) option {
	return func(p *player) {
		p.name = name
	}
}

func Active(active bool) option {
	return func(p *player) {
		p.active = active
	}
}

func New(id int, opts ...option) player {
	ret := player{
		id:     id,
		name:   "annonymous",
		active: false,
	}
	for _, opt := range opts {
		opt(&ret)
	}
	return ret
}

func main() {

	players := []player{
		New(1, Name("noob saibot"), Active(true)),
		New(2, Name("yoshi")),
		New(3),
	}

	for _, p := range players {
		fmt.Printf("%+v\n", p)
	}

}
