package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep.

func main() {
	//сканируем
	fmt.Println("введите время сна программы")
	var seconds int64
	fmt.Scan(&seconds)
	sleep(seconds)
	fmt.Println("работаем дальше")
}

func sleep(seconds int64) {
	timeout := time.After(time.Duration(seconds) * time.Second) //используем time.After для создания канала на который спустя определенное время придут данные
	for {                                                       //в бесконечном цикле for смотрим канал и завершаем работу когда канал time.After сработает
		<-timeout
		fmt.Println("sleep end")
		return
	}
}
