package main

import "fmt"

func main() {
	//sum using range
	arr := []int{2, 3, 4}
	sum := 0
	for _, n := range arr {
		sum += n
	}
	fmt.Println("Sum: ", sum)
	fmt.Println("")

	//index using range
	for i, n := range arr {
		if n == 3 {
			fmt.Println("index of 3:", i)
		}
	}
	fmt.Println("")

	//key value pairs in maps using range
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println("")

	//string and indexes using range
	for i, c := range "go-lang" {
		fmt.Println(i, "->", c)
	}
}
