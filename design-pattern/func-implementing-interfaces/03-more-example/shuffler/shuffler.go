package shuffler

import "math/rand"

type Shuffler interface {
	Shuffle(s rand.Source)
}

type ShufflerFunc func(s rand.Source)

func (f ShufflerFunc) Shuffle(s rand.Source) {
	f(s)
}
