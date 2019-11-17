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
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":8003", router))
}
