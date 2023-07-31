package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем слай
	arr := []int{2, 4, 6, 8, 10}

	// Вычисляем сумму квадратов и выводим
	fmt.Println(sum(arr))
}

func sum(arr []int) int {
	// Создаем буферезированный канал, в который будем записывать результаты вычислений
	ch := make(chan int, len(arr))

	// Используем sync.WaitGroup, чтобы закрыть канал, когда все горутины закончили свои вычисления
	var wg sync.WaitGroup

	// Увеличиваем счетчик на количество элементов в слайсе = количеству горутин
	wg.Add(len(arr))

	// Запускаем горутины
	for i := 0; i < len(arr); i++ {
		go calc(&wg, arr[i], ch)
	}

	go func() {
		// Ждем пока все горутины не закончат вычисления
		wg.Wait()

		// Закрываем канал
		close(ch)
	}()

	var res int

	// Суммируем результаты вычислений
	for v := range ch {
		res += v
	}

	return res
}

// Функция вычисления квадрата числа
func calc(wg *sync.WaitGroup, n int, ch chan<- int) {
	// Уменьшаем счетчик на 1
	defer wg.Done()

	// Отправляем результат вычисления в канал
	ch <- n * n
}
