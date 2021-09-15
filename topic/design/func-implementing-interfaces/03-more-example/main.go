package main

import (
	"fmt"
	"math/rand"
	"sandbox/go-jottings/design/func-implementing-interfaces/03-more-example/shuffler"
	"time"
)

// ========================================

type numbers []int

func (n numbers) Shuffle(s rand.Source) {
	r := rand.New(s)
	for i, j := range r.Perm(len(n)) {
		n[i], n[j] = n[j], n[i]
	}
}

// ========================================

type block uint8

const (
	blockI block = iota
	blockO
	blockT
	blockS
	blockZ
	blockJ
	blockL
)

const (
	minBlock = blockI
	maxBlock = blockL
)

type bag struct {
	index  int
	blocks []block
}

func newBag(idx int) bag {
	ret := bag{
		index:  idx,
		blocks: []block{blockI, blockO, blockT, blockS, blockZ, blockJ, blockL},
	}
	return ret
}

func (b *bag) Shuffle(s rand.Source) {
	r := rand.New(s)
	for i, j := range r.Perm(len(b.blocks)) {
		b.blocks[i], b.blocks[j] = b.blocks[j], b.blocks[i]
	}
}

// ========================================

func newShufflerFunc(nums ...int) shuffler.ShufflerFunc {
	return func(s rand.Source) {
		r := rand.New(s)
		for i, j := range r.Perm(len(nums)) {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}

// ========================================

func handleShuffle(source rand.Source, shufflers ...shuffler.Shuffler) {
	for _, s := range shufflers {
		s.Shuffle(source)
	}
}

// ========================================

func main() {

	source := rand.NewSource(time.Now().UnixNano())

	s1 := numbers{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := newBag(3)

	n := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
	s3 := newShufflerFunc(n...)

	handleShuffle(source, s1, &s2, s3)
	fmt.Println(s1)
	fmt.Printf("%+v\n", s2)
	fmt.Println(n)

}
