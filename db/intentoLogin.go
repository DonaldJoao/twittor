package db

import (
	"github.com/DonaldJoao/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentoLogin realiza el chequeo de login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordByes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordByes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
