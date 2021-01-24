package main

import (
	"fmt"
	"math/rand"
	"sandbox/go-jottings/design/untitled/problem/randomer"
	"time"
)

type Displayer interface {
	Display()
}

// ====================
// []int type

type numbers []int

func (n numbers) Random(s rand.Source) {
	r := rand.New(s)
	for i, j := range r.Perm(len(n)) {
		n[i], n[j] = n[j], n[i]
	}
}

func (n numbers) Display() {
	fmt.Print("numbers: ")
	for _, v := range n {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func newNumbers(n int) numbers {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return numbers(nums)
}

// ====================
// struct type

const (
	minDiceScore = 1
	maxDiceScore = 6
)

type die struct {
	id    int
	score int
}

type dice struct {
	d []die
}

func newDice(n int) dice {
	ret := dice{
		d: make([]die, n),
	}
	for i := 0; i < n; i++ {
		ret.d[i].id = i
	}
	return ret
}

func (d dice) Random(s rand.Source) {
	rnd := rand.New(s)
	for i, _ := range d.d {
		d.d[i].score = rnd.Intn(maxDiceScore) + minDiceScore
	}
}

func (d dice) Display() {
	for _, v := range d.d {
		fmt.Printf("dice: %v, score: %v\n", v.id, v.score)
	}
	fmt.Println()
}

// ====================
// func type
// what if we want to implement an interface on func type?

// type randomerFunc func()

// func (f randomerFunc) Random(s rand.Source) {
//   f()
// }

// func newRandomFunc(nums []int) randomerFunc {
//   return func() {
//     r := rand.New(s)
//     for i, j := range r.Perm(len(n)) {
//       n[i], n[j] = n[j], n[i]
//     }
//   }
// }

func newRandomerFunc(data []string) randomer.RandomerFunc {
	return func(s rand.Source) {
		r := rand.New(s)
		for i, j := range r.Perm(len(data)) {
			data[i], data[j] = data[j], data[i]
		}
	}
}

// ====================

func handleRandomer(rdm randomer.Randomer) {
	s := rand.NewSource(time.Now().UnixNano())
	rdm.Random(s)
}

// func handleRandomers(rdms []randomer.Randomer) {
//   s := rand.NewSource(time.Now().UnixNano())
//   for _, r := range rdms {
//     r.Random(s)
//   }
// }

// ====================

func main() {

	r1 := newNumbers(9)
	r2 := newDice(3)

	s := []string{"Anna", "Bob", "Charlie", "Daniel", "Emma", "Fiona", "Gorge", "Harry", "Ian", "Jack", "Kay"}
	r3 := newRandomerFunc(s)

	handleRandomer(r1)
	handleRandomer(r2)
	handleRandomer(r3)

	fmt.Printf("%+v\n", r1)
	fmt.Printf("%+v\n", r2)
	fmt.Printf("%+v\n", s)

}
