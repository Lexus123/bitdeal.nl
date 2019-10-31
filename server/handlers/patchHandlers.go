package handlers

// import (
// 	"net/http"
// 	"encoding/json"
// 	"io/ioutil"
// )

// /*
// Message ...
// */
// type Message struct {
// 	Address string `json:"address"`
// }

// /*
// PatchAddress ...
// */
// func PatchAddress(w http.ResponseWriter, r *http.Request) {
// 	// Read body
// 	b, err := ioutil.ReadAll(r.Body)
// 	defer r.Body.Close()

// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	// Unmarshal
// 	var msg Message
// 	err = json.Unmarshal(b, &msg)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	d1 := []byte("proxy_pass http://" + msg.Address + ";")
// 	err = ioutil.WriteFile("/var/www/node.lexvanderwerff.com/address", d1, 0644)

// 	output, err := json.Marshal(msg)
// 	if err != nil {
// 		http.Error(w, err.Error(), 500)
// 		return
// 	}

// 	w.Header().Set("content-type", "application/json")
// 	w.Write(output)

// 	// Check login creds
// 	// If OK
// 		// Update het bestand
// }
