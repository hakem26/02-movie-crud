package main

import (
	"example/moviecrud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main() {
	//define router
	router := mux.NewRouter()

	//routes
	routes.RegisterMovieRoutes(router)
	routes.RegisterUserRoutes(router)

	//starting server and log
	fmt.Print("Server Starts on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
