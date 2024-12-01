package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	fmt.Println()
}

func main() {
	sum(78, 99, 13)
	nums := []int{11, 72, 33, 47}
	sum(nums...)
}
