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
		Name:        "GetStatsPage",
		Method:      "GET",
		Pattern:     "/stats",
		HandlerFunc: handlers.GetStatsPage,
	},
	models.Route{
		Name:        "GetInformationPage",
		Method:      "GET",
		Pattern:     "/informatie",
		HandlerFunc: handlers.GetInformationPage,
	},
	models.Route{
		Name:        "GetStats",
		Method:      "GET",
		Pattern:     "/api/getstats",
		HandlerFunc: handlers.GetStats,
	},
}
