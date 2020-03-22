package main

import (
	"bufio" //paquete para aceptar entradas por teclado
	"fmt"   //paquete para mostrar en pantalla
	"os"    //paquete para manejar el sistema operativo
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("ingrese numero 1: ")
	fmt.Scanf("%d", &numero1) // en windows se debe usar Scanln

	fmt.Println("ingrese numero 2: ")
	fmt.Scanf("%d", &numero2) // en linux se puede usar Scanf o Scanln

	fmt.Println("Descripcion: ")
	scanner := bufio.NewScanner(os.Stdin) //el paquete bufio permite obtener datos de entrada y el os permite obtener entradas por los perifericos
	if scanner.Scan() {                   //esto se hace solo para controlar la entrada, solo se ejecuta si obtengo algo por el teclado, en este cado
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
