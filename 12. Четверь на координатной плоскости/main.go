package main

import "../base/console"

// TODO Вынести определение четверти в отдельную функцию func DetectPosition(x, y int) int
// TODO Функция должна возвращать константы (своя константа для каждой ситуации)
// TODO func должна определять что вывести пользователю через switch/case

func main() {
	var x, y int

	console.Writeln("Введите координаты точки по оси Х")
	x = console.ReadInt()
	console.Writeln("Введите координаты точки по оси Y")
	y = console.ReadInt()

	if x == 0 {
		if y > 0 {
			console.Writeln("Координата находится на границе 2-й и 1-ой четвертей")
			return
		}

		if y < 0 {
			console.Writeln("Координата находится на границе 3-ей и 4-ой четвертей")
			return
		}
	}

	if x > 0 {
		if y > 0 {
			console.Writeln("Координата находится в 1-й четверти")
			return
		}

		if y == 0 {
			console.Writeln("Координата находится на границе 1-й и 4-й четвертей")
			return
		}

		console.Writeln("Координата находится в 4-й четверти")
	}

	if x < 0 {
		if y < 0 {
			console.Writeln("Координата находится в 3-ей четвертей")
			return
		}

		if y == 0 {
			console.Writeln("Координата находится на границе 2-й и 3-й четвертей")
			return
		}

		console.Writeln("Координата находится во 2-й четверти")
	}

	if x == 0 && y == 0 {
		console.Writeln("Координата находится на границе всех четырех четвертей")
		return
	}
}
