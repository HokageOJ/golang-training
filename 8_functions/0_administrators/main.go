package main

import "fmt"

/* База данных админов */
var administrators [3]string = [3]string{"ivanov", "petrov", "popov"}
var passwords [3]string = [3]string{"qwerty123", "12345678", "admin"}

/* Функция: является ли пользователь админом? */
func isAdmin(login string, password string) bool {
	for i := 0; i < len(administrators); i++ {
		if login == administrators[i] && password == passwords[i] {
			return true
		}
	}

	return false
}

func main() {
	var login string
	fmt.Print("Введите логин: ")
	fmt.Scan(&login)
	fmt.Println()

	var password string
	fmt.Print("Введите пароль: ")
	fmt.Scan(&password)
	fmt.Println()

	var authorized bool = isAdmin(login, password)

	if !authorized {
		fmt.Println("Неправильный пользователь или пароль! Вход на страницу запрещён!")
		return
	}

	if authorized {
		fmt.Println("Вы вошли в панель управления!")
	}
}