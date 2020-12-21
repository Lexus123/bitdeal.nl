package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

/*
Host will init the routing and fire up a server
*/
func Host() {
	// Make a new Mux router
	router := NewRouter()

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable  VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	if viper.GetString("environment") == "production" {
		router.PathPrefix("/static/").Handler(
			http.StripPrefix("/static/", http.FileServer(http.Dir("/var/www/bitdeal.nl/static"))),
		)
	} else {
		router.PathPrefix("/static/").Handler(
			http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
		)
	}

	log.Fatal(http.ListenAndServe(":8003", router))
}
