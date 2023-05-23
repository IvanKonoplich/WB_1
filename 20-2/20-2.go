package main

import (
	"fmt"
	"strings"
)

func main() {
	var input = "snow € dog € sun €"   //переменная со строкой
	split := strings.Split(input, " ") //разделяем ее по пробелу
	input = ""                         //делаем уже созданную переменную пустой строкой
	for i := len(split) - 1; i >= 0; i-- {
		input += split[i] //добавляем к ней слова в обратном порядке
		input += " "      //разделяем пробелом
	}
	fmt.Println(input)
}
