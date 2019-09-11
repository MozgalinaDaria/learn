package main

import "../dexes/console"

func main () {
	var first, second int

	console.Writeln("Последовательно введите четное и нечетное число:")
	first = console.ReadInt()
	second = console.ReadInt()

	if isOdd(first) {
		console.Write("Нечетное число = ", first)
		return
	}

	console.Write("Нечетное число = ", second)
}

func isOdd(number int) bool {
	return (number | 1) == number
}
