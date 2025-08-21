package main

import "fmt"

/*
	-----------------------------------------------
	| Тема занятия "Логический тип данных (bool)" |
	-----------------------------------------------
	Тип данных bool может хранить только true или false
	Используется для хранения состояний (вкл/выкл, есть/отсутствует, закрыто/открыто)
*/

// Программа "Статус заказа"
func main() {
	var isPaid bool = true
    var isPacked bool = false
    var isShipped bool = true

    fmt.Println("Заказ оплачен?", isPaid)
    fmt.Println("Заказ упакован?", isPacked)
    fmt.Println("Заказ отправлен?", isShipped)
}