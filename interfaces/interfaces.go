package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	shape         string
	width, height float64
}

type circle struct {
	shape  string
	radius float64
}

func (r *rect) area() float64 {
	return r.height * r.width
}

func (r *rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius

}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim(), "\n")
}

func main() {
	r := &rect{"Rectangle", 3.24, 9.78}
	c := &circle{"Circle", 8.88}

	measure(r)
	measure(c)
}
