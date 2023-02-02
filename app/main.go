package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"tags-finder/Infrastructure/Database/config"
	"tags-finder/Infrastructure/Validator"
)

func main() {

	config.DatabaseInit()
	router := InitializeRouter()
	Validator.InitValidate()

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
