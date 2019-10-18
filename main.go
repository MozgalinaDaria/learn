package main

import (
	"./base/console"
	"unicode/utf8"
)
func main() {
	//	var number string
	//	var length int
	//	var figureToRemove rune
	//	var runes []rune
	//
	//	number = console.ReadString("Введите число: ")
	//	figureToRemove = console.ReadChar("Введите цифру, которую нужно удалить из числа: ")
	//
	//	length = len(number)
	//	runes, length = utf8.DecodeRuneInString(number)
	//	for i := 0; i < length; i++ {
	//		if runes[i] == figureToRemove {
	//			console.Write("")
	//		}
	//	}

	str :=  "bd фыва"

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		console.Writeln(str)
		console.Writeln(string(r), size)
		console.Writeln()

		str = str[size:]
	}
}
