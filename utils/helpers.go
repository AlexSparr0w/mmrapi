package utils

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// func HandleWithApikey(r *http.Request) {
// 	apikey := r.Header.Get("apikey")
// 	check := model.DB.Exec("SELECT * FROM apikeys where apikey = $1", apikey)
// 	if check == nil {
// 		fmt.Println("no user found")
// 	}
// }
