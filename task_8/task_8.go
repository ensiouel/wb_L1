package main

import (
	"fmt"
)

func main() {
	// 100 | (1 << 1) == 100 | 010 == 110
	fmt.Printf("%b\n", setBit(4, 1))

	// 110 & ^(1 << 1) == 110 & 101 == 100
	fmt.Printf("%b\n", unsetBit(6, 1))
}

func setBit(n int64, pos uint) int64 {
	return n | (1 << pos)
}

func unsetBit(n int64, pos uint) int64 {
	return n & ^(1 << pos)
}
