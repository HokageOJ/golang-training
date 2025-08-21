package main

import (
	"fmt"
)

/*
	-------------------------
	| Урок Бесконечный ввод |
	-------------------------
	Реализуется с помощью чистого цикла for

	По окончанию урока переделайте предыдущие задания с его использованием.
*/

func main() {
	var numbers [10]int = [10]int{6, 23, 56, 1, 7, 67, 76, 3, 8, 45}

	var num int

	for {
		fmt.Println("Массив:", numbers)

		fmt.Print("Введите число для поиска: ")
		fmt.Scan(&num)

		var isFound bool

		for i := 0; i < len(numbers); i++ {
			if (num == numbers[i]) {
				isFound = true
				fmt.Printf("Число %d найдено! Индекс %d\n", num, i)
			}
		}

		if (isFound == false) {
			fmt.Printf("Число %d не найдено!\n", num)
		}

		var input string

		fmt.Print("Хотите продолжить поиск? Y/y - да, N/n - нет: ")
		fmt.Scan(&input)

		if (input == "Y" || input == "y") {
			fmt.Println("\033c")
		} else {
			break
		}
	}
}