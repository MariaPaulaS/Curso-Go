package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

const pi float64 = 3.14
const pi2 = 3.14159

func main() {

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

	resultArea := rectangleArea(12, 10)
	fmt.Println("Area del rectangulo", resultArea)

	resultArea = trapeziumArea(13, 20, altura)
	fmt.Println("Area del trapecio", resultArea)

	resultAreaFloat := circleArea(5.0)

	fmt.Println("Area del circulo", resultAreaFloat)

	//Uso de paquete fmt

	//Impresion con salto de linea

	hello := "Hello"
	world := "World"

	fmt.Println(hello, world)
	fmt.Println(hello, world)

	nombre := "Platzi"
	cursos := 500

	//Impresion con concatenación en un texto

	fmt.Printf("%s tiene mas de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v cursos\n", nombre, cursos)

	//Guardar mensaje concatenado en una variable

	message := fmt.Sprintf("%s tiene mas de %d cursos", nombre, cursos)
	fmt.Println(message)

	//Obtener tipo de dato de una variable
	fmt.Printf("Hello data type: %T \n", hello)

	//Funciones
	printMessage("Hello world function")
	tripeArgument(1, 2, "Hola")
	fmt.Println("Return value", returnValue(2))

	value1, _ := doubleReturn(5)
	fmt.Println("Value1:", value1)

	//Ciclos

	//For básico - Condicional

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//For while

	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For inverso

	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}

	//For forever - Tiene que haber una forma dentro del codigo para detenerlo

	counterForever := 0

	for {
		fmt.Println(counterForever)
		counterForever++

		if counterForever > 1000 { //Condicion para romper el for infinito
			break
		}
	}

	//Convertir texto a numero
	value, error := strconv.Atoi("53")

	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(value)

	isParOrImparNumber(value)

	user := "Maria"
	password := "12345"
	fmt.Printf("Is the user %s authorized?: %t", user, isUserAuthorized(user, password))

}

func printMessage(message string) {
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

func rectangleArea(rectangleBase, heightBase int) int {
	return rectangleBase * heightBase
}

func trapeziumArea(firstLenght, secondLenght, height int) int {
	return ((firstLenght + secondLenght) * height) / 2
}

func circleArea(radio float64) float64 {
	return pi2 * math.Pow(radio, 2)

}

func isParOrImparNumber(number int) {
	if number%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
}

func isUserAuthorized(user string, password string) bool {
	return user == "Maria" && password == "12345"
}
