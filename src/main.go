package main

import "fmt"

func main() {
	// for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// for range
	for i, v := range "go" {
		fmt.Println(i, v)
	}

	// for forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever > 10 {
			break
		}
	}
}
