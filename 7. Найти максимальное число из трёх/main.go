package main

import "../base/console"

func main() {
	a := console.ReadInt("Введите первое число: ")
	b := console.ReadInt("Введите второе число: ")
	c := console.ReadInt("Введите третье число: ")

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
