package main

import "fmt"

func main() {
	//сканируем
	fmt.Println("введите колличество элементов в слайсе")
	var numberOfDigits int
	fmt.Scan(&numberOfDigits)
	digits := make([]int, numberOfDigits)
	var digit int
	fmt.Println("введите элементы слайса")
	for i := 0; i < numberOfDigits; i++ {
		fmt.Scan(&digit)
		digits[i] = digit
	}
	var i int
	fmt.Println("введите индекс элемента слайса который хотите удалить")
	fmt.Scan(&i)
	fmt.Println(remove(digits, i))
}

func remove(slice []int, s int) []int {
	//соединяем два слайса созданных из изначального
	//первый хранит элементы до нужного индекса
	//второй от следующего после индекса элемента и до конца
	return append(slice[:s], slice[s+1:]...)
}
