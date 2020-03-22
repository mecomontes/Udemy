package main

import "fmt"

//las interfaces son un tipo de polimorfismo
type humano interface { //esta es la interfaz
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

/* Genero Humano*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true } // el * es un puntero que toma la interface
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Homber"
	} else {
		return "Mujer"
	}
}

func HumanoRespirando(hu humano) { //se crea el objeto hu al cual se le lleva la interfaz humano
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/*--------------------------------------*/
/* Genero Animal */

type perro struct {
	respirando bool
	comiendo   bool
	carnivo    bool
}

func (p *perro) respirar()         { p.respirando = true }
func (p *perro) comer()            { p.respirando = true }
func (p *perro) EsCarnivoro() bool { return p.carnivo }

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

func AnimalesRespirar(an animal) { //se crea el objeto hu al cual se le lleva la interfaz humano
	an.respirar()
	fmt.Printf("Soy un animal, y ya estoy respirando \n")
}

func main() {
	Pedro := new(hombre)
	Pedro.esHombre = true
	HumanoRespirando(Pedro)

	Maria := new(mujer)
	HumanoRespirando(Maria)

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros += AnimalesCarnivoros(Dogo)

	fmt.Printf("Total carnivoros %d", totalCarnivoros)
}
