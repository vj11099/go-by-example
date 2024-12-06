package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	messages <- "MESSAGE INCOMING"

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	signals <- true

	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	// case signals <- true:
	// 	fmt.Println("sent signal")
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no activity")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no signal")
	}
}
