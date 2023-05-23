package main

import (
	"fmt"
)

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	//смотрим сколько строк будет в вводе
	var numberOfWords int
	//подсказочка
	fmt.Println("введите колличество строк")
	//сканируем кол-во строк
	fmt.Scan(&numberOfWords)
	resultSlice := make([]string, 0, numberOfWords)
	//incoming string
	var inc string
	// сканируем каждую строку
	fmt.Println("введите строки")
	for i := 0; i < numberOfWords; i++ {
		fmt.Scan(&inc)
		resultSlice = append(resultSlice, inc)
	}
	fmt.Println(makeMn(resultSlice))
}

func makeMn(inc []string) [][]string {
	result := make([][]string, len(inc)) //заранее аллоцируем память под результат
	for i := 0; i < len(inc); i++ {
		result = append(result, make([]string, 0, len(inc)))
	}
	checker := make(map[string]int, len(inc)) //создаем мапу для проверки, где индексом является номер слайса в результате
	var index int
	for _, j := range inc { //идем по входным строкам
		if i, ok := checker[j]; ok { //если значение уже было то аппендим результат
			result[i] = append(result[i], j)
		} else { //если не было то добавляем в мапу и в результат
			checker[j] = index
			result[index] = append(result[index], j)
			index++ //сдвигаем индекс слайса
		}
	}
	return result[:index] //возвращаем созданные подмножества
}
