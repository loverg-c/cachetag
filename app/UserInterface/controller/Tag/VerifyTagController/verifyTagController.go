package VerifyTagController

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"tags-finder/Application/Query/Tag/GetTag"
	"tags-finder/Infrastructure/Validator"
	"tags-finder/UserInterface/controller"
)

func VerifyTagIndex(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tagId, ok := vars["id"]
	if !ok {
		controller.ErrorResponse(http.StatusBadRequest, "Id is wrong", w)
		return
	}

	decoder := json.NewDecoder(r.Body)

	input := VerifyTagInput{}
	err := decoder.Decode(&input)

	if err != nil {
		controller.ErrorResponse(http.StatusUnprocessableEntity, "Invalid JSON", w)
		return
	}

	if ok, errors := Validator.ValidateInputs(input); !ok {
		controller.ValidationResponse(errors, w)
		return
	}

	tag := GetTag.HandleGetTag(GetTag.GetTag{TagId: tagId})

	if tag == nil {
		controller.ErrorResponse(http.StatusNotFound, "Tag does not exist", w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	valid := struct {
		TagId int  `json:"id"`
		Valid bool `json:"valid"`
		Score int  `json:"score"`
	}{
		TagId: tag.Id,
		Valid: tag.Secret == input.Secret,
		Score: tag.Score,
	}

	json.NewEncoder(w).Encode(valid)
}
