package validate

import (
	"errors"
	"github.com/nico11alex/tienda-onlie/internal/domain"
)

var (
	ErrorNombreCorto   = errors.New("el nombre debe tener al menos 3 letras")
	ErrorMenorDeEdad   = errors.New("debes ser mayor de edad")
	ErrorEdadImposible = errors.New("la edad no puede ser registrada")
	ErrorEmail         = errors.New("formato incorrecto")
	ErrorPassword      = errors.New("debe tener minimo 8 caracteres")
)

func ValidateInfo(u domain.Usuario) error {
	if err := validarNombre(u.Name); err != nil {
		return err
	}
	if err := validarEdad(u.Edad); err != nil {
		return err
	}
	if err := validarEmail(u.Email); err != nil {
		return err
	}
	if err := validarContraseña(u.Password); err != nil {
		return err
	}
	return nil
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
