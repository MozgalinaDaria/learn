package main

import (
	"../base/console"
)

func main() {
	var a, b int
	var operation string

	console.Write("Введите первое число: ")
	a = console.ReadInt()

	console.Write("Введите знак требуемой операции: ")
	operation = console.ReadString()

	console.Write("Введите второе число: ")
	b = console.ReadInt()

	if operation == "+" {
		Output(a, operation, b, Sum(a, b))
		return
	}

	if operation == "-" {
		Output(a, operation, b, Difference(a, b))
		return
	}

	if operation == "*" {
		Output(a, operation, b, Multiple(a, b))
		return
	}

	if operation == "/" {
		Output(a, operation, b, Division(a, b))
		return
	}

	console.Writeln("Введен некорректный знак операции")
}

func Output(first int, operation string, second int, result int) {
	console.Writeln(first, operation, second, "=", result)
}

func Sum(a, b int) int {
	return a + b
}

func Difference(a, b int) int {
	return a - b
}

func Multiple(a, b int) int {
	return a * b
}

func Division(a, b int) int {
	return a / b
}
