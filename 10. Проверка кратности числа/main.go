package main

import "../base/console"

func main() {
	a := console.ReadInt("Введите первое число: ")
	b := console.ReadInt("Введите второе число: ")

	if b == 0 {
		console.Writeln("Делить на ноль нельзя!")
		return
	}

	if IsMultiplicity(a, b) {
		console.Writeln("Первое число кратно второму")
		return
	}

	console.Writeln("Первое число не кратно второму")
}

func IsMultiplicity(a, b int) bool {
	return a%b == 0
}
