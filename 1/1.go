package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

//определяем human
type Human struct {
}

//определяем action
type Action struct {
	Human
}

//метод для human
func (h *Human) one() {
	fmt.Print("one")
}

func main() {
	//создаем action
	a := Action{}
	//вызываем встроенный метод
	a.one()
}
