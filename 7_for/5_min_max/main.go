package main

import "fmt"

/*
	------------------------------------
	| Максимальное и минимальное число |
	------------------------------------
	Дан массив целых чисел.

	Найдите минимальное и максимальное числа.

	---------
	| Вывод |
	---------
	Min: X
	Max: X
*/

func main() {
	var numbers [10]int = [10]int{5, 12, 8, 3, 17, 9, 1, 14, 6, 10}

	var min int = numbers[0]
	var max int = numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		} 
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}

	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}