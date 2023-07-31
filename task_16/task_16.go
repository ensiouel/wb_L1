package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{2, 34, 5, 1, 7, 9, 3, -6, -10}
	fmt.Println(arr)

	Sort(arr)
	fmt.Println(arr)
}

func Sort[T Ordered](arr []T) {
	// Если элементов меньше 2 (0 или 1) - слайс уже отсортирован
	if len(arr) < 2 {
		return
	}

	// Разбиение Ломуто*

	left, right := 0, len(arr)-1

	// Выбираем случайный опорный элемент
	pivot := rand.Int() % len(arr)

	// Перемещаем опорный элемент в конец
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Проходимся по слайсу
	// Если i-ый элемент меньше опорного, то мы помещаем его на позицию left и увеличиваем значение left
	for i := 0; i < len(arr); i++ {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	// Перемещаем опорный элемент после наименьшего значения
	arr[left], arr[right] = arr[right], arr[left]

	// Проделываем те же операции для левой и правой части от pivot
	Sort(arr[:left])
	Sort(arr[left+1:])
}

type Ordered interface {
	Integer | Float | ~string
}

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface {
	~float32 | ~float64
}
