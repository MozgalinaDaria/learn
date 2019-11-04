package main

import "../base/console"
import "../base/rand"

func main() {
	var number int

	number = GetTheRandomNumber()
	console.Writeln("Сгенерированное число = ", number)

	console.Writeln("Числа от 0 до ", number, ": ")
	if number >= 0 {
		FromZeroToNumberPositive(number)
		console.Writeln()
		console.Writeln("Числа от ", number, "до 0: ")
		FromNumberToZeroPositive(number)
		console.Writeln()
		console.Writeln("Все четные числа от 0 до ", number, ": ")
		EvenNumbersPositive(number)
	} else {
		FromZeroToNumberNegative(number)
		console.Writeln()
		console.Writeln("Числа от ", number, "до 0: ")
		FromNumberToZeroNegative(number)
		console.Writeln()
		console.Writeln("Все четные числа от 0 до ", number, ": ")
		EvenNumbersNegative(number)
	}
}

func GetTheRandomNumber() int {
	period := rand.RandomInt(0, 1)

	if period == 0 {
		return rand.RandomInt(-15, -5)
	}
	return rand.RandomInt(5, 15)
}

func FromZeroToNumberNegative(number int) {
	for i := 0; i >= number; i-- {
		console.Write(i, " ")
	}
}

func FromZeroToNumberPositive(number int) {
	for i := 0; i <= number; i++ {
		console.Write(i, " ")
	}
}

func FromNumberToZeroPositive(number int) {
	for i := number; i >= 0; i-- {
		console.Write(i, " ")
	}
}

func FromNumberToZeroNegative(number int) {
	for i := number; i <= 0; i++ {
		console.Write(i, " ")
	}
}

func EvenNumbersNegative(number int) {
	for i := 0; i >= number; i = i - 2 {
		console.Write(i, " ")
	}
}

func EvenNumbersPositive(number int) {
	for i := 0; i <= number; i = i + 2 {
		console.Write(i, " ")
	}
}
