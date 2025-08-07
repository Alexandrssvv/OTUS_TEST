package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ch = make(chan int)

	go func() {
		val := rand.Int()

		ch <- val

		time.Sleep(time.Second)

		ch <- val

		time.Sleep(time.Second)

		close(ch)
	}()

	//for val := range ch {
	//
	//	fmt.Println(val)
	//}

	_, ok := <-ch
	fmt.Println(ok)

	val, ok := <-ch
	fmt.Println(val)

	val, ok = <-ch
	fmt.Println(ok)
	fmt.Println(val)

}
