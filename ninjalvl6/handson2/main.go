package main

import "fmt"

func main() {
	xi := []int{1, 2, 3}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func bar(x []int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
