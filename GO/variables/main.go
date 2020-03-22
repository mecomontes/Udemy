package main

import (
	"fmt"
	"strconv"
)

var numero, numero2 int
var texto string
var status bool

func main() {
	numero3 := 4 //se crea la variable y se le asigna 4
	numero3 = 23 //se reasigna 23 a la variable
	numero5, numero6, numero7, texto1 := -4, 5, 6, "textonuevo"
	fmt.Println("Hola Mundo")
	fmt.Println(numero3)
	fmt.Println(numero5)
	fmt.Println(numero6)
	fmt.Println(numero7)
	fmt.Println(texto1)

	var moneda float64 = 2
	numero2 = int(moneda)
	texto1 = strconv.Itoa(int(moneda)) //se debe importar la libreria strconv
	fmt.Println(texto1)
}
