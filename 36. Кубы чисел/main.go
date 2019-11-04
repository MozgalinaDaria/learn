package main

import (
	"../base/console"
)

func main() {
	var first, second float64

	first = console.ReadFloat("Введите первое число: ")
	second = console.ReadFloat("Введите второе число: ")
	console.Writeln("Кубы чисел от ", first, "до ", second, " : ")

	for i := first; i <= second; i++ {
		console.Write(i*i*i, " ")
	}
}
