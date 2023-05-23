package main

import (
	"fmt"
	"log"
	"sync"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

//структура-счетчик
type Counter struct {
	num int
	sync.Mutex
}

//инкрементация
func (c *Counter) Inc() {
	//обеспечиваем безопасный доступ к счетчику
	c.Lock()
	defer c.Unlock()
	c.num += 1
}

func main() {
	//создаем структуру-счетчик
	cnt := &Counter{}

	//канал для завершения
	finish := make(chan struct{})
	//получаем инпут
	fmt.Println("введите значение счетчика")
	var inc int
	fmt.Scan(&inc)
	//вызываем горутины
	go Do(cnt, finish, inc)
	//когда приходит значение из канала main разблокируется и завершит работу
	<-finish
	log.Printf("программа завершилась со значением счетчика - %d", cnt.num)
}

func Do(cnt *Counter, finish chan struct{}, inc int) {
	//wg для контроля горутин
	wg := sync.WaitGroup{}

	//создаем необходимое колличество горутин
	for i := 0; i < inc; i++ {
		wg.Add(1)

		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("горутина %d стартовала\n", num)
			cnt.Inc()
			fmt.Printf("горутина %d закончила\n", num)
		}(i, cnt, &wg)
	}

	wg.Wait()
	//когда все воркеры отработали посылаем значение в канал
	finish <- struct{}{}
}
