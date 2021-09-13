package main

import (
	"fmt"
	"log"
	"net/http"
	"teste-api/src/config"
	"teste-api/src/routers"
)

func main() {

	config.LoadConfig()

	r := routers.GenerateRouter()
	fmt.Printf("Escutando a porta %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
