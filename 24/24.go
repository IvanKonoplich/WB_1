package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

//определяем структуру Point
type Point struct {
	x float64
	y float64
}

//определяем конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	//сканируем
	fmt.Println("введите координаты первой точки")
	var x, y float64
	fmt.Scan(&x)
	fmt.Scan(&y)
	first := NewPoint(x, y)
	fmt.Println("введите координаты второй точки")
	fmt.Scan(&x)
	fmt.Scan(&y)
	second := NewPoint(x, y)
	fmt.Println("расстояние между точками равно - " + fmt.Sprint(findDistance(first, second)))
}

func findDistance(a, b *Point) float64 {
	result := math.Sqrt(math.Pow(a.x-b.x, 2)) + math.Pow(a.y-b.y, 2) //используем формулу AB = √ (x b - x a) 2 + (y b - y a)
	return result
}
