package main

import "fmt"


// Программа "Длина имени"
// 1) Нужно сохранить имя человека
// 2) Вычислить его длину
//
// Вывод должен быть таким:
//
// В вашем имени 4 букв(ы)
func main() {
	var name string = "Oleg"
	var nameLength int = len(name)

	fmt.Println("В вашем имени", name, nameLength, "букв(ы)")
}