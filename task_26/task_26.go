package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}

func unique(s string) bool {
	// Преобразуем строку в нижний регистр
	s = strings.ToLower(s)

	m := make(map[rune]struct{})
	for _, r := range s {
		// Если символ уже есть - не все символы уникальные
		if _, ok := m[r]; ok {
			return false
		}

		m[r] = struct{}{}
	}

	return true
}
