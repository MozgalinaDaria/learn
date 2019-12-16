package main

import "../base/console"

func main() {
	a := console.ReadInt("Введите первое число: ")
	b := console.ReadInt("Введите второе число: ")
	c := console.ReadInt("Введите третье число: ")

	console.Writeln("Среднее число = ", Average(a, b, c))
}

func Average(a, b, c int) int {
	if a < b && b < c || a > b && b > c {
		return b
	}

	if a < c && c < b || a > c && c > b {
		return c
	}

	return a
}
