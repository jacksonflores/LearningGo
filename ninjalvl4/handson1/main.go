package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	// how to assing values to each position independently
	x[0] = 42
	x[1] = 420
	x[2] = 69
	x[3] = 400
	x[4] = 600
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)
}
