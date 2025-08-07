package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	for i := range 100 {
		go func(a int) {
			fmt.Println(a)

			someFunc(a)
		}(i)
	}

	//m := runtime.MemStats{}
	//
	//runtime.ReadMemStats(&m)
	//
	//fmt.Println(runtime.NumGoroutine())

	time.Sleep(time.Second)
}

func someFunc(a int) {
	a++
}
