package main

import "fmt"

func main() {
	th := thermometer{}
	tha := NewThermometerAdapter(&th)
	fmt.Println(tha.GetWeather())
}

// допустим есть сервис по получение погоды
type weather interface {
	GetWeather() string
}

// есть термометр
type thermometer struct {
}

// NewThermometerAdapter конструктор
func NewThermometerAdapter(thermometer *thermometer) *thermometerAdapter {
	return &thermometerAdapter{thermometer}
}

// giveMeTemperature запрос на получение температуры который возвращает float64
func (a *thermometer) giveMeTemperature() float64 {
	return 36.6
}

// thermometerAdapter бует реализовывать интерфейс weather
type thermometerAdapter struct {
	*thermometer
}

// GetWeather адаптированный метод
func (a *thermometerAdapter) GetWeather() string {
	return fmt.Sprint(a.giveMeTemperature())
}
