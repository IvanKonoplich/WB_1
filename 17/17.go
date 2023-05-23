package main

import (
	"fmt"
)

//Реализовать бинарный поиск встроенными методами языка.

func main() {
	//сканируем отсортированный массив
	fmt.Println("введите колличество элементов массива")
	var inc int
	fmt.Scan(&inc)
	var digit int
	fmt.Println("введите числа массива")
	digits := make([]int, inc)
	for i := 0; i < inc; i++ {
		fmt.Scan(&digit)
		digits[i] = digit
	}
	//сканируем число которое хотим найти
	fmt.Println("введите индекс для поиска")
	fmt.Scan(&inc)
	//используем  функцию бинарного поиска
	result := binarySearch(digits, inc)
	if result == -1 {
		fmt.Println("такого элемента нет")
	} else {
		fmt.Printf("число %d было найдено под индексом %d", inc, result)
	}
}

func binarySearch(a []int, search int) (result int) {
	mid := len(a) / 2 //находим середину слайса
	switch {
	case len(a) == 0:
		result = -1 // если длина слайса 0 возвращаем -1
	case a[mid] > search:
		result = binarySearch(a[:mid], search) //если значение по середине больше необходимого индекса рекурсивно вызываем функцию для первой половины слайса
	case a[mid] < search:
		result = binarySearch(a[mid+1:], search) //если дзначение по середине меньше необходимого индекса рекурсивно вызываем функцию для второй половины слайса
		if result >= 0 {                         // если результат не равен -1 (не нашли)
			result += mid + 1 //то результат равен следующему после середины индексу, ведь рекурсивный вызов функции вернет в виде результата индекс переданного меньшего слайса
		}
	default: //значение по индексу равно искомому
		result = mid
	}
	return
}
