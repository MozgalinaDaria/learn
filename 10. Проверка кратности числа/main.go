package main

import "../base/console"

func main() {
	var a, b int

	console.Writeln("Введите 2 числа")
	a = console.ReadInt()
	b = console.ReadInt()

	if b == 0 {
		console.Writeln("Делить на ноль нельзя! Начните заново")
		return
	}

	if isMultiplicity(a, b) {
		console.Writeln("Первое число кратно второму")
		return
	}

	console.Writeln("Первое число не кратно второму")
}

func isMultiplicity(a, b int) bool {
	return a%b == 0
}
