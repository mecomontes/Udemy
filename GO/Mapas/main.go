package main

import "fmt"

func main() {
	paises := make(map[string]string, 5) //Los mapas son el equivalente a los diccionarios
	paises["Mexico"] = "Mexico D.F."
	paises["Colombia"] = "Bogota"

	fmt.Println(paises)
	campeonato := map[string]int{
		"Barcelona":   23,
		"Real Madrid": 34,
		"Chivas":      24}

	campeonato["River Plate"] = 25
	campeonato["Chivas"] = 56
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s tiene un puntaje %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
