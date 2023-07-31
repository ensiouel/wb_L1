package main

import (
	"fmt"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	m := make(map[int][]float64)

	for i := 0; i < len(arr); i++ {
		// Округляем число
		r := round(arr[i])

		// Добавляем в срез
		m[r] = append(m[r], arr[i])
	}

	fmt.Println(m)
}

func round(n float64) int {
	return int(n) / 10 * 10
}
