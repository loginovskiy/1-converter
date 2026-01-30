package main

import (
	"fmt"
)

func main() {
	const usdToEur float64 = 0.82
	const usdToRub float64 = 75.50
	const eurToRub float64 = 75.50 / 0.82
	fmt.Println(eurToRub)
}
