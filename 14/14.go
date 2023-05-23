package main

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func main() {
	var a, b, c, d interface{}
	a = 1
	b = "1"
	c = true
	d = make(chan int)

	at := reflect.TypeOf(a).Kind()
	bt := reflect.TypeOf(b).Kind()
	ct := reflect.TypeOf(c).Kind()
	dt := reflect.TypeOf(d).Kind()

	fmt.Printf("%s: %s\n", fmt.Sprint(a), at)
	fmt.Printf("%s: %s\n", fmt.Sprint(b), bt)
	fmt.Printf("%s: %s\n", fmt.Sprint(c), ct)
	fmt.Printf("%s: %s\n", fmt.Sprint(d), dt)

}
