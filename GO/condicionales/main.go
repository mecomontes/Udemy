package main

import (
	"fmt"
	"strconv"
)

var estado bool
var numero int
var texto string
var status bool

func main() {
	estado = true
	if estado = false; estado == true { //se puede asignar una condicion y evaluarla en la misma linea
		fmt.Println(estado)
	} else if estado == false {
		fmt.Println("algo loco")
	} else {
		fmt.Println("Es Falso")
	}

	switch numero := 5; numero {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("nada")
	}

}
