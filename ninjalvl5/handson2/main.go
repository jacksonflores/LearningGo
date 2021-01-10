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
	fmt.Println(p1, "\n", p2)

	m := make(map[string]person)
	m[p1.second] = p1
	m[p2.second] = p2
	fmt.Println(m[p1.second])

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.second)
		for _, val := range v.cream {
			fmt.Println(val)
		}
		fmt.Println("-----")
	}
}
