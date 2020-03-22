package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("inicio")
	
	result = operacionMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionMidd(multiplicar)(2, 3)
	fmt.Println(result)
	result = operacionMidd(restar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {func(int, int)
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionMidd(f func(int, int) int) func(int, int) int{
	return func(a, b int){
		fmt.Println("inicio de operacion")
		return f(a, b)
	}
} 