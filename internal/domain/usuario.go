package domain

import (
	"errors"
)

type Usuario struct {
	Name     string
	Edad     int
	Email    string
	Password string
}

var (
	ErrorNombreCorto   = errors.New("el nombre debe tener al menos 3 letras")
	ErrorMenorDeEdad   = errors.New("debes ser mayor de edad")
	ErrorEdadImposible = errors.New("la edad no puede ser registrada")
	ErrorEmail         = errors.New("formato incorrecto")
	ErrorPassword      = errors.New("debe tener minimo 8 caracteres")
)

func ValidateInformationUser(name string, edad int, email string, password string) error {
	if err := validarNombre(name); err != nil {
		return err
	}
	if err := validarEdad(edad); err != nil {
		return err
	}
	if err := validarEmail(email); err != nil {
		return err
	}
	if err := validarContraseña(password); err != nil {
		return err
	}
	return nil
}

func CrearUsuario(name string,edad int, email string, password string)(Usuario,error){
	if err := ValidateInformationUser(name,edad,email,password); err != nil{
		return Usuario{},err
	}
	return Usuario{
		Name: name,
		Edad: edad,
		Email: email,
		Password: password,
	},nil
}

func validarNombre(nombre string) error {
	if len(nombre) < 3 {
		return ErrorNombreCorto
	}
	return nil
}

func validarEdad(edad int) error {
	if edad <= 0 || edad >= 120 {
		return ErrorEdadImposible
	}
	if edad < 18 {
		return ErrorMenorDeEdad
	}
	return nil
}

func validarEmail(email string) error {
	for _, char := range email {
		if char == '@' {
			return nil
		}
	}
	return ErrorEmail
}

func validarContraseña(password string) error {
	if len(password) < 8 {
		return ErrorPassword
	}
	return nil
}
