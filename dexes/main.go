package dexes

import "fmt"

func ReadString() string {
	var result string
	fmt.Scanln(&result)

	return result
}

func Write(a ...interface{}) {
	fmt.Print(a...)
}

func Writeln(a ...interface{}) {
	fmt.Println(a...)
}
