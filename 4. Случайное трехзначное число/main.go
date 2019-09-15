package main

import "../base/console"
import "../base/rand"

func main() {
	console.Write("Рандомное трехзначное число = ", rand.RandomInt(10, 99)*10)
}
