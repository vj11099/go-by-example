package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// for {
		// 	j, more := <-jobs
		// 	if more {
		// 		fmt.Println("recieved a job", j)
		// 	} else {
		// 		fmt.Println("recieved all jobs")
		// 		done <- true
		// 		return
		// 	}
		// }
		for i := 0; i < 5; i++ {
			j, more := <-jobs
			if (more) && (j < 4) {
				fmt.Println("recieved a job", j)
			} else {
				fmt.Println("recieved all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j < 6; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("sent all jobs")
	close(jobs)
	<-done

	_, ok := <-jobs
	fmt.Println("recieved more jobs:", ok)

}
