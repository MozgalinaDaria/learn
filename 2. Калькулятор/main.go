package main

import (
	"../base/console"
)

func main() {

	a := console.ReadInt("Введите первое число: ")
	operation := console.ReadString("Введите знак требуемой операции: ")
	b := console.ReadInt("Введите второе число: ")

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
