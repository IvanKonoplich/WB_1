package main

import (
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	stop := make(chan bool)       // канал для остановки
	squareCh := make(chan uint32) //канал для ответа

	go generateSquares(10, stop, squareCh) // запускаем горутину генерирующую квадраты чисел

	go func() { // запускаем еще одну горутину которая спит некоторое время и после отправляет в stop канал значение
		time.Sleep(2500 * time.Millisecond)
		stop <- true
	}()

	for s := range squareCh { //читаем из канала для ответа
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func generateSquares(n uint32, stop <-chan bool, out chan<- uint32) { //функция для рассчета квадратов чисел которую запустим в горутине
	defer close(out)

	for i := uint32(1); i <= n; i++ {
		select {
		case out <- i * i:
		case <-stop:
			return
		}
	}
}
