package main

import "fmt"


// Программа "Конвертер валют" (рубли в доллары)
// 1) Cохранить рубли и курс обмена
// 2) Перевести в доллары
//
// Вывод должен быть таким:
//
// 4500 рублей будет 50 долларов
func main() {
	var rubbles float64 = 4500
	var exchangeRate float64 = 90

	fmt.Printf("При обмене %.2f рубля(ей) по курсу %.2f рублей за доллар, получиться: %.2f долларов\n", rubbles, exchangeRate, rubbles/exchangeRate)
}