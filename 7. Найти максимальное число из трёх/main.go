package main

import "../dexes/console"

func main() {
	var a, b, c int

	console.Writeln("Введите 3 числа")
	a = console.ReadInt()
	b = console.ReadInt()
	c = console.ReadInt()

	console.Writeln("Наибольшее из 3 чисел = ", max(a, b, c))
}

func max(a, b, c int) int {
	if a > b && a > c {
		return a
	}

	if b > c {
		return b
	}

	return c
}
