package main

import "fmt"

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	set2 := []int{6, 1, 2, 3, 4, 9, 10, 5, 7, 8}

	fmt.Println(intersection(set1, set2))
}

func intersection(set1 []int, set2 []int) []int {
	// Создаем мапу для хранения наличия элементов
	m := make(map[int]struct{})
	for i := 0; i < len(set1); i++ {
		m[set1[i]] = struct{}{}
	}

	var res []int

	// Проверяем пересечение элементов
	for i := 0; i < len(set2); i++ {
		if _, ok := m[set2[i]]; ok {
			res = append(res, set2[i])
		}
	}

	return res
}
