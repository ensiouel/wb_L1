package main

import (
	"fmt"
)

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(Find(data, 4))
}

func Find(arr []int, target int) int {
	i, j := 0, len(arr)-1
	for i <= j {
		// Получаем центральный элемент
		h := i + (j-i)>>1

		// Сравниваем центральный элемент с тем, который нужно найти и смещаем указатели
		if arr[h] > target {
			j = h - 1
		} else if arr[h] < target {
			i = h + 1
		} else {
			return h
		}
	}

	return -1
}
