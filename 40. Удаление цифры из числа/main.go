package main

import (
	"../base/console"
	"unicode/utf8"
)

func main() {
	var number string
	var figureToRemove string

	number = console.ReadString("Введите число: ")
	figureToRemove = console.ReadString("Введите цифру, которую нужно удалить из числа: ")

	for len(number) > 0 {
		figure, size := utf8.DecodeRuneInString(number) // Возвращает первый символ и его размер в байтах

		if string(figure) != figureToRemove {
			console.Write(string(figure))
		}

		number = number[size:] //Сдвигаем строку вправо на столько байт, сколько было в предыдущем, уже обработанном символе
	}

	//	str :=  "bd фыва"
	//
	//	for len(str) > 0 {
	//		r, size := utf8.DecodeRuneInString(str) // Возвращает первый символ и его размер в байтах
	//		console.Writeln(str)
	//		console.Writeln(string(r), size)
	//		console.Writeln()
	//
	//		str = str[size:]
	//	}
}
