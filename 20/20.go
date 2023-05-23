package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	var input = "snow € dog € sun €"       //переменная со строкой
	split := strings.Split(input, " ")     //разделяем ее по пробелу
	for i := len(split) - 1; i >= 0; i-- { //выводим слова в обратном порядке добавляя пробел
		fmt.Print(split[i] + " ")
	}
}
