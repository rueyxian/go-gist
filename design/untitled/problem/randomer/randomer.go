package randomer

import "math/rand"

type Randomer interface {
	Random(s rand.Source)
}

type RandomerFunc func(s rand.Source)

func (f RandomerFunc) Random(s rand.Source) {
	f(s)
}
