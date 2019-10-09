package main

import "../base/console"

func main() {
	var count, first, second, next int
	count = console.ReadInt("Сколько чисел Фибоначчи отобразить? ")

	if count == 0 {
		console.Writeln("Ни одного числа Фибоначчи")
		return
	}
	if count == 1 {
		console.Writeln("Число Фибоначчи: 0")
		return
	}
	if count == 2 {
		console.Writeln("Числа Фибоначчи: 0, 1")
		return
	}
	if count > 2 {
		first = 0
		second = 1
		console.Write("Числа Фибоначчи: ", first, " ", second, " ")

		for i:= 0; i <= count - 3; i++ {
			next = first + second
			console.Write(next, " ")
			first = second
			second = next
		}
	}
}
