package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//	aabcd — false

func main() {
	//сканируем
	fmt.Println("введите строку")
	var input string
	fmt.Scan(&input)
	fmt.Println(findDups(input))
}

func findDups(s string) bool {
	checker := make(map[byte]bool, len(s)) //создаем мапу для проверки
	s = strings.ToLower(s)                 //приводим все к нижнему регистру
	for _, j := range []byte(s) {          //проверяем каждый байт в строке на уникальность
		if _, ok := checker[j]; ok {
			return false //если значение есть в мапе то возвращаем false
		} else {
			checker[j] = true //если значения нет в мапе то добавляем его
		}
	}
	return true //если дубликатов не было то возвращаем true
}
