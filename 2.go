package main

import (
	"fmt"
	"sync"
)

var (
	mu    sync.RWMutex
	cache = make(map[string]int)
)

func get(key string) (int, bool) {
	mu.RLock()
	defer mu.RUnlock()
	value, ok := cache[key]
	return value, ok
}

func set(key string, value int) {
	mu.Lock()
	cache[key] = value
	mu.Unlock()
}

func main() {
	set("alice", 31)

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value, ok := get("alice")
			if ok {
				fmt.Println("alice:", value)
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		set("alice", 32)
	}()

	wg.Wait()
	value, ok := get("alice")
	if ok {
		fmt.Println("После обновления:", value)
	}
}
