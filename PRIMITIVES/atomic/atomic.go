package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var dd atomic.Int32
//dd.Add(123)
//fmt.Println(dd.Load())

func main() {
	var counter atomic.Int32
	//var counter1 int64
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
			//atomic.AddInt64(&counter1, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.Load())
	//fmt.Println("Counter1:", counter1)
}
