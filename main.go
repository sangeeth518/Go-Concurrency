package main

import (
	"fmt"
	"sync"
	"time"
)

func new(wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("studing goroutines and channels")
	}
	fmt.Println("task done")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("hello")
	go new(&wg)
	// time.Sleep(6 * time.Second)
	wg.Wait()

	fmt.Println("main exit")
}
