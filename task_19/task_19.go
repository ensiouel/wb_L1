package main

import "fmt"

func main() {
	fmt.Println(reverse("главрыба"))
}

func reverse(s string) string {
	res := []rune(s)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}
