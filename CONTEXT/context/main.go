package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cancel := context.WithCancel(context.Background())
	cancel()
	//ctx2 := context.WithValue(ctx, "key", "value")
	ctx2 := context.WithValue(ctx1, "key", "value")

	go doSmt(ctx2)
	//cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("Отменили контекст")

	//go doSmt(context.WithoutCancel(ctx2))
	go doSmt(ctx2)
	time.Sleep(1 * time.Minute)
}

func doSmt(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)

	if ctx.Err() != nil {
		fmt.Println("А контекст уже отменен")
	}
	for {
		select {
		case <-ticker.C:
			fmt.Println(123)
		case <-ctx.Done():
			fmt.Println(ctx.Err().Error())
			fmt.Println("Горутина получила отмену контекста")
			return
		}
	}

	//fmt.Println("Горутина ждет отмену контекста")
	//<-ctx.Done()
	//fmt.Println("Горутина получила отмену контекста")

	//fmt.Println(123)
	//time.Sleep(1 * time.Second)
	//fmt.Println(123)
	//v := ctx.Value("key")
	//fmt.Println(v)
}
