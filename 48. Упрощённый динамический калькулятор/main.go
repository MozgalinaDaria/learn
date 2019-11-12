package main

import (
	"../base/cast"
	"../base/console"
)

func main() {
	var str string

	s := ""
	console.Writeln("Введите пример вида: ")
	console.Writeln("1")
	console.Writeln("+")
	console.Writeln("1")
	console.Writeln("=")

	for s != "=" {
		s = console.ReadString()
		str = str + s
	}
	example := [] rune (str)

	Calculator(example)
}

func Calculator(example []rune) {
	var i, total int
	var numberFirst []rune

	length := len(example)

	for i = 0; i < length; i++ {
		if example [i] != '+' && example [i] != '-' {
			numberFirst = append(numberFirst, example [i])
		} else {
			break
		}
	}

	operation := example [i]
	numberSecond := (example [i+1 : length-1])

	switch operation {
	case '+':
		total = cast.RuneToInt(numberFirst) + cast.RuneToInt(numberSecond)
		console.Writeln(total)
	case '-':
		total = cast.RuneToInt(numberFirst) - cast.RuneToInt(numberSecond)
		console.Writeln(total)
	}
}
