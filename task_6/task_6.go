package main

import (
	"context"
	"time"
)

func main() {
	// С помощью контекста
	{
		ctx, cancel := context.WithCancel(context.Background())
		/*
			Еще варианты:
			- context.WithCancelCause(context.Background())
			- context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
			- context.WithTimeout(context.Background(), 5*time.Second)
		*/

		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					// some work
					time.Sleep(1 * time.Second)
				}
			}
		}(ctx)

		time.Sleep(5 * time.Second)
		cancel()
	}

	// С помощью канала
	{
		done := make(chan struct{})

		go func(done <-chan struct{}) {
			for {
				select {
				case <-done:
					return
				default:
					// some work
					time.Sleep(1 * time.Second)
				}
			}
		}(done)

		time.Sleep(5 * time.Second)
		done <- struct{}{}
	}
}
