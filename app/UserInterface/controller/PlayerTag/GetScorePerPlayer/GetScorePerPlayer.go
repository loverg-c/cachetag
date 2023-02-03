package GetScorePerPlayer

import (
	"encoding/json"
	"net/http"
	"tags-finder/Infrastructure/Database/Repository"
)

func GetScorePerPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(repository.GetScorePerPlayers())

	//jwt := JWT.GenerateJWT()
	//var bearer = "Bearer " + jwt
	//
	//data, _ := query.Values(repository.GetScorePerPlayers())
	//// JSON body
	//body := []byte(data.Encode())
	//
	//log.Println(body)
	//
	//req, err := http.NewRequest("POST",
	//	"https://mercure/.well-known/mercure?topic=https://example.com/my-private-topic",
	//	bytes.NewBuffer(body))
	//
	//req.Header.Add("Authorization", bearer)
	//req.Header.Add("Content-Type", "application/json")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//
	//if err != nil {
	//	log.Println("Error on response.\n[ERROR] -", err)
	//}
	//defer resp.Body.Close()

}