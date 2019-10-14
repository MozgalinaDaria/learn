package main

import "../base/console"

func main() {
	var number string
	var length int
	var figureToRemove uint8

	number = console.ReadString("Введите число: ")
	figureToRemove = uint8(console.ReadChar("Введите цифру, которую нужно удалить из числа: "))

	length = len(number)

	for i := 0; i < length; i++ {
		if number[i] == figureToRemove {
			console.Write("")
		} else {
			console.Write(string(number[i]))
		}
	}
}
