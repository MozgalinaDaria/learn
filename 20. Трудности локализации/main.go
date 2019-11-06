package main

import "../base/console"

func main() {

	count := console.ReadInt("Введите количество воинов")

	switch {
	case count <= 4:
		console.Writeln("few")
	case count <= 9:
		console.Writeln("several")
	case count <= 19:
		console.Writeln("pack")
	case count <= 49:
		console.Writeln("lots")
	case count <= 99:
		console.Writeln("horde")
	case count <= 249:
		console.Writeln("throng")
	case count <= 499:
		console.Writeln("swarm")
	case count <= 999:
		console.Writeln("zounds")
	default:
		console.Writeln("legion")
	}
}
