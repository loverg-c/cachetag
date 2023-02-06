package GetScorePerPlayer

import (
	"encoding/json"
	"net/http"
	"tags-finder/Infrastructure/Database/Repository"
)

func GetScorePerPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	json.NewEncoder(w).Encode(repository.GetScorePerPlayers())
}
