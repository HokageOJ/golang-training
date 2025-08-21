package main

import "fmt"

// Программа "Килограммы в граммы"
// 1) Cохранить килограммы
// 2) Перевести в граммы
//
// Вывод должен быть таким:
//
// 5 кг будет 5000 грамм
func main() {
	var kilograms int = 5

	fmt.Println(kilograms, "кг будет", kilograms*1000, "грамм")
}
