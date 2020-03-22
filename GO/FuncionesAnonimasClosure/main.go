package main

import "fmt"

// Funcion anonima
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("\n Suma = %d \n", Calculo(5, 7))

	//Se puede redefinir la funcion
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("\n Resta = %d \n", Calculo(6, 4))

	Operaciones()

	/* CLosure */

	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 2
		return a + b
	}
	fmt.Println("\n otra resta = ", resultado())
}

//Closure
func Tabla(valor int) func() int { //Regresa otra funcion
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero * secuencia
	}
}
