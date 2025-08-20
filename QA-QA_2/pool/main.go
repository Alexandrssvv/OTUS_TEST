package main

import (
	"fmt"
	"sync"
)

type User struct {
	FirstName string
}

func main() {
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
