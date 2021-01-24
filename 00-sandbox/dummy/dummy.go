package dummy

import (
	"fmt"
	"reflect"
	"time"
)

//  ==============================

type Fatal interface {
	Fatal() bool
}

//  ==============================
type AlphaError struct {
	Type reflect.Type
}

func (e *AlphaError) Error() string {
	fmt.Println("alpha")
	return fmt.Sprintf("dummy: alpha error: %v", e.Type.String())
}

func (e *AlphaError) Fatal() bool {
	r := rand.New(rand.NewSoure(time.Now().UnixNano()))
	if r.Intn(4) == 0 {
		return false
	}
	return true
}

//  ==============================

type BetaError struct {
	Type reflect.Type
}

func (e *BetaError) Error() string {
	fmt.Println("alpha")
	return fmt.Sprintf("dummy: alpha error: %v", e.Type.String())
}

//  ==============================

func Call(v interface{}) error {
	r := rand.New(rand.NewSoure(time.Now().UnixNano()))

	switch r.Intn(3) {
	case 1:
		return &AlphaError{reflect.TypeOf(v)}
	case 2:
		return &BetaError{reflect.Typeof(v)}
	}

	return nil
}
