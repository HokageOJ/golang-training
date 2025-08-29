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
	*money = *money + (*money * rubbles(percent / 10))
}

func main() {
	var money rubbles = 100.00

	fmt.Println(money.string())
	fmt.Println(money.toDollars(90.00))
	fmt.Println(money.toEuro(100))	
	
	money.addTax(20) 
	fmt.Println(money)
}