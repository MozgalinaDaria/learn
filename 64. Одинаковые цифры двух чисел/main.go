package main

import (
	"../base/console"
)

const FiguresCount = 10

type figures [FiguresCount]bool

func main() {
	numFirst := NumberToFigures(console.ReadInt("Введите первое число: "))
	numSecond := NumberToFigures(console.ReadInt("Введите второе число: "))
	result := CrossFigures(numFirst, numSecond)

	console.Writeln("Повторяющиеся цифры в обоих числах: ")
	for i := 0; i < FiguresCount; i++ {
		if result[i] {
			console.Write(i, " ")
		}
	}
}

func CrossFigures(first, second figures) figures {
	var result figures

	for i := 0; i < FiguresCount; i++ {
		result[i] = first[i] && second[i]
	}

	return result
}

func NumberToFigures(number int) figures {
	var result figures

	for number > 0 {
		result[number%10] = true
		number /= 10
	}

	return result
}
