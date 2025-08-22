package main

import "fmt"

/*
	-------------------------
	| Конвертер температуры |
	-------------------------
	Создайте функцию, которая переводит температуру из одной единицы измерения в другую.

	Функция должна поддерживать перевод из/в градусы Цельсия, Фаренгейта, Кельвины.
*/

func converterTemp(temp float64, fromUnit string, toUnit string) float64 {
	if fromUnit == "C" {
		if toUnit == "F" {
			return (temp * 9/5) + 32
		} else if toUnit == "K" {
			return temp + 273.15
		} else {
			fmt.Printf("Вы передали некорректную единицу измерения в аргумент 'toUnit': %s\n", toUnit)
			return 0.0
		}
	} else if fromUnit == "F" {
		if toUnit == "C" {
			return (temp - 32) * 5/9
		} else if toUnit == "K" {
			return (temp - 32) * 5/9 + 273.15
		} else {
			fmt.Printf("Вы передали некорректную единицу измерения в аргумент 'toUnit': %s\n", toUnit)
			return 0.0
		}
	} else if fromUnit == "K" {
		if toUnit == "C" {
			return temp - 273.15
		} else if toUnit == "F" {
			return temp * 9/5 - 459.67
		} else {
			fmt.Printf("Вы передали некорректную единицу измерения в аргумент 'toUnit': %s\n", toUnit)
			return 0.0
		}
	}

	return 0.0
}

func main() {
	fmt.Printf("Перевод 't'- %.1f из 'C' в 'F': %.1f\n", 10.0, converterTemp(10.0, "C", "F" ))
	fmt.Printf("Перевод 't'- %.1f из 'C' в 'K': %.1f\n", 10.0, converterTemp(10.0, "C", "K" ))
	fmt.Println()
	fmt.Printf("Перевод 't'- %.1f из 'F' в 'C': %.1f\n", 10.0, converterTemp(10.0, "F", "C" ))
	fmt.Printf("Перевод 't'- %.1f из 'F' в 'K': %.1f\n", 10.0, converterTemp(10.0, "F", "K" ))
	fmt.Println()
	fmt.Printf("Перевод 't'- %.1f из 'K' в 'C': %.1f\n", 10.0, converterTemp(10.0, "K", "C" ))
	fmt.Printf("Перевод 't'- %.1f из 'K' в 'F': %.1f\n", 10.0, converterTemp(10.0, "K", "F" ))
}