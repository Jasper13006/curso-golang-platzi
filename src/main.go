package main

import "fmt"

func normalFuncion(message string) {
	fmt.Println(message)
}

func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFuncion("Hello World")
	tripeArgument(1, 2, "Hello")
	fmt.Println(returnValue(3))

	value1, value2 := doubleReturn(3)
	fmt.Println(value1, value2)
}
