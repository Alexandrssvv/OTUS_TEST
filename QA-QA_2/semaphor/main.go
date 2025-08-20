package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	data := make([]int, 100)
	for i := range data {
		data[i] = i
	}

	processData(data)
}

func processData(data []int) {
	var (
		wg        sync.WaitGroup
		semaphore = make(chan struct{}, 5)
	)

	wg.Add(len(data))

	for i := range data {
		semaphore <- struct{}{}

		go func() {
			defer func() { <-semaphore }()
			defer wg.Done()

			fmt.Println(data[i])
			time.Sleep(2 * time.Second)
		}()
	}

	wg.Wait()

}
