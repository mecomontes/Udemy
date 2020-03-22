package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()

}

func leoArchivo() { //lee el archivo pero no deja manipularlo
	archivo, err := ioutil.ReadFile("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./archivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
	} else {
		scanner := bufio.NewScanner(archivo) //lee la linea
		for scanner.Scan() {
			registro := scanner.Text() //convierte lo leido en una cadena de texto
			fmt.Println("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una linea nueva")
	archivo.Close()
}

func graboArchivo2() { //adicionar texto a un archivo existente
	filename := "./nuevoArchivo.txt"
	if Append(filename, "\nEsta es una segunda linea") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un error")
		return false
	}

	_, err = arch.WriteString(texto) //se adiciona lo que hay en texto al final del archivo
	if err != nil {
		fmt.Println("Hubo error")
		return false
	}
	return true
}
