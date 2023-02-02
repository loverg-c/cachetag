package main

import (
	"github.com/gorilla/mux"
	"tags-finder/UserInterface/controller/Player/CreatePlayerController"
	"tags-finder/UserInterface/controller/Player/PlayerListController"
	"tags-finder/UserInterface/controller/PlayerTag/GetScorePerPlayer"
	"tags-finder/UserInterface/controller/PlayerTag/PlayerHasValidateTagController"
	"tags-finder/UserInterface/controller/PlayerTag/ValidateTagController"
	"tags-finder/UserInterface/controller/Tag/ListTagController"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/players").
		Name("PlayerIndex").HandlerFunc(PlayerListController.PlayersIndex)
	router.Methods("POST").Path("/players").
		Name("PlayerCreateIndex").HandlerFunc(CreatePlayerController.PlayerCreateIndex)

	router.Methods("GET").Path("/tags").
		Name("TagIndex").HandlerFunc(ListTagController.TagsIndex)

	router.Methods("GET").Path("/players/validated_tags").
		Name("PlayerHasValidateTagIndex").HandlerFunc(PlayerHasValidateTagController.PlayerHasValidateTagIndex)
	router.Methods("POST").Path("/players/validated_tags").
		Name("ValidateTagController").HandlerFunc(ValidateTagController.ValidateTagController)

	router.Methods("GET").Path("/players/scores").
		Name("GetScorePerPlayers").HandlerFunc(GetScorePerPlayer.GetScorePerPlayers)

	return router
}
