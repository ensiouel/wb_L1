package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			counter.Add(1)
		}(&wg)
	}

	wg.Wait()

	fmt.Println(counter.Get())
}

type Counter struct {
	// Используем пакет atomic для создания конкурентного счетчика
	v atomic.Int64
}

func (counter *Counter) Add(delta int64) {
	counter.v.Add(delta)
}

func (counter *Counter) Get() int64 {
	return counter.v.Load()
}
