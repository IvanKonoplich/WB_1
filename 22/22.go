package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая
//перемножает,
//делит,
//складывает,
//вычитает
//две числовых переменных a,b, значение которых > 2^20.

func main() {
	//сканируем как строку
	fmt.Println("введите два числа ")
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	multiMath(a, b)

}

func multiMath(f, s string) {
	//парсим строки в числа
	a, _, _ := big.ParseFloat(f, 10, 0, big.ToNearestEven)
	b, _, _ := big.ParseFloat(s, 10, 0, big.ToNearestEven)
	a.Mul(a, b) // умножение
	fmt.Println("умножение: ", a)
	a.Quo(a, b) // возвращаем как было
	a.Quo(a, b) // деление
	fmt.Println("деление: ", a)
	a.Mul(a, b) // возвращаем как было
	a.Add(a, b) // сложение
	fmt.Println("сложение: ", a)
	a.Sub(a, b) // возвращаем как было
	a.Sub(a, b) // вычитание
	fmt.Println("вычитание: ", a)
}
