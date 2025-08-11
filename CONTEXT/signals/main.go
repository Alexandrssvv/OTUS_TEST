package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go asd(ctx)

	go gracefulShut(cancel)

	<-ctx.Done()
	time.Sleep(1 * time.Second)
}

func gracefulShut(cancel context.CancelFunc) {
	// Создаем канал для получения сигналов
	sigs := make(chan os.Signal, 1)

	// Уведомляем канал о поступлении указанных сигналов
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println(("Ожидаем сигнал (нажмите Ctrl + C)..."))
	sig := <-sigs
	cancel()

	fmt.Println("Получен сигнал:", sig)

	time.Sleep(1 * time.Second)

	fmt.Println("Выход из программы.")
}

func asd(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст отменился")
			return
		default:
			fmt.Println(133)
			time.Sleep(time.Second)
		}

		fmt.Println("Новая итерация")
	}
}
