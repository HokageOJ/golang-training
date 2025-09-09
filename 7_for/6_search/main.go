package main

import "fmt"

/*
	---------------
	| Поиск числа |
	---------------
	Дан массив целых чисел.

	Напишите программу, для поиска в массиве числа, введённого с терминала.

	--------
	| Ввод |
	--------
	Введите число для поиска: X

	---------
	| Вывод |
	---------
	Массив: [42 7 <...>]

	(Если число найдено) Число X найдено! Индекс элемента: Y
	(Число не найдено) Число X не найдено в массиве!
*/

func main() {
	var numbers [10]int = [10]int{42, 4, 15, 23, 2, 11, 30, 4, 19, 42}

	fmt.Printf("Массив: %d\n", numbers)

	var search int

	fmt.Print("Введите число для поиска: ")
	fmt.Scan(&search)

	var isFound bool

	for i := 0; i < len(numbers); i++ {
		if search == numbers[i] {
			fmt.Printf("Число %d найдено! Индекс элемента: %d\n", search, i)
			isFound = true
			break
		} 
	}
	
	if !isFound {
		fmt.Printf("Число %d не найдено в массиве!\n", search)
	}
}