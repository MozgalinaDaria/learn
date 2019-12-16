package main

import "../base/console"

func main() {
	var term, sum, i, rangeForExample float64

	rangeForExample = console.ReadFloat("Введите любое натуральное число: ")
	for i = 1; i <= rangeForExample; i++ {
		//console.Writeln("В")
		term = 1 / (i * i)
		//console.Writeln("term = ", term)
		sum += term
		//console.Writeln("sum = ", sum)
	}
	console.Writeln("Сумма ряда составила: ", sum)
}
