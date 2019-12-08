package main

import "../base/console"

func main() {
	number := console.ReadString("Введите число для преобразования: ")
	numberOfFiguresToPrint := console.ReadInt("Введите длину строки вывода: ")

	countOfZeroFigures := numberOfFiguresToPrint - len(number)
	inputZeroes(countOfZeroFigures)
	console.Write(number)

}

func inputZeroes(count int) {
	for i := 0; i < count; i++ {
		console.Write(0)
	}
}
