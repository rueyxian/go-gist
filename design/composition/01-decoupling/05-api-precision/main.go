package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

// =============================================================================
// ========== primitive level api (method base api) ==========
type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {
	r := rand.Intn(10)
	switch r {
	case 1, 9:
		return io.EOF
	case 5:
		return errors.New("Error reading data from Xenia")
	default:
		d.Line = fmt.Sprint("Data ", r)
		fmt.Println("In:", d.Line)
		return nil
	}
}

type Pillar struct {
	Host    string
	Timeout time.Duration
}

func (*Pillar) Store(d *Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// =============================================================================
// ========== lower level api (function base api) ==========

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}
	return len(data), nil
}

// =============================================================================
// ========== higher level api ==========

// type System struct {
//   Puller
//   Storer
// }

// func Copy(sys *System, batch int) error {
//   data := make([]Data, batch)
//   for {
//     i, err := pull(sys, data)
//     if err != nil {
//       return err
//     }
//     if i > 0 {
//       _, err := store(sys, data[:i])
//       if err != nil {
//         return err
//       }
//     }
//   }
// }

// although data injection is cool, but it doesn't make the api easier to understand
// this is much more precise
func Copy(p Puller, s Storer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(p, data)
		if err != nil {
			return err
		}
		if i > 0 {
			_, err := store(s, data[:i])
			if err != nil {
				return err
			}
		}
	}
}

// =============================================================================

func main() {

	// sys := System{
	//   Puller: &Xenia{
	//     Host:    "localhost:8000",
	//     Timeout: time.Second,
	//   },
	//   Storer: &Pillar{
	//     Host:    "localhost:9000",
	//     Timeout: time.Second,
	//   },
	// }

	x := &Xenia{
		Host:    "localhost:8000",
		Timeout: time.Second,
	}

	p := &Pillar{
		Host:    "localhost:9000",
		Timeout: time.Second,
	}

	err := Copy(x, p, 3)
	if err != nil {
		fmt.Println(err)
	}

}
