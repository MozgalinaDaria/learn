package main

import (
	"../dexes/console"
)

func main() {
	var a, b, result int
	var operation string
	var isSum, isDifference, isMultiple, isDivision bool

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
		console.Write("Результат вычислений: ", a, " ", operation, " ", b, " = ", result)
		return
	}

	if isDifference {
		result = a - b
		console.Write("Результат вычислений: ", a, " ", operation, " ", b, " = ", result)
		return
	}

	if isMultiple {
		result = a * b
		console.Write("Результат вычислений: ", a, " ", operation, " ", b, " = ", result)
		return
	}

	if isDivision {
		result = a / b
		console.Write("Результат вычислений: ", a, " ", operation, " ", b, " = ", result)
		return
	}
}
