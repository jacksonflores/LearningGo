package main

import "fmt"

func main() {
	p := struct {
		first string
		last  string
	}{
		first: "jackson",
		last:  "flores",
	}

	fmt.Println(p)

}
