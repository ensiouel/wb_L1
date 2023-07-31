package main

import "fmt"

func main() {
	a, b := 4, 6

	{
		a, b = b, a
		fmt.Println(a, b)
	}

	{
		a = a + b
		b = a - b
		a = a - b
		fmt.Println(a, b)
	}

	{
		a = a ^ b
		b = a ^ b
		a = a ^ b
		fmt.Println(a, b)
	}
}
