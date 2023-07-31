package main

import "fmt"

func main() {
	fmt.Println(set([]string{"cat", "cat", "dog", "cat", "tree"}))
}

func set[T comparable](arr []T) []T {
	// Создаем карту для хранения наличия элементов
	m := make(map[T]struct{})

	var res []T
	for i := 0; i < len(arr); i++ {
		v := arr[i]

		// Если элемент отсутствует - добавляем и помечаем флаг
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = struct{}{}
		}
	}

	return res
}
