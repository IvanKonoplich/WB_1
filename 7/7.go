package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.

func main() {
	var counters sync.Map    //используем встроенную потокобезопасную мапу
	var wg sync.WaitGroup    //создвем счетчик для контроля горутин
	wg.Add(5)                //добавляем значения к waitgroup
	for i := 0; i < 5; i++ { //создаем 5 горутин, которые принимают ссылочку ну мапу
		go func(m *sync.Map) {
			for i := 1; i < 10; i++ {
				m.Store(i, i) //пишем в мапу значения от 1 до 10
			}
			wg.Done()
		}(&counters)
	}
	wg.Wait()
	fmt.Println(counters.Load(8)) //вывводим какие нибудь значение в stdout
}
