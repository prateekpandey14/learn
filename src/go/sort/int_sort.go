package main

import (
	"fmt"
	"sort"
)

func main() {

	list := []int{32, 5, 14, 23, 10}

	// Sort sorts data.
	// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
	// data.Less and data.Swap. The sort is not guaranteed to be stable.

	sort.Ints(list)
	fmt.Println(list)
}
