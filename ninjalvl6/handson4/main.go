package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "jackson",
		last:  "flores",
		age:   17,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
