package main

import (
	"fmt"
)

func userEntry(userEnv *string) {
	fmt.Print("Введите любую строку: ")
	fmt.Scan(userEnv)
}

func casheConvertation(machMoney int, firstCurrency, secondCurrency string) {

}

func main() {
	const usdToEur float64 = 0.82
	const usdToRub float64 = 75.50
	const eurToRub float64 = 75.50 / 0.82
	fmt.Println(eurToRub)
	var userEntryString string
	userEntry(&userEntryString)
	fmt.Println("Вы ввели: ", userEntryString)
}
