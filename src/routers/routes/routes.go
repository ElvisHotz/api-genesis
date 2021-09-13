package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri      string
	Metodo   string
	Function func(http.ResponseWriter, *http.Request)
}

func ConfigRouters(r *mux.Router) *mux.Router {
	routers := routesExchanges

	for _, rout := range routers {
		r.HandleFunc(rout.Uri, rout.Function)

	}

	return r
}
