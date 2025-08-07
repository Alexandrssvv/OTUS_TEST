package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{}, 2)
	var isStop = 0

	go func() {
		for range 100 {
			ch1 <- rand.Int()
			ch1 <- rand.Int()
			time.Sleep(time.Second)
			close(ch1)
			ch3 <- struct{}{}
			break
		}
	}()

	go func() {
		for range 100 {
			ch2 <- rand.Int()
			ch2 <- rand.Int()
			time.Sleep(time.Second)
			close(ch2)
			ch3 <- struct{}{}
			break
		}
	}()

	ticker := time.NewTicker(time.Second)

MAIN_LOOP:
	for {
		select {
		case <-ticker.C:
			fmt.Println("Тикер")
		case val, ok := <-ch1:
			if ok {
				fmt.Printf("chan1: %d\n", val)
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Printf("chan2: %d\n", val)
			}
		case <-ch3:
			isStop++
			if isStop == cap(ch3) {
				break MAIN_LOOP
			}
		default:
			fmt.Println("nothing")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
