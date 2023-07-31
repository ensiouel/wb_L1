package main

import (
	"fmt"
)

func main() {
	{
		arr := []int{1, 2, 3, 4, 5, 6}
		arr = removeStable(arr, 3)
		fmt.Println(arr)
	}

	{
		arr := []int{1, 2, 3, 4, 5, 6}
		arr = remove(arr, 3)
		fmt.Println(arr)
	}
}

// Копируем слайс без i-ого элемента
func removeStable(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

// Можно заменить элемент с индексом i на последний и усеч срез, может привести к утечке памяти
func remove(arr []int, i int) []int {
	arr[i] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
