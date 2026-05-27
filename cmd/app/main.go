package main

import (
	"fmt"

	"github.com/nico11alex/tienda-onlie/internal/domain"
)

func main() {
	name := "Nicola"
	edad := 25
	email := "nicolas@gmail.com"
	password := "Password123"
	user, err := domain.CrearUsuario(name,edad,email,password)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	fmt.Println("Usuario creado exitosamente")
}
