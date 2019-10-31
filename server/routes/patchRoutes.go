package routes

import (
	"node.lexvanderwerff.com/models"
	"node.lexvanderwerff.com/server/handlers"
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
