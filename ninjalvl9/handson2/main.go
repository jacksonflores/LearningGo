package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("yo im", p.name, "and im", p.age, "years old")
}

type human interface {
	speak()
}

func main() {
	p1 := person{
		name: "jackson",
		age:  17,
	}
	// following works because this is of type pointer to a person
	saySomething(&p1)
	// following does not work because it is not of type pointer to a person
	// method speak can only take type pointer to a person
	// saySomething(p1)
}

func saySomething(h human) {
	h.speak()
}
