package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	//сканируем
	fmt.Println("введите два числа")
	var first int
	var second int
	fmt.Scan(&first)
	fmt.Scan(&second)
	//меняем местами
	first, second = second, first
	fmt.Println(first, second)
}
