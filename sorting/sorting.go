package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"y", "m", "k"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	ints := []int{9, 3, 5, 0, 12}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	s := slices.IsSorted(ints)
	fmt.Println("Sorted status:", s)
}
