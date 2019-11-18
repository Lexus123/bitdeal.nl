package routes

import (
	"bitdeal.nl/models"
	"bitdeal.nl/server/handlers"
)

/*
GetRoutes ...
*/
var GetRoutes = models.Routes{
	models.Route{
		Name:        "GetHomePage",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: handlers.GetHomePage,
	},
}
