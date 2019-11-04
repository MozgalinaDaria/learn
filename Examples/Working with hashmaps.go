package examples

import "fmt"

func hashmap() {
	var key string
	var value int
	var exists bool

	// Создаём map
	data := make(map[string]int)

	// вписываем значение 5 в ключ test
	data["test"] = 5
	// получаем значение из ключа test
	value = data["test"]
	fmt.Println(value)
	fmt.Println("----------------------------")

	// второй вариант получения значения из ключа test, во вторую переменную кладётся True если ключ есть и False если ключа нет
	value, exists = data["test"]
	fmt.Println(value, exists)
	fmt.Println("----------------------------")

	// когда ключа нет...
	value, exists = data["not exists"]
	fmt.Println(value, exists)
	fmt.Println("----------------------------")

	// ключа всё ещё нет, даже после верхней проверки, когда мы обратились к map, мы туда ещё ничего не писали
	value, exists = data["not exists"]
	fmt.Println(value, exists)
	fmt.Println("----------------------------")

	// даже если ключа нет, то всё равно можно делать инкремент, по умолчанию там будет 0
	data["not exists"]++
	fmt.Println(data["not exists"])
	fmt.Println("----------------------------")

	// теперь ключ есть
	value, exists = data["not exists"]
	fmt.Println(value, exists)
	fmt.Println("----------------------------")

	// цикл по всем ключам map
	for key, value = range data {
		fmt.Println(key, "|", value)
	}
}
