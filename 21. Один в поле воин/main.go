package main

import "../base/console"

func main() {
	var startCoordinates string
	var x, y int

	console.Writeln("Введите начальную координату ферзя")
	startCoordinates = console.ReadString()

	x = LetterToInt(startCoordinates[0])
	y = NumeralToInt(startCoordinates[1])

	ChessMoves(x, y)
}

func LetterToInt(c uint8) int {
	return int(c - 64) // 65 - код символа "А". Таким образом, "а" вернет 0; "b" вернет 1 и т.д.
}

func NumeralToInt(c uint8) int {
	return int(c - 48) // 49 - код символа "1". Таким образом, "0" вернет 0; "1" вернет 1 и т.д.
}

func IntToLetter(n int) string {
	return string(n + 64)
}

func IntToNumeral(n int) string {
	return string(n + 48)
}

func CoordinatesToString(x, y int) string {
	return IntToLetter(x) + IntToNumeral(y)
}

func ChessMoves(x, y int) {
	var count int

	if (x+1) >= 1 && (x+1) <= 8 && (y+2) >= 1 && (y+2) <= 8 { // Вверх и направо
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x+1, y+2))
	}

	if (x-1) >= 1 && (x-1) <= 8 && (y+2) >= 1 && (y+2) <= 8 { // Вверх и налево
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x-1, y+2))
	}

	if (x+1) >= 1 && (x+1) <= 8 && (y-2) >= 1 && (y-2) <= 8 {
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x+1, y-2))
	}

	if (x-1) >= 1 && (x-1) <= 8 && (y-2) >= 1 && (y-2) <= 8 {
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x-1, y-2))
	}

	if (x+2) >= 1 && (x+2) <= 8 && (y+1) >= 1 && (y+1) <= 8 { // Вверх и направо
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x+2, y+1))
	}

	if (x-2) >= 1 && (x-2) <= 8 && (y+1) >= 1 && (y+1) <= 8 { // Вверх и налево
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x-2, y+1))
	}

	if (x+2) >= 1 && (x+2) <= 8 && (y-1) >= 1 && (y-1) <= 8 {
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x+2, y-1))
	}

	if (x-2) >= 1 && (x-2) <= 8 && (y-1) >= 1 && (y-1) <= 8 {
		count++
		console.Writeln("Возможный ход: ", CoordinatesToString(x-2, y-1))
	}

	console.Writeln("Всего возможных ходов: ", count)
}
