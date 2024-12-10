package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int) {
	fmt.Printf("Worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
}
