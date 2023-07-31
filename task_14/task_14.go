package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a interface{}

	a = 3
	fmt.Println(typeOf(a))

	a = "3"
	fmt.Println(typeOf(a))

	a = true
	fmt.Println(typeOf(a))

	a = make(chan int)
	fmt.Println(typeOf(a))
}

func typeOf(v interface{}) string {
	// Используем пакет reflect для получения типа
	// так же можно было заменить на v.(type), но возникают проблемы с каналами
	return reflect.TypeOf(v).Kind().String()
}
