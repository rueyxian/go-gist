package main

import (
	"fmt"
	"sort"
)

func main() {

	{
		s := []string{"oliver", "jake", "maxwell", "logan", "eli", "benjamin", "anna"}
		sort.Slice(s, func(i, j int) bool {
			return len(s[i]) < len(s[j])
		})
		fmt.Println(s)
	}

	{
		s := []string{"oliver", "jake", "maxwell", "logan", "eli", "benjamin", "anna"}
		fn := func(i, j int) bool {
			return len(s[i]) < len(s[j])
		}
		sort.Slice(s, fn)
		fmt.Println(s)
	}

}
