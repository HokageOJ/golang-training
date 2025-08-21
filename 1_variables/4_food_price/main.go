package main

import "fmt"


// Программа "Стоимость покупки"
// 1) Cохранить кол-во пачек бумаги А4 "Снегурочка" и цену пачки
// 2) Посчитать стоимость всех пачек бумаги
//
// Вывод должен быть таким:
//
// Я купил 4 пачки/пачек бумаги А4 "Снегурочка" на сумму 1000 рублей
func main() {
	var paperPacketsCount int = 15
	var paperPacketPrice = 300

	fmt.Println("Я купил", paperPacketsCount, "пачки/пачек бумаги А4 "Снегурочка" на сумму", paperPacketsCount * paperPacketPrice, "рублей")
}