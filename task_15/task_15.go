package main

import (
	"crypto/rand"
	"encoding/hex"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	/*
		Проблема: функция createHugeString возвращает большую строку, от которой берется срез, который ссылается на исходный массив строки.
		К чему это может привести: утечка памяти, потому что сборщик мусора будет считать, что исходный массив строки используется.

		Решение: скопировать нужный нам срез.

		Примеры:
			Заменить:
				justString = v[:100]

			На:
				buf := make([]byte, 100)
				copy(buf, v[:100])
				justString = string(buf)

			Или:
				var builder strings.Builder
				builder.Write([]byte(v[:100]))
				justString = builder.String()
	*/

	buf := make([]byte, 100)
	copy(buf, v[:100])
	justString = string(buf)
}

func main() {
	someFunc()
}

func createHugeString(n int) string {
	buf := make([]byte, n)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}
