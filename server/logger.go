package server

import (
	"log"
	"net/http"
	"time"

	"bitdeal.nl/database"
)

/*
Logger logs everything the server does to the console.
*/
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		created := start.Unix()
		requestTime := time.Since(start).Milliseconds()

		if r.Method == "POST" && r.RequestURI == "/api/getprices" {
			database.SaveResponseTime(requestTime, created)
		}

		log.Printf(
			"%s\t%s\t%s\t%d",
			r.Method,
			r.RequestURI,
			name,
			requestTime,
		)
	})
}
