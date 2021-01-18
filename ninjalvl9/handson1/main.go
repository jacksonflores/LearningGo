package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println("main routine")
	go yo()
	wg.Wait()

}

func yo() {
	fmt.Println("goroutinebaby")
	wg.Done()
}
