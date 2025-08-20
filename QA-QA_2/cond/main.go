package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(new(sync.Mutex))

	go func() {
		fmt.Println("Горутина 1. Жду")
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()

		fmt.Println("Горутина 1. Завершилась")
	}()

	go func() {
		fmt.Println("Горутина 2. Жду")
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()

		fmt.Println("Горутина 2. Завершилась")
	}()

	go func() {
		fmt.Println("Горутина 3. Жду")
		cond.L.Lock()
		cond.Wait()
		cond.L.Unlock()

		fmt.Println("Горутина 3. Завершилась")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Запуск горутин...")

	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(2 * time.Second)
}
