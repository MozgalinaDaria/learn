package console

import "fmt"

func ReadString() string {
	var result string
	fmt.Scan(&result)

	return result
}

func ReadInt() int {
	var result int
	fmt.Scan(&result)

	return result
}

func Write(a ...interface{}) {
	fmt.Print(a...)
}

func Writeln(a ...interface{}) {
	fmt.Println(a...)
}
