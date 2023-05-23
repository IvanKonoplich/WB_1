package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	//сканируем инпут
	fmt.Println("введите колличество чисел в первом массиве")
	var numberOfDigits int
	fmt.Scan(&numberOfDigits)
	digits1 := make([]int, numberOfDigits)
	var digit int
	fmt.Println("введите числа")
	for i := 0; i < numberOfDigits; i++ {
		fmt.Scan(&digit)
		digits1[i] = digit
	}

	fmt.Println("введите колличество чисел во втором массиве")
	fmt.Scan(&numberOfDigits)
	digits2 := make([]int, numberOfDigits)
	fmt.Println("введите числа")
	for i := 0; i < numberOfDigits; i++ {
		fmt.Scan(&digit)
		digits2[i] = digit
	}
	fmt.Print(intersection(digits1, digits2))
}

func intersection(one, two []int) []int {
	//создаем мапу для проверки, сразу определяем ее cap
	var m = make(map[int]uint, len(one))
	//заполняем мапу элементами первого массива, добавляем сам элемент как ключ и колличество его появлений как значение
	for i := range one {
		m[one[i]]++
	}
	//создаем слайс для результата, сразу определяя его cap
	var result = make([]int, 0, len(one))

	//пробегаемся по второму слайсу
	for i := range two {
		//если значение присутствует в мапе, то уменьшаем колличество появления этого числа и добавляем его в результат
		if value, ok := m[two[i]]; ok {
			if value > 0 {
				m[two[i]]--
				result = append(result, two[i])
			} else {
				delete(m, two[i])
			}
		}
	}
	return result
}
