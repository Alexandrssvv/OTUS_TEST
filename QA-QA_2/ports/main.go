package main

import (
	"fmt"
	"net"
	"sync"
	"time"
	//"time"
	//"hw12/internal/service"
)

func scan(wg *sync.WaitGroup, i int) {
	defer wg.Done()

	address := fmt.Sprintf("localhost:%d", i)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		// порт закрыт или отфильтрован
		return
	}
	conn.Close()
	fmt.Printf("%d open\n", i)
}

func worker(ports chan int, result chan int) {
	for p := range ports {
		address := fmt.Sprintf("localhost:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// порт закрыт или отфильтрован
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}

func worker2(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		address := fmt.Sprintf("localhost:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			wg.Done()
			continue
		}
		fmt.Printf("%d open\n", p)
		conn.Close()
		wg.Done()
	}
}

func scanPorts() {
	fmt.Println("Scan Ports *************************")
	var wg sync.WaitGroup

	for i := 1; i <= 65_000; i++ {
		wg.Add(1)
		go scan(&wg, i)
	}

	wg.Wait()
}

func scanBufferedChan() {
	fmt.Println("scanBufferedChan *************************")

	ports := make(chan int, 100)

	var wg sync.WaitGroup

	for i := 0; i < cap(ports); i++ {
		go worker2(ports, &wg)
	}

	for i := 1; i <= 65_000; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()

	close(ports)
}

func rangeOverChannels() {
	ch := make(chan int, 3) // Create a buffered channel

	// Sender goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i                            // Send values to the channel
			time.Sleep(100 * time.Millisecond) // Simulate some work
		}
		close(ch) // Close the channel when done sending
	}()

	// Receiver goroutine using for...range
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}

	fmt.Println("Channel closed and all values processed.")
}

func main() {
	scanPorts()
	scanBufferedChan()
	scanTwoChannels()

}

func scanTwoChannels() {

	fmt.Println("scanTwoChannels *************************")

	ports := make(chan int, 100)
	result := make(chan int)

	for i := 0; i < cap(ports); i++ {
		go worker(ports, result)
	}

	portCount := 65_000

	go func() {
		for i := 0; i < portCount; i++ {
			ports <- i
		}
	}()

	for i := 0; i < portCount; i++ {
		port := <-result
		if port != 0 {
			fmt.Printf("%d open\n", port)
		}
	}

	close(ports)
	close(result)
}
