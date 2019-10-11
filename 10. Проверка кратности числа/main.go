package main

import "../base/console"

func main() {
	var a, b int

	a = console.ReadInt("Введите первое число: ")
	b = console.ReadInt("Введите второе число: ")

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
	return a % b == 0
}
