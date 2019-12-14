package main

import "../base/console"

func main() {
	console.Writeln("Введите a, b")
	a := console.ReadFloat("Введите a: ")
	b := console.ReadFloat("Введите b: ")

	if a == 0 {
		console.Writeln("Решения нет")
		return
	}

	min := -b / (2 * a)

	console.Writeln("Функция принимает минимальное значение при x = ", min)
}
