package main

import "fmt"


// Программа "Сколько лет будет человеку"
// 1) Нужно сохранить возвраст и сколько лет должно пройти
// 2) Вывести сколько будет человеку через N лет
//
// Вывод должен быть таким:
//
// Через 5 лет вам будет 25 лет
func main() {
	var age int = 25
	var yearsLater int = 5
	var futureAge int = age + yearsLater

	fmt.Println("Через", yearsLater, "лет вам будет", futureAge, "лет")
}