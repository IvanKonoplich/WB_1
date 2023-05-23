package main

import (
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	//сканируем
	fmt.Println("введите время работы программы")
	var seconds int64
	fmt.Scan(&seconds)
	//используем time.After для создания канала на который спустя определенное время придут данные
	timeout := time.After(time.Duration(seconds) * time.Second)
	//создаем канал для последовательной отправки данных
	c := make(chan int)
	//вызываем горутину для отправки данных
	go func() {
		for i := 0; ; i++ {
			//отправляем последовательно числа и ждем секунду
			c <- i
			time.Sleep(time.Second)
		}
	}()
	//в бесконечном цикле for смотрим два канала и завершаем работу когда канал time.After сработает
	for {
		select {
		case i := <-c:
			fmt.Println(i)
		case <-timeout:
			fmt.Println("end")
			return
		}
	}
}
