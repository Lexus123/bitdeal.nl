package routes

import (
	"bitdeal.nl/models"
	"bitdeal.nl/server/handlers"
)

/*
PostRoutes ...
*/
var PostRoutes = models.Routes{
	models.Route{
		Name:        "GetPrices",
		Method:      "POST",
		Pattern:     "/api/getprices",
		HandlerFunc: handlers.GetPrices,
	},
}
