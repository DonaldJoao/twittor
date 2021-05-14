package main

import (
	"log"

	"github.com/DonaldJoao/twittor/db"
	"github.com/DonaldJoao/twittor/handlers"
)

func main() {
	if db.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
