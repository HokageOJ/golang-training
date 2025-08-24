package main

import "fmt"

/*
	----------
	| Методы |
	----------
	Методы — это функции, которые принадлежат именованному типу данных.

	Добавляют поведение для типов, то есть задавать действия, которые тип может выполнять.
*/

type MyInt int

/*
  Отличительная особенность методов — наличие получателя (receiver).
  Он пишется в начале и необходим, чтобы обратится к значению переменной.

  Тип получателя влияет на то, может ли метод изменять исходное значение.
*/

/*
  Получатель может быть значением (value receiver):
*/

// Метод для проверки, является ли число положительным
func (num MyInt) IsPositive() bool {
	return num > 0
}

// Метод для проверки, является ли число отрицательным
func (num MyInt) IsNegative() bool {
	return num < 0
}

// Метод для проверки, является ли число чётным
func (num MyInt) IsEven() bool {
	return num % 2 == 0
}

/*
  Получатель может быть указателем (pointer receiver):
*/

// Инкремент
func (num *MyInt) Increment() {
	*num += 1
}

// Декремент
func (num *MyInt) Decrement() {
	*num -= 1
}

func main() {
	var number MyInt = 10
	fmt.Println(number)

	fmt.Println("Число является положительным?", number.IsPositive())
	fmt.Println("Число является отрицательным?", number.IsNegative())
	fmt.Println("Число является чётным?", number.IsEven())

	number.Increment()
	fmt.Println(number)

	number.Decrement()
	fmt.Println(number)
}