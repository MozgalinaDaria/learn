package main

import "../base/console"

func main (){
	var number, count int

	number = console.ReadInt("Введите число: ")
	for i:= 0; i < number; i ++{
		count = count + 1

		for j := i * number + 1 ; j <= number * count; j++{
			console.Write(j, " ")
		}
		console.Writeln()
	}
}

