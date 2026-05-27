package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNuevoUsuario_Validar(t *testing.T) {
	tests := []struct {
		name        string
		usuario     Usuario
		quiereError bool
		mensaje     string
	}{
		{
			name: "usuario_válido",
			usuario: Usuario{
				Name:     "Nicolas",
				Edad:     25,
				Email:    "nicolas@gmail.com",
				Password: "Password123",
			},
			quiereError: false,
		},
		{
			name: "nombre_muy_corto",
			usuario: Usuario{
				Name:     "Ni",
				Edad:     25,
				Email:    "nicolas@gmail.com",
				Password: "Password123",
			},
			quiereError: true,
			mensaje:     "el nombre debe tener al menos 3 letras",
		},
		{
			name: "edad_menor_de_18",
			usuario: Usuario{
				Name:     "Nicolas",
				Edad:     16,
				Email:    "nicolas@gmail.com",
				Password: "Password123",
			},
			quiereError: true,
			mensaje:     "el usuario debe ser mayor de edad",
		},
		{
			name: "contraseña_corta",
			usuario: Usuario{
				Name:     "Nicolas",
				Edad:     25,
				Email:    "nicolas@gmail.com",
				Password: "123",
			},
			quiereError: true,
			mensaje:     "la contraseña debe tener minimo 8 caracteres",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_,err := CrearUsuario(tt.usuario.Name, tt.usuario.Edad, tt.usuario.Email, tt.usuario.Password)

			if tt.quiereError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
