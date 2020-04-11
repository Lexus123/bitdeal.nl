package routes

import (
	"bitdeal.nl/models"
	"bitdeal.nl/server/handlers"
)

/*
PatchRoutes ...
*/
var PatchRoutes = models.Routes{
	models.Route{
		Name:        "PatchSunshade",
		Method:      "PATCH",
		Pattern:     "/api/sunshade",
		HandlerFunc: handlers.PatchSunshade,
	},
}
