package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	timeBefore := time.Now()
	// withMutex()
	withoutMutex()
	// sequentialWithoutMutex()
	timeAfter := time.Now()
	fmt.Printf("The hours difference is: %f", timeAfter.Sub(timeBefore).Seconds())
}

func sequentialWithoutMutex() {
	var someMap = make(map[string]string, 0)

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(1) * time.Second)
		key := fmt.Sprintf("key-%d", i)
		someMap[key] = "value"
	}
}

func withoutMutex() {
	var someMap = make(map[string]string, 0)

	for i := 1; i <= 1000; i++ {
		go func(iteration int) {
			key := fmt.Sprintf("key-%d", iteration)
			someMap[key] = "value"
		}(i)
	}
}

func withMutex() {
	var (
		someMap      = make(map[string]string, 0)
		someMapMutex = sync.RWMutex{}
	)

	for i := 1; i <= 10000; i++ {
		go func(iteration int) {
			time.Sleep(time.Duration(1) * time.Second)
			key := fmt.Sprintf("key-%d", iteration)

			someMapMutex.Lock()
			someMap[key] = "value"
			someMapMutex.Unlock()
		}(i)
	}
}
