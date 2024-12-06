package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Printf("Working...")
	time.Sleep(time.Second)
	fmt.Println("Finished!")
	done <- true

}

func main() {

	done := make(chan bool)
	go worker(done)
	<-done //waits for the channel task to complete
	go worker(done)
	<-done //waits for the channel task to complete
	fmt.Println("All tasks are completed succesfully")
}
