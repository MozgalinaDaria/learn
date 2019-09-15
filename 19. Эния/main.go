package main

import "../base/console"

func main() {
	var a, b, n int

	console.Writeln("Введите длину и ширину панелей, количество панелей")
	a = console.ReadInt()
	b = console.ReadInt()
	n = console.ReadInt()

	console.Writeln("Необходимый вес = ", a*b*n*2)
}
