package main

import "fmt"


// Программа "Сколько лет будет собаке"
// 1) Нужно сохранить возвраст и сколько лет должно пройти
// 2) Вывести сколько будет собаке через N лет
//
// Вывод должен быть таким:
//
// Через 2 года вашей собаке будет 6 лет/года
func main() {
	var currentDogAge int = 4
	var yearsLater int = 5
	var futureDogAge = currentDogAge + yearsLater

	fmt.Println("Через", yearsLater, "года", "вашей собаке будет", futureDogAge, "лет/года")
}