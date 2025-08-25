package main

import (
	"fmt"
)

/*
	---------------
	| Температуры |
	---------------
	Создайте именованные типы для градусов Цельсия, Фаренгейтов и Кельвинов.

	1) Напишите методы для красивого вывода типов.
	2) Напишите методы для перевода типов в другие единицы измерения.
*/

type Celsius float64
type Fahrenheit float64
type Kelvin float64

func (temp Celsius) String() string {
	return fmt.Sprintf("%.1f°С", temp)
} 

func (temp Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit(temp * 9/5 + 32)
}

func (temp Celsius) ToKelvin() Kelvin {
	return Kelvin(temp + 273.15)
}

func (temp Fahrenheit) String() string {
	return fmt.Sprintf("%.1f°F", temp)
}

func (temp Fahrenheit) ToCelsius() Celsius {
	return Celsius((temp - 32) * 5/9)
}

func (temp Fahrenheit) ToKelvin() Kelvin {
	return Kelvin((temp - 32) * 5/9 + 273.15)
}

func (temp Kelvin) String() string {
	return fmt.Sprintf("%.1fK", temp)
}

func (temp Kelvin) ToCelsius() Celsius {
	return Celsius(temp - 273.15)
}

func (temp Kelvin) ToFahrenheit() Fahrenheit {
	return Fahrenheit(temp * 9/5 - 459.67)
}

func main() {
	var celsiusTemp Celsius = 20.0
	
	fmt.Println(celsiusTemp.String())
	fmt.Println(celsiusTemp.ToFahrenheit())
	fmt.Println(celsiusTemp.ToKelvin())
	fmt.Println()

	var fahrenheitTemp Fahrenheit = 10.0

	fmt.Println(fahrenheitTemp.String())
	fmt.Println(fahrenheitTemp.ToCelsius())
	fmt.Println(fahrenheitTemp.ToKelvin())
	fmt.Println()
	
	var kelvinTemp Kelvin = 5.0

	fmt.Println(kelvinTemp.String())
	fmt.Println(kelvinTemp.ToCelsius())
	fmt.Println(kelvinTemp.ToFahrenheit())
}