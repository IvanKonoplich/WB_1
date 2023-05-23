package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2500*time.Millisecond) //создаем контекст с таймаутом через некоторое время
	defer cancel()

	squareCh := make(chan uint32) //канал для приема ответов от горутины

	go generateSquares(ctx, 10, squareCh) //запускаем горутину

	for s := range squareCh { //читаем из канала
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func generateSquares(ctx context.Context, n uint32, out chan<- uint32) {
	defer close(out)

	for i := uint32(1); i <= n; i++ {
		select {
		case <-ctx.Done():
			return
		case out <- i * i:
		}
	}
}
