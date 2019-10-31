package routes

import (
	"bitdeal.nl/models"
	"bitdeal.nl/server/handlers"
)

/*
PatchRoutes ...
*/
var PatchRoutes = models.Routes{
	// All PATCH requests
	models.Route{
		Name:        "PatchAddress",
		Method:      "PATCH",
		Pattern:     "/update",
		HandlerFunc: handlers.PatchAddress,
	},
}
