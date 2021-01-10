package main

import "fmt"

type person struct {
	first  string
	second string
	cream  []string
}

func main() {
	p1 := person{
		first:  "sammy",
		second: "coom",
		cream:  []string{"cum", "piss", "blood"},
	}
	p2 := person{
		first:  "jackson",
		second: "florenzes",
		cream:  []string{"vanilla", "barf", "beef"},
	}

	fmt.Println(p1.first, p1.second)
	for i, v := range p1.cream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.second)
	for i, v := range p2.cream {
		fmt.Println(i, v)

	}
}
