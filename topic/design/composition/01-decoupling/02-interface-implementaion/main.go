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

// both pull & store function, instead of accepting concrete data, change it interface

// func pull(x *Xenia , data []Data) (int, error) {
//   for i := range data {
//     if err := x.Pull(&data[i]); err != nil {
//       return i, err
//     }
//   }
//   return len(data), nil
// }

// func store(p *Pillar, data []Data) (int, error) {
//   for i := range data {
//     if err := p.Store(&data[i]); err != nil {
//       return i, err
//     }
//   }
//   return len(data), nil
// }

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
	Xenia
	Pillar
}

// func Copy(sys *System, batch int) error {
//   data := make([]Data, batch)
//   for {
//     i, err := pull(&sys.Xenia, data)
//     if err != nil {
//       return err
//     }
//     if i > 0 {
//       _, err := store(&sys.Pillar, data[:i])
//       if err != nil {
//         return err
//       }
//     }
//   }
// }

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

	//the code still not completely decoupled yet.
	//System type is embedded concrete types (Xenia & Pillar)
	//what if we wanted different combination in the future?
	sys := System{
		Xenia: Xenia{
			Host:    "localhost:8000",
			Timeout: time.Second,
		},
		Pillar: Pillar{
			Host:    "localhost:9000",
			Timeout: time.Second,
		},
	}

	err := Copy(&sys, 3)
	if err != nil {
		fmt.Println(err)
	}

}
