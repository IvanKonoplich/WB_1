package main

import "fmt"

//Разработать программу, которая переворачивает
//подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	//сканируем
	fmt.Println("введите строку")
	var input string
	fmt.Scan(&input)
	//читаем байты строки и выводим через функцию отложенного вызова
	for _, j := range []rune(input) {
		defer fmt.Print(string(j))
	}
}
