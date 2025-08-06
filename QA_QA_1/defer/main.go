package main

import (
	"fmt"
)

func main() {
	a := 10

	defer func() {
		fmt.Println("call 0", a+10)
	}()

	defer abc(&a)

	a++

	fmt.Println("call 2", a)
}

func abc(a *int) {
	*a++
}
