package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func main() {
	res := add(10, 13, 49)

	fmt.Println("sum:", res)
}
