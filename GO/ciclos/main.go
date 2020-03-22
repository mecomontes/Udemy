package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	numero := 0
	for {
		fmt.Println("Continua")
		fmt.Println("Ingrese el numero: ")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	var j = 0

	for j < 10 {
		fmt.Printf("\n valor de j: %d", j) // El Printf no cambia de linea como si lo hace el Println
		if j == 5 {
			fmt.Printf(" multiplicamos por 2 \n")
			j = j * 2
			continue
		}
		fmt.Printf("Paso por aca \n")
		j++
	}

	var k int = 0

RUTINA: //Esto es una etiqueta de codigo
	for k < 10 {
		if k == 4 {
			k = k + 2
			fmt.Println("Voy a rutina")
			goto RUTINA //esto manda el codigo donde esta la etiqueta
		}
		fmt.Printf("valor de k: %d \n", k)
		k++
	}

}
