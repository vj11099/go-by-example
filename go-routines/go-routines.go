package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func n(msg string) {
	fmt.Println(msg)
}

func main() {
	f("direct")

	go f("go-routine")

	go n("thread interval")

	time.Sleep(time.Second)
	fmt.Println("done")
}
