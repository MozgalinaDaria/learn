package main

import (
	"../base/cast"
	"../base/console"
)

func main() {
	var equal, number []rune
	var i int

	equal = []rune(console.ReadString("Введите пример вида: число, знак, число через пробел: "))
	length := len(equal)

	for i = 0; i < length; i++ {
		if equal[i] != ' ' {
			number = append(number, equal[i])
		} else {
			break
		}
	}

	numberFirst := cast.RuneToInt(number)
	operation := equal[i+1]
	number = (equal[i+3:])
	numberSecond := cast.RuneToInt(number)

	switch operation {
	case '+':
		sum := numberFirst + numberSecond
		console.Writeln(sum)
	case '-':
		console.Writeln(numberFirst - numberSecond)
	case '*':
		console.Writeln(numberFirst * numberSecond)
	case '/':
		console.Writeln(numberFirst / numberSecond)
	}
}
