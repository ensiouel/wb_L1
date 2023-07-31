package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("snow dog sun"))
}

func reverse(s string) string {
	res := strings.Split(s, " ")

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return strings.Join(res, " ")
}
