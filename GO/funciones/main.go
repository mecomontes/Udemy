package main

import "fmt"

func main() {
	fmt.Println(uno(5))

	numero, estado := dos(1)

	fmt.Println(numero)
	fmt.Println(estado)
	fmt.Println(Calculo(1, 2, 3, 4, 5, 5, 6))
	fmt.Println(Calculo(23, 45, 56))
	fmt.Println(Calculo(4))

}

func uno(numero int) (z int) { //el segundo int es para el valor de retorno
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int { //Los ... indican que no se sabe la cantidad de datos de entrada
	total := 0
	for item, num := range numero { //como contador se puede cambiar item por _
		total += total + num
		fmt.Printf("\n Item    %d \n", item)
	}
	return total
}
