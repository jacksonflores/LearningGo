package main

import "fmt"

func main() {
	c1 := customEr{
		info: "bruhruhrurruhrurh",
	}

	foo(c1)
}

type customEr struct {
	info string
}

func (ce customEr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
}
