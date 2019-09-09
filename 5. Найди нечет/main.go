package main

import "../dexes/console"

func main () {
	var first, second int

	console.Writeln("Последовательно введите четное и нечетное число:")
	first = console.ReadInt()
	second = console.ReadInt()

	// 1011 | 1 = 1011
	// 1000 | 1 = 1001
	if first | 1 == first {
		console.Write("Нечетное число = ", first)
		return
	}


	console.Write("Нечетное число = ", second)
}
