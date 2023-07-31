package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var workerCount = flag.Int("worker_count", 10, "count of worker")

func main() {
	flag.Parse()

	// Создаем контекст который будет завершаться сигналом Ctrl+C
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Создаем канал для работы с данными
	ch := make(chan int, *workerCount)

	// Запускаем генератор данных
	go producer(ctx, ch)

	// Запускаем обработчики
	for i := 0; i < *workerCount; i++ {
		go worker(i, ch)
	}

	// Ожидаем сигнала Ctrl+C
	<-ctx.Done()
}

func worker(id int, ch <-chan int) {
	// При чтении из канала с помощью range мы выйдем из цикла когда канал будет закрыт
	// тем самым мы заверишим работу воркера
	for v := range ch {
		fmt.Printf("worker %d: %d\n", id, v)
	}
}

func producer(ctx context.Context, ch chan<- int) {
	// В цикле отправляем данные в канал
	for {
		select {
		// Проверяем закрытие контекста
		case <-ctx.Done():
			close(ch)
			return
		case ch <- rand.Int():
			time.Sleep(100 * time.Millisecond)
		}
	}
}
