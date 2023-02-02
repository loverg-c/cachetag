package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"tags-finder/Infrastructure/Database/config"
)

func main() {

	config.DatabaseInit()
	router := InitializeRouter()

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
