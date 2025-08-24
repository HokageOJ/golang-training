package main

import "fmt"

type Celsius float64
type Fahrenheit float64

// Метод для красивого вывода градусов Цельсия
func (temp Celsius) String() string {
	return fmt.Sprintf("%.1f°C", temp)
}

// Метод для перевода градусов Цельсия в Фаренгейты
func (temp Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(temp * 9/5 + 32)
}

func main() {
	var c Celsius = 20.0
	fmt.Println(c.String())
	fmt.Println(c.ToFahrenheit())
}

