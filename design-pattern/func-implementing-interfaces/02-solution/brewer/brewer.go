package brewer

import (
	"sandbox/go-jottings/design-pattern/func-implementing-interfaces/drink"
	"sync"
	"time"
)

type Brewer interface {
	Brew(d time.Duration, a ...drink.AddIn)
}

type BrewerFunc func(d time.Duration, add ...drink.AddIn)

func (f BrewerFunc) Brew(d time.Duration, add ...drink.AddIn) {
	f(d, add...)
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
	b   Brewer
	d   time.Duration
	add []drink.AddIn
}

func newBrewEntry(b Brewer, d time.Duration, add ...drink.AddIn) brewEntry {
	addIns := make([]drink.AddIn, 0, len(add))
	addIns = append(addIns, add...)
	return brewEntry{b, d, addIns}
}

func (bar *Barista) Assign(b Brewer, d time.Duration, add ...drink.AddIn) {
	if b == nil {
		panic("nil Brewer")
	}
	newEntry := newBrewEntry(b, d, add...)
	bar.entries = append(bar.entries, newEntry)
}

// func (bar *Barista) AssignFunc(b BrewerFunc){
//   if b == nil {
//     panic("nil BrewerFunc")
//   }

// }

func (bar *Barista) Brew() {
	for _, entry := range bar.entries {
		bar.wg.Add(1)
		e := entry
		go func() {
			e.b.Brew(e.d, e.add...)
			bar.wg.Done()
		}()
	}
}

func (bar Barista) Wait() {
	bar.wg.Wait()
}
