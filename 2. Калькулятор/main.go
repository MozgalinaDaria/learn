package main

import "../dexes/console"

func main () {
	var a, b, result int
	var operation string

	console.Write("Введите первое число: ")
	a = console.ReadInt()

	console.Write("Введите знак требуемой операции: ")
	operation = console.ReadString()

	console.Write("Введите второе число: ")
	b = console.ReadInt()

	if operation == "+" {
		result = a + b
	} else if operation == "-" {
		result = a - b
	} else if operation == "*" {
		result = a * b
	} else {
		result = a / b
	}

	console.Write("Результат вычислений: ", result)
}
