package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})   //канал для остановки
	squareCh := make(chan uint32) //канал для приема ответов

	go generateSquares(10, done, squareCh) //вызываем нашу горутину

	go func() { //спим некоторое время и закрываем канал done
		time.Sleep(2500 * time.Millisecond)
		close(done)
	}()

	for s := range squareCh { //читаем из канала с ответами
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func generateSquares(n uint32, done <-chan struct{}, out chan<- uint32) {
	defer close(out)

	for i := uint32(1); i <= n; i++ {
		select {
		case <-done:
			return
		case out <- i * i:
		}
	}
}
