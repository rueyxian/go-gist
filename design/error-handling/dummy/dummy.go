package dummy

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//  ==============================

type temporary interface {
	Temporary() bool
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(3) == 0 {
		return false
	}
	return true
}

//  ==============================

type BetaError struct {
	Type reflect.Type
}

func (e *BetaError) Error() string {
	fmt.Println("beta")
	return fmt.Sprintf("dummy: beta error: %v", e.Type.String())
}

func (e *BetaError) Fatal() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(3) == 0 {
		return false
	}
	return true
}

//  ==============================

type GammaError struct {
	Type reflect.Type
}

func (e *GammaError) Error() string {
	fmt.Println("gamma")
	return fmt.Sprintf("dummy: gamma error: %v", e.Type.String())
}

func (e *GammaError) Fatal() bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(3) == 0 {
		return false
	}
	return true
}

//  ==============================

func Call(v interface{}) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch r.Intn(4) {
	case 1:
		return &AlphaError{reflect.TypeOf(v)}
	case 2:
		return &BetaError{reflect.TypeOf(v)}
	case 3:
		return &GammaError{reflect.TypeOf(v)}
	default:
	}
	return nil

}

//  ==============================
