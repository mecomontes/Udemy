package main

import (
	"fmt"
	"time"

	us "./user" //el archivo dentro de la carpeta debe llamrse igual
)

//Creacion de una estructura

type robinson struct {
	us.Usuario
}

func main() {
	u := new(robinson) //pepe e una herencia del tipo Usuario
	u.AltaUsuario(8063200, "Robinson Montes", time.Now(), true)
	fmt.Println(u.Usuario)
}
