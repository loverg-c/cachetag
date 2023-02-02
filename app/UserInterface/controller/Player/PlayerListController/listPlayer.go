package PlayerListController

import (
	"encoding/json"
	"net/http"
	repository "tags-finder/Infrastructure/Database/Repository"
)

func PlayersIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(repository.GetAllPlayer())
}
