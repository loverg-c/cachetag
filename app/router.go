package main

import (
	"github.com/gorilla/mux"
	"tags-finder/UserInterface/controller"
	playerController "tags-finder/UserInterface/controller/player"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/players").
		Name("PlayerIndex").HandlerFunc(controller.PlayersIndex)
	router.Methods("POST").Path("/players").
		Name("PlayerCreateIndex").HandlerFunc(playerController.PlayerCreateIndex)

	router.Methods("GET").Path("/tags").
		Name("TagIndex").HandlerFunc(controller.TagsIndex)

	router.Methods("GET").Path("/players/validated_tags").
		Name("PlayerHasValidateTagIndex").HandlerFunc(controller.PlayerHasValidateTagIndex)

	router.Methods("GET").Path("/players/scores").
		Name("GetScorePerPlayers").HandlerFunc(controller.GetScorePerPlayers)

	return router
}
