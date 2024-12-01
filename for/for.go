package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println("I-th iteration", i)
		i++
	}
	fmt.Println()

	for j := 0; j < 3; j++ {
		fmt.Println("J-th iteration", j)
	}
	fmt.Println()

	for i := range 3 {
		fmt.Println("range", i)
	}
	fmt.Println()

	for {
		fmt.Println("Loop")
		break
	}
	fmt.Println()

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	fmt.Println()

}
