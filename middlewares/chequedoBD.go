package middlewares

import (
	"net/http"

	"github.com/DonaldJoao/twittor/db"
)

/* ChequeoBD es el middleware que me permite conocer el estado de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
