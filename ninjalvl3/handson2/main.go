package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for ii := 0; ii <= 2; ii++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
