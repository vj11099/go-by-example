package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	buffer := time.Tick(200 * time.Millisecond)

	for i := 0; i < 5; i++ {
		requests <- i + 1
	}

	close(requests)

	for items := range requests {
		<-buffer
		fmt.Println("request:", items, time.Now())
	}

	burstyBuffer := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyBuffer <- time.Now()
	}

	go func() {
		for req := range time.Tick(500 * time.Millisecond) {
			burstyBuffer <- req
		}
	}()

	// close(burstyBuffer)

	burstyRequests := make(chan int, 5)

	for i := 0; i < 5; i++ {
		burstyRequests <- i + 1
	}

	close(burstyRequests)

	for items := range burstyRequests {
		<-burstyBuffer
		fmt.Println("bursty request:", items, time.Now())
	}

}
