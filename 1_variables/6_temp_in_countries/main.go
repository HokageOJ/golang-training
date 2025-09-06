package main

import "fmt"


// Программа "Температура в странах"
// 1) Cохранить две страны и температуру в них
// 2) Вывести
//
// Вывод должен быть таким:
//
// В стране Норвегия сейчас 10 °C
// В стране Испания сейчас 28 °C
func main() {
	var countries string = "Норвегия"
	var countries2 string = "Испания"
	var tempNorway int = 10
	var tempSpain int = 28

	fmt.Printf("В стране %s сейчас %d °C\n", countries, tempNorway)
	fmt.Printf("В старне %s сейчас %d °C\n", countries2, tempSpain)
}