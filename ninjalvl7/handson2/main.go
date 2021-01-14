package main

import "fmt"

func main() {
	p1 := person{
		first: "jackson",
		last:  "flores",
	}
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	p.first = "blast"
}
