package main

import "./dexes/console"


func main() {
	var name string
	var age int

	console.Write("Enter your name: ")
	name = console.ReadString()
	console.Writeln("Hello, ", name)

	console.Write("Enter your age: ")
	age = console.ReadInt()
	console.Writeln("Really?", age, "?")
}
