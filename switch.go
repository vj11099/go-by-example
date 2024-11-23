package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	fmt.Println()

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println()

	t := time.Now().Hour()

	switch true {
	case t < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	fmt.Println()

	whereAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("My type is %T\n", t)
		}
	}

	whereAmI(true)
	whereAmI(6)
	whereAmI("hello")
}
