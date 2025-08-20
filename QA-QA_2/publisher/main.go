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
		batchCh   = make(chan []int)
		resultCh  = make(chan int)
		semaphore = make(chan struct{}, 2)
	)

	// Producer (Разбивает на батчи)
	go func() {
		batchSize := 5

		for i := 0; i < len(data); i += batchSize {
			end := i + batchSize
			if end > len(data) {
				end = len(data)
			}
			batchCh <- data[i:end]
		}

		close(batchCh)
	}()

	//Consumer (обработка с семафором)
	go func() {
		var wg sync.WaitGroup
		for batch := range batchCh {
			semaphore <- struct{}{}
			wg.Add(1)

			go func(b []int) {
				defer func() { <-semaphore }()
				defer wg.Done()

				var result int
				for _, item := range b {
					result += item
				}

				time.Sleep(1 * time.Second)

				resultCh <- result
			}(batch)
		}

		wg.Wait()
		close(resultCh)
	}()

	// Вывод результатов
	for res := range resultCh {
		fmt.Println(res)
	}
}
