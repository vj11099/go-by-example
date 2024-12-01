package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("Un-initialized slice:", s, ", Length:", len(s), ", Is nil:", s == nil)
	fmt.Println()

	s = make([]string, 3)
	fmt.Println("Restructured slice:", s, ", Length:", len(s), ", Capacity:", cap(s))
	fmt.Println()

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("New slice:", s)
	fmt.Println()
	fmt.Println("Slice at index 2:", s[2])
	fmt.Println()

	fmt.Println("Slice Length:", len(s))
	fmt.Println()

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Appended Slice:", s)
	fmt.Println()

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied Slice:", c)
	fmt.Println()

	l := s[2:5]
	fmt.Println("Slice from idx 2 to 5:", l)
	fmt.Println()

	l = s[:5]
	fmt.Println("Slice till idx 5 but not 5:", l)
	fmt.Println()

	l = s[2:]
	fmt.Println("Slice from idx 2:", l)
	fmt.Println()

	t1 := []string{"a", "b", "c"}
	fmt.Println("Declared slice:", t1)
	fmt.Println()

	t2 := []string{"a", "b", "c"}
	if slices.Equal(t1, t2) {
		fmt.Println("t1 === t2")
	}
	fmt.Println()

	twoD := make([][]int, 3)

	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2-D Array:", twoD)
	fmt.Println()

}
