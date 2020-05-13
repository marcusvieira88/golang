package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(2)
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	// GO reuses a thread to run many goroutines
	// GOROUTINE 1
	go func() {
		defer waitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Task Delayed")
	}()

	// GOROUTINE 2
	go func() {
		defer waitGrp.Done()

		fmt.Println("Normal Task")
	}()

	waitGrp.Wait()
}
