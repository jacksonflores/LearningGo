package main

import "fmt"

func main() {
	x := foo
	x()()
}

func foo() func() {
	fmt.Println("----")
	return bar
}

func bar() {
	fmt.Println("bar says sup")
}
