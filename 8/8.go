package main

import "fmt"

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

// On устанавливает i-й бит в 1, используя побитовое ИЛИ
// с маской типа 00100 с единицей в позиции i
func On(n int64, i int) int64 {
	return n | 1<<(i-1)
}

// Off устанавливает i-й бит в 0, используя побитовое И
// с маской типа 11011 с нулем в позиции i
func Off(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func main() {
	fmt.Println("введите число")
	var inc int64
	fmt.Scan(&inc)
	fmt.Println("введите бит")
	var pos uint
	fmt.Scan(&pos)
	fmt.Println("введите 1 или 0")
	var val int
	fmt.Scan(&val)

	if val == 1 {
		fmt.Println(setBit(inc, pos))
	} else if val == 0 {
		fmt.Println(clearBit(inc, pos))
	} else {
		fmt.Println("некорректное значение")
	}
}

//setBit устанавливает i-й бит в 1
func setBit(n int64, pos uint) int64 {
	n |= 1 << pos
	return n
}

//clearBit устанавливает i-й бит в 0
func clearBit(n int64, pos uint) int64 {
	n &^= (1 << pos)
	return n
}
