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

type PullStorer interface {
	Pull(d *Data) error
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

type System struct {
	// Xenia
	// Pillar
	Puller
	Storer
}

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, err := pull(ps, data)
		if err != nil {
			return err
		}
		if i > 0 {
			_, err := store(ps, data[:i])
			if err != nil {
				return err
			}
		}
	}
}

// =============================================================================

func main() {

	// sys := System{
	//   Xenia: Xenia{
	//     Host:    "localhost:8000",
	//     Timeout: time.Second,
	//   },
	//   Pillar: Pillar{
	//     Host:    "localhost:9000",
	//     Timeout: time.Second,
	//   },
	// }

	sys := System{
		Puller: &Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Storer: &Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	err := Copy(&sys, 3)
	if err != nil {
		fmt.Println(err)
	}

}
