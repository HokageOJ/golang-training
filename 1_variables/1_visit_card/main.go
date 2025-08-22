package main

import "fmt"


// Программа "Визитка"
// 1) Нужно сохранить имя, возраст и город проживания гражданина
// 2) Вывести в красивом виде
//
// Вывод должен быть таким:
//
// Меня зовут Иван
// Мне 25 лет
// Я живу в городе Москва
func main() {
	var name string = "Дмитрий"
	var age int = 24
	var city string = "Пенза"

	fmt.Println("Меня зовут", name)
	fmt.Println("Мне", age)
	fmt.Println("Я живу в городе", city)
}