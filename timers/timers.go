package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Timer 1 started")

	<-timer1.C
	fmt.Println("Timer 1 finished")

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("Timer 2 started")
		<-timer2.C
		fmt.Println("Timer 2 finished")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
