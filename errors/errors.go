package main

import (
	"errors"
	"fmt"
)

func f(i int) (int, error) {
	if i == 43 {
		return -1, errors.New("cant work with 43")
	}
	return i + 3, nil
}

var ErrNoCoffee = fmt.Errorf("we need to buy more coffee")
var ErrPower = fmt.Errorf("cant make coffee with no power =(")

func makeCoffee(i int) error {
	if i == 2 {
		return ErrNoCoffee
	} else if i == 4 {
		return fmt.Errorf("making coffee... %w", ErrPower)
	}
	return nil
}

func main() {
	for _, e := range []int{10, 43} {
		if r, e := f(e); e != nil {
			fmt.Println("Task failed: ", e)
		} else {
			fmt.Println("Task complete", r)
		}
	}

	for i := range 5 {
		if err := makeCoffee(i); err != nil {
			if errors.Is(err, ErrNoCoffee) {
				fmt.Println("We should get more coffee!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("It is dark now")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Your coffee is ready!")
	}
}
