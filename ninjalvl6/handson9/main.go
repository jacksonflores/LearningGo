package main

import "fmt"

func main() {
	foo(bar)
}

func foo(f func()) {
	f()
}

func bar() {
	fmt.Println("whaddup ball licker")
}
