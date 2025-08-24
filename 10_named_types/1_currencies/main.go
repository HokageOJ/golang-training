package main

import "fmt"

/*
	----------
	| Валюты |
	----------
	Создайте именованные типы для рублей и долларов.
	Заведите соответствующие переменные.

	Попробуйте их сложить и убедитесь, что компилятор выдаёт ошибку.

	Напишите функции:
	— для конвертации рублей в доллары
	— для конвертации долларов в рубли

	Внутри функций используйте приведение типов.
*/

type rubbles float64
type dollars float64

func rubblesToDollars(money rubbles, rate rubbles) dollars { 
	return dollars(money / rate) 
}

func dollarsToRubbles(money dollars, rate rubbles) rubbles {
	return rubbles(money * dollars(rate))
}

func main() {
	var rubblesMoney rubbles = 1000.0
	var dollarsMoney dollars = 1000.0

	// result := rubblesMoney + dollarsMoney

	fmt.Println(rubblesToDollars(rubblesMoney, rubbles(90)))
	fmt.Println(dollarsToRubbles(dollarsMoney, rubbles(82)))
}