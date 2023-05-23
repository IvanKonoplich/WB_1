package main

import (
	"fmt"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

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
	fmt.Println(quicksort(digits))
}

func quicksort(a []int) []int {
	if len(a) < 2 { //если один элемент то его и возвращаем
		return a
	}
	pivot := a[0]               //отпределяем от чего отталкиваемся
	var less, greater []int     //создаем слайсы для больших и меньших элементов относительно опоры
	for _, num := range a[1:] { //распределяем значения
		if num < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(quicksort(less), pivot) //рекурсивно сортируем полученные слайсы
	result = append(result, quicksort(greater)...)
	return result
}
