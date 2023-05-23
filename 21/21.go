package main

//Реализовать паттерн «адаптер» на любом примере.

// Target определяет как должна работать программа
type Target interface {
	Request() string
}

// Adaptee система которую адаптируем
type Adaptee struct {
}

// NewAdapter конструктор
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

// SpecificRequest реализации того, что адаптируем
func (a *Adaptee) SpecificRequest() string {
	return "Request"
}

// Adapter бует реализовывать интерфейс Target
type Adapter struct {
	*Adaptee
}

// Request адаптированный метод
func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
