package main

import "../base/console"

func main() {

	var x []int // делимое

	n := console.ReadInt("Введите степень: ")
	m := console.ReadInt("Введите делитель: ")
	y := console.ReadInt("Введите получившийся остаток от деления: ")

	console.Writeln("Следующие числа могут быть делимыми в этом равенстве:")

	for i := 0; i <= m-1; i++ {
		nToDegree := 1

		for j := 0; j < n; j++ {
			nToDegree *= i
		}

		if nToDegree%m == y {
			x = append(x, i)
		}
	}

	if len(x) > 0 {
		console.Write(x)
	} else {
		console.Writeln("-1")
	}
}
