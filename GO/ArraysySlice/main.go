package main

import "fmt"

var tabla1 [10]int   //Tabla con 10 posiciones en 0, Vectos estatico
var matriz [5][7]int //Matriz

func main() {
	tabla1[0] = 1
	tabla1[5] = 14
	fmt.Println(tabla1)

	tabla := [10]int{5, 6, 8, 12, 34, 45, 2, 7, 34, 0}
	fmt.Println(tabla)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[4][3] = 45
	fmt.Println(matriz)

	//Slice es similar a una lista: Vector dinamico donde puedo cambiar su longitud
	lista := []int{2, 5, 4} //Cuando no se indica su longitud se crea un slice
	fmt.Println(lista)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:]
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos)) // Len es el tamaño del slice mientras que cap es la capacidad del slice

}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums)) //La capacidad reservada es de 2^n
}
