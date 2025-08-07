package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 100)

	_ = len(ch)
	_ = cap(ch)

	go func() {
		for i := range 1000 {
			ch <- i
			ch <- i
			ch <- i
			time.Sleep(time.Second)
		}

		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

	some()
}

func some() chan int {
	ch := make(chan int)
	//
	return ch
}
