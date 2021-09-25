package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"bitdeal.nl/database"
)

/*
SunshadeMessage ...
*/
type SunshadeMessage struct {
	Length int64 `json:"length"`
}

/*
PatchSunshade ...
*/
func PatchSunshade(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg SunshadeMessage
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	database.PatchSunshade(msg.Length)
}
