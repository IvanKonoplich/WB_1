package main

import (
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

//если я правильно понял, то createHugeString возвращает строку нужного размера с любым содержимым.
//Непонятна необходимость использовать <<.
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

func createHugeString(inc int) string {
	b := strings.Builder{}
	b.Grow(inc)
	for i := 0; i < inc; i++ {
		b.WriteString("1")
	}
	return b.String()
}
