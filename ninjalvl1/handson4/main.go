package main

import "fmt"

type cum int

func main() {
	var x cum
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

}
