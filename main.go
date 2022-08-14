package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 4, 2, 5, 3}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	fmt.Print(a)

}
