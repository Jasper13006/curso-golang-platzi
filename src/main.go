package main

import "fmt"

func main() {
	// Declaracion de constantes
	const pi float64 = 3.14159265358979323846
	const pi2 = 3.4
	fmt.Println("El valor de pi es: ", pi)
	fmt.Println("El valor de pi es: ", pi2)

	// Declaracion de variables enteras
	var altura int = 10
	base := 12 // los dos puntos dicen que la variable nunca fue asignada antes
	var area int

	fmt.Println("El valor de la base es: ", base)
	fmt.Println("El valor de la altura es: ", altura)
	fmt.Println("El valor de la area es: ", area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("El valor de a es: ", a)
	fmt.Println("El valor de b es: ", b)
	fmt.Println("El valor de c es: ", c)
	fmt.Println("El valor de d es: ", d)

	// Area cuadrado.
	const basdeCuadrado = 10
	areaCuadrado := basdeCuadrado * basdeCuadrado
	fmt.Println("El valor de la area del cuadrado es: ", areaCuadrado)
}
