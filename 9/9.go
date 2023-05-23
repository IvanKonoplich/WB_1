package main

import "fmt"

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из
//второго канала должны выводиться в stdout.

func main() {
	//создаем канал в который пишутся числа
	inputChannel := make(chan int)
	//канал в который идет рузультат
	outputChannel := make(chan int)

	//переменная для колличества чисел в вводе
	var numberOfDigits int
	//подсказка
	fmt.Println("введите колличество чисел в массиве")
	//сканируем
	fmt.Scan(&numberOfDigits)
	digits := make([]int, numberOfDigits)
	var digit int
	fmt.Println("введите числа")
	for i := 0; i < numberOfDigits; i++ {
		fmt.Scan(&digit)
		digits[i] = digit
	}
	//создаем горутину которая считает квадраты чисел, передаем оба созданных канала
	go func(ic, oc chan int) {
		//берем числа из первого канала
		for {
			d := <-inputChannel
			//отправляем результат во второй
			fmt.Printf("горутина получила число - %d\n", d)
			outputChannel <- d * d
		}
	}(inputChannel, outputChannel)
	//передаем числа из инпута в первый канал и читаем результат из второго
	for _, j := range digits {
		inputChannel <- j
		fmt.Println(<-outputChannel)
	}
}
