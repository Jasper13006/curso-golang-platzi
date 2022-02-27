package main

import (
	"fmt"

	mypackage "./mypackage"
)

func main() {
	var myCar mypackage.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 2018

	// Print the car's brand and year
	fmt.Println(myCar.Brand, myCar.Year)
}
