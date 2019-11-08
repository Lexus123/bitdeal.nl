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
	models.Route{
		Name:        "GetPrices",
		Method:      "GET",
		Pattern:     "/api/getprices",
		HandlerFunc: handlers.GetPrices,
	},
}
