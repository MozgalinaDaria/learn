package main

import "../base/console"

func main() {
	var factorial = 1

	number := console.ReadInt("Введите число от 3 до 20: ")

	for i := 1; i <= number; i++ {
		factorial = factorial * i
	}
	console.Writeln("Факториал числа ", number, " = ", factorial)
}
