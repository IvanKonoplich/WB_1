package main

import (
	"fmt"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	//подсказочка для ввода
	fmt.Println("введите число горутин")
	//объявляем переменную
	var num int
	//сканируем ввод
	fmt.Scan(&num)
	//создаем канал
	c := make(chan any)
	//создаем необходимое колличество горутин
	for i := 0; i < num; i++ {
		//принимаем канал и номер горутины
		go func(c chan any, i int) {
			//бесконечный цикл чтобы горутина не завершилась
			for {
				//принимаем и выводим в stdout данныебесконечное чтение из канала
				inc := fmt.Sprint(<-c)
				fmt.Println(fmt.Sprintf("канал номер %d начал работу с данными: %s", i+1, inc))
				time.Sleep(time.Second * 2)
				fmt.Println(fmt.Sprintf("канал номер %d завершил работу с данными: %s", i+1, inc))
			}
		}(c, i)
	}
	//бесконечно принимаем данные из терминала и передаем в канал
	for {
		fmt.Println("введите данные")
		var data string
		fmt.Scan(&data)
		c <- data
	}
}