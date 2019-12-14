package main

import "../base/console"

func main() {
	console.Writeln("Введите длины каждой из трех сторон треугольника: ")
	a := console.ReadInt()
	b := console.ReadInt()
	c := console.ReadInt()

	if CompareLength(a, b, c) && CompareLength(b, a, c) && CompareLength(c, a, b) {
		console.Writeln("Треугольник существует")
		return
	}

	console.Writeln("Треугольник не существует")
}

func CompareLength(a, b, c int) bool {
	return (b + c) > a
}
