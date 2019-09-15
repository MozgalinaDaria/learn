package main

import "../base/console"

func main() {
	var a, b, min float64

	console.Writeln("Введите a, b")
	a = console.ReadFloat()
	b = console.ReadFloat()

	if a == 0 {
		console.Writeln("Решения нет")
		return
	}

	min = -b / (2 * a)

	console.Writeln("Функция принимает минимальное значение при x = ", min)
}
