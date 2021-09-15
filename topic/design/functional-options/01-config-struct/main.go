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

func New(id int, opts playerOpts) player {
	return player{
		id: id,
		playerOpts: playerOpts{
			name:   opts.name,
			active: opts.active,
		},
	}
}

func main() {

	players := []player{
		New(1, playerOpts{"noob saibot", true}),
		New(2, playerOpts{"yoshi", false}),
		New(3, playerOpts{"", false}),
	}

	for _, p := range players {
		fmt.Printf("%+v\n", p)
	}

}
