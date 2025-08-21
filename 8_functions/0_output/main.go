package main

import "fmt"

/* БД */
var users [3]string = [3]string{"Иван Кондратьев", "Алина Машева", "Женёчек Пирожков"}
var ages [3]int = [3]int{37, 17, 56}
var city [3]string = [3]string{"Пенза", "Москва", "Одесса"}


func isUserFromPenza(name string, age int, city string) bool {
    fmt.Printf("Имя: %s, Возраст: %d\n", name, age)
    if age >= 18 {
        fmt.Println("Пользователь", name, "совершеннолетний!")
    } else {
        fmt.Println("Пользователь", name, "несовершеннолетний!")
    }
    fmt.Println()

    if city == "Пенза" {
        return true
    }

    return false
}

func main() {
    // Количество пензенцев
    var count int

    for i := 0; i < len(users); i++ {
        if isUserFromPenza(users[i], ages[i], city[i]) {
            count++
        }
    }

    fmt.Println("Количество пензенцев:", count)
}