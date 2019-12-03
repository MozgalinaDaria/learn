package program

import "../logger"

func DoSomething(number int) {
	if number == 0 {
		var log logger.Interface

		log = logger.GetLogger()
		log.Err("some err")

		panic("Число не может быть нулём")
	}

	// ...
}
