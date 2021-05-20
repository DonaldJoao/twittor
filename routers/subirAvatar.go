package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/DonaldJoao/twittor/db"
	"github.com/DonaldJoao/twittor/models"
)

// SubirAvatar sube la imagen del avatar al servidor
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	var extencion = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extencion

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extencion
	status, err = db.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
