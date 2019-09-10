package main

import "../dexes/console"

func main () {
	var first, second int
	var oddness bool

	console.Writeln("Последовательно введите четное и нечетное число:")
	first = console.ReadInt()
	second = console.ReadInt()

	oddness = isOdd(first)

	if oddness {
		console.Write("Нечетное число = ", first)
		return
}
	console.Write("Нечетное число = ", second)

	func isOdd(number int) bool {

		return (number | 1) == number
	}



}
