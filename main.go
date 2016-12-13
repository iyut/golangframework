package main

import (
	"log"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/rs/cors"
	"framework/common"
	"framework/routers"
)

//Entry point of the program
func main() {

	//common.StartUp() - Replaced with init method
	// Get the mux router object
	router := routers.InitRoutes()

	// Get the CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	})
	// Create a negroni instance
	n := negroni.Classic()
	n.Use(c)
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
