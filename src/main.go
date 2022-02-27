package main

import "fmt"

func main() {

	//
	var arr [4]int
	fmt.Println(arr, len(arr), cap(arr))

	// slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 11)
	fmt.Println(slice)

	// Append
	newSlice := []int{12, 13, 14}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
