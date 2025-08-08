package main

import (
	"fmt"
	"sync"
)

func main() {
	const goCount = 5
	wg := sync.WaitGroup{}
	wg.Add(goCount) // <===
	for i := 0; i < goCount; i++ {
		go func() {
			defer wg.Done() // <===
			fmt.Println("go-go-go")
		}()
	}
	wg.Wait() // <===
}

// WaitGroup
//Методы sync.WaitGroup
//
//	type WaitGroup struct {
//		}
//	func (wg *WaitGroup) Add(delta int) - увеличивает счетчик WaitGroup.
//	func (wg *WaitGroup) Done() - уменьшает счетчик на 1.
//	func (wg *WaitGroup) Wait() - блокируется, пока счетчик WaitGroup не обнулится.
//
