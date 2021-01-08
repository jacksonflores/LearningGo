package main

import "fmt"

func main() {
	switch {
	case !false:
		fmt.Println("1")
	case true:
		fmt.Println("2")
	case 1 == 2:
		fmt.Println("3")
	}
}
