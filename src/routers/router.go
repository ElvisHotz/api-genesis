package routers

import (
	"teste-api/src/routers/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRouters(r)
}
