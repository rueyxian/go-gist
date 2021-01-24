package brewer

import (
	"sync"
	"time"
)

type Brewer interface {
	Brew(d time.Duration, a []AddIn)
}

//====================

type AddIn struct {
	name string
}

//====================

type Barista struct {
	wg *sync.WaitGroup
	b Brewer
}

func NewBarista() Barista {
	return Barista{
		wg: &sync.WaitGroup{},
	}
}

func (

//====================
