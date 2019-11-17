package server

import (
	"log"
	"net/http"
)

/*
Host will init the routing and fire up a server
*/
func Host() {
	router := NewRouter()
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	)
	log.Fatal(http.ListenAndServe(":8003", router))
}
