package main

import "fmt"

func add(a, b, c int) int {
	return a + b + c
}

func vals() (int, int, int) {
	return 10, 13, 49
}

func main() {

	a, b, c := vals()
	fmt.Println("a from vals:", a)
	fmt.Println("b from vals:", b)
	fmt.Println("c from vals:", c)

	res := add(10, 13, 49)
	fmt.Println("sum:", res)
}
