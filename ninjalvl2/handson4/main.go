package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%d\t%b\t%#x", x, x, x)
	fmt.Printf("\n")
	y := x << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}
