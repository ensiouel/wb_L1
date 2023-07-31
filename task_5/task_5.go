package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var workDuration = flag.Duration("work_duration", 5*time.Second, "work duration")

func main() {
	// Создаем контекст, который завершится через N секунд
	ctx, cancel := context.WithTimeout(context.Background(), *workDuration)
	defer cancel()

	// Создаем канал для работы с данными
	ch := make(chan int)

	// Запускаем генератор данных
	go producer(ctx, ch)

	// Запускаем обработчик
	go consumer(ch)

	// Ожидаем
	<-ctx.Done()
}

func consumer(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
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
