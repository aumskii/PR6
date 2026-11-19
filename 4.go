package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var (
	mu    sync.Mutex
	votes [3]int
)

func vote(i int) {
	mu.Lock()
	votes[i]++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				i := int(rand.Float64() * float64(len(votes)))
				vote(i)
			}
		}()
	}
	wg.Wait()

	total := 0
	for i, n := range votes {
		fmt.Printf("Кандидат %d: %d\n", i+1, n)
		total += n
	}
	fmt.Println("Всего голосов:", total)
}
