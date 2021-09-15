package brewer

import (
	"sync"
	"time"
)

type Brewer interface {
	Brew(d time.Duration, a ...AddIn)
}

type BrewerFunc func(d time.Duration, add ...AddIn)

func (f BrewerFunc) Brew(d time.Duration, add ...AddIn) {
	f(d, add...)
}

type AddIn struct {
	add string
}

func AddInsFunc(adds ...string) func() []AddIn {
	return func() []AddIn {
		ret := make([]AddIn, 0, len(adds))
		for _, a := range adds {
			ret = append(ret, AddIn{a})
		}
		return ret
	}
}

type Barista struct {
	wg      *sync.WaitGroup
	entries []brewEntry
}

func NewBarista() Barista {
	return Barista{
		wg: new(sync.WaitGroup),
	}
}

type brewEntry struct {
	b    Brewer
	d    time.Duration
	adds []AddIn
}

func newBrewEntry(b Brewer, d time.Duration, a []AddIn) brewEntry {
	return brewEntry{b, d, a}
}

func (bar *Barista) Assign(b Brewer, d time.Duration, f func() []AddIn) {
	if b == nil {
		panic("nil Brewer")
	}
	newEntry := newBrewEntry(b, d, f())
	bar.entries = append(bar.entries, newEntry)
}

func (bar *Barista) Brew() {
	for _, entry := range bar.entries {
		bar.wg.Add(1)
		e := entry
		go func() {
			e.b.Brew(e.d, e.adds...)
			bar.wg.Done()
		}()
	}
}

func (bar Barista) Wait() {
	bar.wg.Wait()

