package ValidateTagController

import (
	"encoding/json"
	"net/http"
	Command "tags-finder/Application/Command/Player/PlayerValidateTag"
	"tags-finder/Infrastructure/Validator"
	"tags-finder/UserInterface/controller"
)

func ValidateTagController(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var input = ValidateTagInput{}
	err := decoder.Decode(&input)

	if err != nil {
		controller.ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	if ok, errors := Validator.ValidateInputs(input); !ok {
		controller.ValidationResponse(errors, w)
		return
	}

	//todo check already tagged

	Command.HandlePlayerValidateTag(
		Command.PlayerValidateTag{PlayerId: input.PlayerId, TagId: input.TagId},
	)

	fields := make(map[string]interface{})

	fields["status"] = "success"
	fields["message"] = "Tag has been validated"
	message, err := json.Marshal(fields)

	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}
