package main

import "fmt"

func main() {
	defer foo()

}

func foo() {
	fmt.Println("foo says hello")

	defer func() {
		fmt.Println("foo defered yo")
	}()
}
