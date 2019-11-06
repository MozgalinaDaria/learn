package main

import "../base/console"

func main() {

	a := console.ReadInt("Введите длину панелей: ")
	b := console.ReadInt("Введите ширину панелей: ")
	n := console.ReadInt("Введите количество панелей: ")

	console.Writeln("Необходимый вес = ", a*b*n*2)
}
