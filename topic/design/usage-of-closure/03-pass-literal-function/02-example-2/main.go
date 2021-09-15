package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	{
		s := genIntSlice(0, 100, 10)
		sort.Ints(s)

		idx := sort.Search(len(s), func(i int) bool {
			return s[i] >= 50
		})
		fmt.Println(s)
		fmt.Printf("1st >= 50 | idx: %v | val: %v\n", idx, s[idx])
		fmt.Println()

	}

}

func genIntSlice(min, max, num int) []int {
	out := make([]int, num)
	for i := 0; i < num; i++ {
		out[i] = rand.Intn(max-min+1) + min
	}
	return out
}
