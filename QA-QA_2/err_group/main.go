package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	data := make([]int, 10)
	for i := range data {
		data[i] = i
	}

	err := processData(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func processData(data []int) error {
	var (
		wg          = sync.WaitGroup{}
		errChan     = make(chan error, len(data))
		ctx, cancel = context.WithCancel(context.Background())
	)

	defer cancel()

	for _, v := range data {
		wg.Add(1)

		go func() {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
			}

			if rand.Intn(10) == 0 {
				cancel()
				errChan <- errors.New("processData error")
			} else {
				fmt.Println(v)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}
