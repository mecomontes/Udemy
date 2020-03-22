package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Meco") //la instruccion go hace que esto se ejecute en paralelo mientras el programa sigue en ejecucion
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
