package main

import (
	"fmt"
	"github.com/nico11alex/tienda-onlie/internal/Validate"
	"github.com/nico11alex/tienda-onlie/internal/domain"
)

func main() {
	user := domain.Usuario{
		Name:     "Nico",
		Edad:     17,
		Email:    "nicolas@gmail.com",
		Password: "bxwnwqwqx",
	}

	if err := validate.ValidateInfo(user); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Usuario creado exitosamente")
}
