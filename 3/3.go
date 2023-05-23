package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
//с использованием конкурентных вычислений.

//вызываем функцию
func main() {
	//для наглядного отображения отсутствия data race при использовании мьютекса
	for i := 0; i < 1000; i++ {
		if square([5]int{2, 4, 6, 8, 10}) != 220 {
			fmt.Println("data race")
			break
		}
		fmt.Println("ok")
	}
	fmt.Println(square([5]int{2, 4, 6, 8, 10}))
}

//определяем функцию
func square(m [5]int) int {
	//создаем waitgroup чтобы все горутины гарантированно отработали
	var wg sync.WaitGroup
	//создаем переменную для результата
	var result int
	//создаем mutex для контроля доступа к result
	var mu sync.Mutex
	//проходимся по полученному массиву
	for _, j := range m {
		//инкрементируем счетчик
		wg.Add(1)
		//создаем горутину для каждого элемента в массиве
		go func(j int, result *int) {
			//используем функцию отложенного вызова для уменьшения счетчика
			defer wg.Done()
			//вызываем lock для предотвращения data race
			mu.Lock()
			//отменяем lock через функцию отложенного вызова
			defer mu.Unlock()
			//считаем квадрат числа
			s := j * j
			//используем stdout вместо fmt.Print, ведь так поинтереснее
			*result += s
		}(j, &result) //передаем j сразу, чтобы все горутины не получили одно значение
	}
	//запускаем работу счетчика
	wg.Wait()
	//возвращаем результат
	return result
}