package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func writeLog(id, n int) {
	mu.Lock()
	fmt.Printf("Горутина %d: сообщение %d\n", id, n)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := 1; n <= 3; n++ {
				writeLog(id, n)
			}
		}(i)
	}
	wg.Wait()
}
