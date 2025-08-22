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
	
}