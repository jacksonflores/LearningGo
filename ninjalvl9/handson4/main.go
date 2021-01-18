package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex sync.Mutex

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mutex.Lock()
			v := counter
			// time.Sleep(time.Second)
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)
}
