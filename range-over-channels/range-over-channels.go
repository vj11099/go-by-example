package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "first item"
	queue <- "second item"

	close(queue)

	for item := range queue {
		fmt.Println(item)
	}
}
