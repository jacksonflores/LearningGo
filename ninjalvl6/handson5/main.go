package main

import (
	"fmt"
	"math"
)

func main() {
	s1 := square{
		len:   3,
		width: 4,
	}
	c1 := circle{
		radius: 5,
	}
	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", c1)

	info(s1)
	info(c1)
}

type shape interface {
	area() float64
}

type square struct {
	len   float64
	width float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.len * s.width
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}
