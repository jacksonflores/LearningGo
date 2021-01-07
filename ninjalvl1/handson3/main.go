package main

import "fmt"

// gave vars x,y,z default values because not specified. also known as zero values :D

var x int
var y string
var z bool

func main() {
	fmt.Println(x, y, z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
