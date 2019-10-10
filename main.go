package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	a := 0
	for i := 0; i < 10000000000; i++ {
		if a == 0 {
		} else {}
	}

	log.Printf("Binomial took %s", time.Since(start))
}
