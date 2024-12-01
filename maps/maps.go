package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 9
	m["k2"] = 11
	fmt.Println("Map: ", m)
	fmt.Println()

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println()

	v3 := m["k3"]
	fmt.Println("v3: ", v3)
	fmt.Println()

	delete(m, "k2")
	fmt.Println("Map:", m)
	fmt.Println()

	clear(m)
	fmt.Println("Map:", m)
	fmt.Println()

	_, prs := m["k1"]
	fmt.Println("pairs:", prs)
	fmt.Println()

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map n:", n)
	fmt.Println()

	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Map n1:", n1)
	fmt.Println()

	if maps.Equal(n, n1) {
		fmt.Println("n == n1")
		fmt.Println()
	}
}
