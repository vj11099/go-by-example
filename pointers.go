package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial value: ", i)
	fmt.Println()

	zeroVal(i)
	fmt.Println("After updating value in zeroVal func: ", i)
	fmt.Println()

	zeroPtr(&i)
	fmt.Println("After pointing i's address to value 0: ", i)
	fmt.Println()

	fmt.Println("Address/Pointer of i", &i)
	fmt.Println("")
}
