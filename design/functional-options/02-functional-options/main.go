package main

import "fmt"

type player struct {
	id int
	playerOpts
}

type playerOpts struct {
	name   string
	active bool
}

func Name(name string) func(*player) {
	return func(p *player) {
		p.name = name
	}
}

func Active(active bool) func(*player) {
	return func(p *player) {
		p.active = active
	}
}

func New(id int, opts ...func(*player)) player {
	ret := player{
		id: id,
		playerOpts: playerOpts{
			name:   "annonymous",
			active: false,
		},
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
