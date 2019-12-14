package main

import "../base/console"

func main() {
	console.Writeln("Введите 2 числа с разной четностью: ")
	first := console.ReadInt()
	second := console.ReadInt()

	if IsOdd(first) {
		console.Write("Нечетное число = ", first)
		return
	}

	console.Write("Нечетное число = ", second)
}

func IsOdd(number int) bool {
	return (number | 1) == number
}
