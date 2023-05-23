package main

import (
	"fmt"
)

//Дана последовательность температурных
//колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов.
//Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
func main() {
	//сканируем инпут
	fmt.Println("введите колличество чисел в массиве")
	var numberOfDigits int
	fmt.Scan(&numberOfDigits)
	digits := make([]float32, numberOfDigits)
	var digit float32
	fmt.Println("введите числа")
	for i := 0; i < numberOfDigits; i++ {
		fmt.Scan(&digit)
		digits[i] = digit
	}
	//создаем мапу для градации по температуре где ключом является градация с шагом 10 и значением сама температура из ввода
	result := make(map[int][]float32, numberOfDigits)
	//в цикле проверяем каждый элемент инпута
	for _, j := range digits {
		//определяем градацию
		grade := j / 10
		//хаписываем в мапу
		result[int(grade)*10] = append(result[int(grade)*10], j)
	}
	//выводим результат
	fmt.Println(result)
}
