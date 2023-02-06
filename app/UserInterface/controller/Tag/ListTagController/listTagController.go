package ListTagController

import (
	"encoding/json"
	"net/http"
	"tags-finder/Infrastructure/Database/Repository"
)

func TagsIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(repository.GetAllTag())
}
