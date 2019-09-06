package main

import "../dexes/console"

func main () {
	var a, b, result int
	var operation string
	var isSum, isDifference, isMultiple bool

	console.Write("Введите первое число: ")
	a = console.ReadInt()

	console.Write("Введите знак требуемой операции: ")
	operation = console.ReadString()

	console.Write("Введите второе число: ")
	b = console.ReadInt()

	isSum = operation == "+"
	isDifference = operation == "-"
	isMultiple = operation == "*"

	if isSum {
		result = a + b
	} else if isDifference {
		result = a - b
	} else if isMultiple {
		result = a * b
	} else {
		result = a / b
	}

	console.Write("Результат вычислений: ", result)
}
