package main

import (
	"fmt"
	"math"
)

func main() {

	const pi float64 = 3.14
	const pi2 = 3.14159

	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Area cuadrado

	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado: ", areaCuadrado)

	//Suma
	x := 10
	y := 50
	result := x + y

	fmt.Println("Suma: ", result)

	//Resta
	result = y - x

	fmt.Println("Resta: ", result)

	//Multiplicacion

	result = y * x

	fmt.Println("Multiplicacion: ", result)

	//Division

	result = y / x

	fmt.Println("Division: ", result)

	//Modulo

	result = y % x

	fmt.Println("Modulo: ", result)

	//Autoincremento

	x++

	fmt.Println("Autoincremento: ", x)

	//Decremento

	x--

	fmt.Println("Autoincremento: ", x)

	//Areas
	//Rectangulo - Trapecio - Circulo

	baseRectangulo := 12
	alturaRectangulo := 10
	resultArea := baseRectangulo * alturaRectangulo

	fmt.Println("Area del rectangulo", resultArea)

	primeraLongitud := 13
	segundaLongitud := 20

	resultArea = ((primeraLongitud + segundaLongitud) * altura) / 2
	fmt.Println("Area del trapecio", resultArea)

	radio := 5.0
	resultAreaFloat := pi2 * math.Pow(radio, 2)

	fmt.Println("Area del circulo", resultAreaFloat)

}
