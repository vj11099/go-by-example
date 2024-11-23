package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	//recursion to get the factorial of n
	fmt.Println("Factorial of 7: ", fact(7))

	//recursion to get n-th value of fibonacci series
	var fib func(n int) int
	fib = func(num int) int {
		if num < 2 {
			return num
		}
		return fib(num-1) + fib(num-2)
	}
	fmt.Println("\n7th fibonacci number:", fib(7))
}
