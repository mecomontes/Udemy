package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	archivo := "prueba.txt"
	f, err := os.Open(archivo)

	defer f.Close() //Exception para cerrar la funcion open en caso de inexistencia del archivo

	if err != nil { //Esto se ejecuta antes del defer
		fmt.Println("Error al abrir el archivo")
		os.Exit(1)
	}
	ejemploPanic()
}

func ejemploPanic() {
	defer func() { //No es necesario hacerlo usando una funcion
		//El difer solo ejecuta una instruccion. Si se necesitan ejecutar varias instrucciones se debe usar una funcion anonima
		reco := recover() //Si no hay un panic adentro de la funcion se ejecuta el defer. Si no hay panic, reco se lleva un nil
		if reco != nil {
			log.Fatalf("Ocurrio un error \n %v", reco) //Fatalf funciona igual al printf
		} //con esto se ataja el programa para que no se vaya al panic directamente
	}() //termina con parentesis porque la funcion no devuelve nada
	a := 1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}
}
