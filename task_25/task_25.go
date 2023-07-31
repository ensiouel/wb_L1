package main

import (
	"fmt"
	"time"
)

func main() {
	sleep(5 * time.Second)
	fmt.Println("done")
}

func sleep(d time.Duration) {
	// time.After создает канал, в который отправит сигнал через time.Duration
	<-time.After(d)
}
