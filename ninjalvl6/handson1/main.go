package main

import "fmt"

func main() {
	i := foo()
	fmt.Println(i)

	x, y := bar()
	fmt.Println(x, y)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 69, "big money silvia"
}
