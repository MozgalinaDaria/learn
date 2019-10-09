package main

import "../base/console"

func main() {
	number:= console.ReadInt("Назовите число: ")
	degree:= console.ReadInt("Назовите степень, в которую нужно возвести ", number, ": ")

	if degree == 0 {
		console.Writeln(number, " в степени ", degree, "равен 1")
		return
	}
	if degree != 0 {
		console.Write(number, " в степени ", degree, " = ")
		for i:= 0; i <= degree - 2; i++ {
			number *= number
		}
		if degree > 0 {
			console.Write(number)
		}
		if degree < 0 {
			console.Write("1/", number)
		}
	}
}
