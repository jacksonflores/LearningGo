package main

import "fmt"

const (
	x = 2021 + iota
	y = 2021 + iota
	z = 2021 + iota
	w = 2021 + iota
)

func main() {
	fmt.Println(x, y, z, w)
}
