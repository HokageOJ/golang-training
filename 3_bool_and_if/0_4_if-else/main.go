package main

import "fmt"

/*
	-----------------------------------
	| Тема занятия "Оператор if-else" |
	-----------------------------------
*/

func main() {
	// Скорость ветра (м/с)
	var windSpeed int = 36
	// Категория
    var category string

    if windSpeed < 5 {
        category = "штиль"
    } else if windSpeed <= 20 {
        category = "легкий ветер"
    } else if windSpeed <= 35 {
        category = "сильный ветер"
    } else {
        category = "ураган"
    }

    fmt.Println("Категория ветра:", category)
}