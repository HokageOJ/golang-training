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
			// to do
		} else if toUnit == "K" {
			// to do
		} else {
			fmt.Printf("Вы передали некорректную единицу измерения в аргумент 'toUnit': %s\n", toUnit)
			return 0.0
		}
	} 
	
	return 0.0
}

func main() {
	
}