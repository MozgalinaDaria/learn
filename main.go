package main

import "./dexes"


func main() {
	var name string

	dexes.Write("Enter your name: ")
	name = dexes.ReadString()
	dexes.Writeln("Hello, ", name)
}
