package main

import "fmt"


// Программа "Длина клички животного"
// 1) Нужно сохранить кличку собаки
// 2) Вычислить её длину
//
// Вывод должен быть таким:
//
// В кличке Рекс 4 буквы
func main() {
	var animalName string = "Murzik"
	var animalNameLength int = len(animalName)

	fmt.Println("В кличке", animalName, animalNameLength, "букв(ы)")
}
