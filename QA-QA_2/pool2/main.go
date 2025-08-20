package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task func(ctx context.Context, args User) error

func main() {
	pool := sync.Pool{
		New: func() any {
			return make(chan Task, 1)
		},
	}

	ch := pool.Get().(chan Task)

	go func(taskCh <-chan Task) {
		for task := range taskCh {
			err := task(context.Background(), User{FirstName: "John"})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}(ch)

	go func(taskCh <-chan Task) {
		for task := range taskCh {
			err := task(context.Background(), User{FirstName: "Doe"})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}(ch)

	go func(taskCh <-chan Task) {
		for task := range taskCh {
			err := task(context.Background(), User{FirstName: "FirsName"})
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}(ch)

	for range 100 {
		ch <- func(ctx context.Context, args User) error {
			fmt.Println(args.FirstName)
			return nil
		}

		time.Sleep(1 * time.Second)
	}

	var task Task = func(ctx context.Context, args User) error {
		fmt.Println(args)

		return nil
	}

	task(context.Background(), User{})
}

type User struct {
	FirstName string
}

func userF() {
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Hello World")
			return &User{FirstName: "World"}
		},
	}

	user := pool.Get().(*User)

	user.FirstName = "First Name"

	pool.Put(user)

	user2 := pool.Get().(*User)

	fmt.Println(user)
	fmt.Println(user2)
}
