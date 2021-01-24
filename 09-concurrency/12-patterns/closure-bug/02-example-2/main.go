package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rnd *rand.Rand

func init() {
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// ============================================================

func main() {

	ch1 := pipe("AAA", "BBB", "CCC", "DDDD")
	ch2 := pipe("@@", "##", "$$", "&&")
	ch3 := pipe("1", "2", "3", "4")

	// ====================
	// for v := range fanIn_bug(ch1, ch2, ch3) {
	//   fmt.Println(v)
	// }

	// ====================
	// for v := range fanIn_fix1(ch1, ch2, ch3) {
	//   fmt.Println(v)
	// }

	// ====================
	for v := range fanIn_fix2(ch1, ch2, ch3) {
		fmt.Println(v)
	}

}

// ============================================================
func fanIn_bug(chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	ret := make(chan string)
	go func() {
		defer close(ret)
		for _, ch := range chs {
			go func() {
				defer wg.Done()
				for str := range ch {
					ret <- str
				}
			}()
		}
		wg.Wait()
	}()
	return ret
}

// ============================================================
func fanIn_fix1(chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	ret := make(chan string)
	go func() {
		defer close(ret)
		for _, ch := range chs {
			go func(c <-chan string) {
				defer wg.Done()
				for str := range c {
					ret <- str
				}
			}(ch)
		}
		wg.Wait()
	}()
	return ret
}

// ============================================================
func fanIn_fix2(chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	ret := make(chan string)
	go func() {
		defer close(ret)
		for i, _ := range chs {
			ch := chs[i]
			go func() {
				defer wg.Done()
				for str := range ch {
					ret <- str
				}
			}()
		}
		wg.Wait()
	}()
	return ret
}

// ============================================================

func pipe(strs ...string) <-chan string {
	ret := make(chan string)
	go func() {
		defer close(ret)
		for _, str := range strs {
			ret <- str
		}
	}()
	return ret
}

// ============================================================
