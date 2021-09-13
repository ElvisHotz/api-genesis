package routes

import (
	"net/http"
	"teste-api/src/controllers"
)

var routesExchanges = []Route{
	{
		Uri:      "/exchange/{amount}/{from}/{to}/{rate}",
		Metodo:   http.MethodGet,
		Function: controllers.GetExchange,
	}, {
		Uri:      "/exchange",
		Metodo:   http.MethodGet,
		Function: controllers.GetExchanges,
	},
}
