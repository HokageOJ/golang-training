package main

import (
	"fmt"
)

/*
	---------------
	|    Валюты    |
	---------------
	
*/

type rubbles float64
type dollars float64
type euro float64

func (money rubbles) string() string {
	return fmt.Sprintf("%.2f₽", money)
}

func (money rubbles) toDollars(rate rubbles) dollars {
	return dollars(money / rate)
}

func (money rubbles) toEuro(rate rubbles) euro {
	return euro(money / rate)
}

func (money *rubbles) addTax(percent float64) {
	*money = *money + (*money * rubbles(percent / 100))
}

func (money dollars) string() string {
	return fmt.Sprintf("$%.2f", money)
}

func (money dollars) toRubbles(rate dollars) rubbles {
	return rubbles(money / rate)
}

func (money dollars) toEuro(rate dollars) euro {
	return euro(money * rate)
}

func (money *dollars) addTax(percent float64) {
	*money = *money + (*money * dollars(percent / 100))
}
func (money euro) string() string {
	return fmt.Sprintf("€%.2f", money)
}

func (money euro) toRubbles(rate euro) rubbles {
	return rubbles(money/rate)
}

func (money euro) toDollars(rate euro) dollars {
	return dollars(money*rate)
}

func (money *euro) addTax(percent float64) {
	*money = *money + (*money * euro(percent / 100))
}

func main() {
	var money rubbles = 100.00

	fmt.Println("Рубли:", money.string())
	fmt.Println("Рубли в доллары:", money.toDollars(90.00).string())
	fmt.Println("Рубли в евро:", money.toEuro(100).string())	
	
	money.addTax(20) 
	fmt.Println("Добавлен налог 20% к рублям:", money.string())

	fmt.Println()

	var cash dollars = 100.00

	fmt.Println("Доллары:", cash.string())
	fmt.Println("Доллары в рубли:", cash.toRubbles(0.012).string())
	fmt.Println("Доллары в евро:", cash.toEuro(0.85).string())

	cash.addTax(20)
	fmt.Println("Добавлен налог 20% к долларам:", cash.string())

	fmt.Println()

	var currency euro = 100.00

	fmt.Println("Евро", currency.string())
	fmt.Println("Евро в рубли:", currency.toRubbles(0.011).string())
	fmt.Println("Евро в доллары:", currency.toDollars(1.16).string())
	
	currency.addTax(20)
	fmt.Println("Добавлен налог 20% к евро:", currency.string())
} 