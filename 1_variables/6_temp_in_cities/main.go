package main

import "fmt"


// Программа "Прогноз погоды"
// 1) Cохранить два российских города и температуру в них
// 2) Вывести
//
// Вывод должен быть таким:
//
// В городе Сочи сейчас 23 °C
// В городе Новосибирск сейчас 15 °C

func main() {
	var city string = "Москва"
	var city2 string = "Пенза"
	var tempMoscow int = 23
	var tempPenza = 16

	fmt.Printf("В городе %s сейчас %d °C\n", city, tempMoscow)
	fmt.Printf("В городе %s сейчас %d °C\n", city2, tempPenza)
	
}