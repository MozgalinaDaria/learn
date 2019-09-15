package main

import "../base/console"
import "../base/rand"

func main() {
	var number int
	var numeral1, numeral2, numeral3 int

	number = rand.RandomInt(100, 999)

	numeral1 = number / 100
	numeral2 = number / 10 % 10
	numeral3 = number % 10

	console.Writeln("Получившееся трехзначное число: ", number)
	console.Writeln("Сумма трех чисел: ", numeral1+numeral2+numeral3)
	console.Writeln("Произведение трех чисел: ", numeral1*numeral2*numeral3)
}
