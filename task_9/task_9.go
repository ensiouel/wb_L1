package main

import "fmt"

func main() {
	// Создаем массив
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Берем срез массива и генерируем конвеер чисел
	gen := generate(arr[:])

	// Читаем из конвеера, делаем операции и пишем в новый канал
	res := double(gen)

	// Выводим результат
	for n := range res {
		fmt.Println(n)
	}
}

func generate(arr []int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for i := 0; i < len(arr); i++ {
			res <- arr[i]
		}
	}()

	return res
}

func double(ch <-chan int) <-chan int {
	res := make(chan int)

	go func() {
		defer close(res)

		for n := range ch {
			res <- n * 2
		}
	}()

	return res
}
