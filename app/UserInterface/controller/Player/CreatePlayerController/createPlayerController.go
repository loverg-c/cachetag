package CreatePlayerController

import (
	"encoding/json"
	"net/http"
	Command "tags-finder/Application/Command/Player/CreatePlayer"
	Query "tags-finder/Application/Query/Player/GetPlayer"
	"tags-finder/Infrastructure/Validator"
	"tags-finder/UserInterface/controller"
)

func PlayerCreateIndex(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var input = CreatePlayerInput{}
	err := decoder.Decode(&input)

	if err != nil {
		controller.ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	if ok, errors := Validator.ValidateInputs(input); !ok {
		controller.ValidationResponse(errors, w)
		return
	}

	existingPlayer := Query.HandleGetPlayer(Query.GetPlayer{Username: input.Username})

	if existingPlayer != nil {
		controller.ErrorResponse(http.StatusConflict, "Player already exist", w)
		return
	}

	player := Command.HandleCreatePlayer(
		Command.CreatePlayer{Username: input.Username},
	)

	json.NewEncoder(w).Encode(player)
}
