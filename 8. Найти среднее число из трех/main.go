package main

import "../base/console"

func main() {
	var a, b, c int

	console.Writeln("Введите 3 числа")
	a = console.ReadInt()
	b = console.ReadInt()
	c = console.ReadInt()

	console.Writeln("Среднее число = ", average(a, b, c))
}

func average(a, b, c int) int {
	if a < b && b < c || a > b && b > c {
		return b
	}

	if a < c && c < b || a > c && c > b {
		return c
	}

	return a
}
