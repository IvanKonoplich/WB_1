package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает
//подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	fmt.Println("введите строку")
	var input string
	fmt.Scan(&input)
	b := strings.Builder{}
	//читаем байты строки в обратном порядке и приращиваем к strings.Builder
	for i := len([]rune(input)) - 1; i >= 0; i-- {
		b.WriteRune([]rune(input)[i])
	}
	fmt.Print(b.String())
}
