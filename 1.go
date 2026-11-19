package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.Mutex
	views int
)

func addView() {
	mu.Lock()
	views++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				addView()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Просмотров:", views)
}
