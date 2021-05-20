package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/DonaldJoao/twittor/middlewares"
	"github.com/DonaldJoao/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto, el Handler y pongo a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlewares.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/subirBanner", middlewares.ChequeoBD(middlewares.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlewares.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

	//COMENTARIO DE PRUEBAS GIT
}
