package ValidateTagController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	Command "tags-finder/Application/Command/Player/PlayerValidateTag"
	"tags-finder/Application/Query/Player/GetScore"
	"tags-finder/Application/Query/Tag/GetTag"
	"tags-finder/Application/Query/Tag/GetValidatedByTagAndUser"
	"tags-finder/Infrastructure/Validator"
	"tags-finder/UserInterface/controller"
)

func ValidateTagController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerId, ok := vars["id"]
	if !ok {
		controller.ErrorResponse(http.StatusBadRequest, "Player id is wrong", w)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var input = ValidateTagInput{}
	err := decoder.Decode(&input)

	input.PlayerId, _ = strconv.Atoi(playerId)

	if err != nil {
		controller.ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON + "+err.Error(), w)
		return
	}

	if ok, errors := Validator.ValidateInputs(input); !ok {
		controller.ValidationResponse(errors, w)
		return
	}

	tagExist := GetTag.HandleGetTag(
		GetTag.GetTag{TagId: strconv.Itoa(input.TagId)},
	)

	if tagExist == nil {
		controller.ErrorResponse(http.StatusNotFound, "Ce tag n'existe pas", w)
		return
	}

	if tagExist.Secret != input.Secret {
		controller.ErrorResponse(http.StatusForbidden, "Mauvais secret", w)
		return
	}

	alreadyValidate := GetValidatedByTagAndUser.HandleGetValidatedByTagAndUser(
		GetValidatedByTagAndUser.GetValidatedByTagAndUser{PlayerId: input.PlayerId, TagId: input.TagId},
	)

	if alreadyValidate != nil {
		controller.ErrorResponse(http.StatusConflict, "Tag déjà validé", w)
		return
	}

	Command.HandlePlayerValidateTag(
		Command.PlayerValidateTag{PlayerId: input.PlayerId, TagId: input.TagId},
	)

	currentScore := GetScore.HandleGetScore(
		GetScore.GetScore{PlayerId: input.PlayerId},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(currentScore)
}
